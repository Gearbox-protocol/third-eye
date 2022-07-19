package credit_manager

import (
	"math/big"
	"sort"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (mdl *CreditManager) processExecuteEvents() {
	if len(mdl.executeParams) > 0 {
		mdl.handleExecuteEvents(mdl.executeParams)
		mdl.executeParams = []ds.ExecuteParams{}
	}
}

func (mdl *CreditManager) ProcessDirectTokenTransfer(oldBlockNum, newBlockNum int64) {
	data := mdl.Repo.GetAccountManager().CheckTokenTransfer(mdl.GetAddress(), oldBlockNum, newBlockNum)
	blockNums := []int64{}
	for blockNum, _ := range data {
		blockNums = append(blockNums, blockNum)
	}
	sort.Slice(blockNums, func(i, j int) bool { return blockNums[i] < blockNums[j] })
	for _, blockNum := range blockNums {
		mdl.ProcessDirectTransfersOnBlock(blockNum, data[blockNum])
	}
}

func (mdl *CreditManager) ProcessAccountEvents(newBlockNum int64) {
	data := mdl.Repo.GetAccountManager().CheckTokenTransfer(mdl.GetAddress(), mdl.lastEventBlock, newBlockNum)
	blockNums := []int64{}
	for blockNum, _ := range data {
		blockNums = append(blockNums, blockNum)
	}
	sort.Slice(blockNums, func(i, j int) bool { return blockNums[i] < blockNums[j] })
	// no direct token transfer or the first token transfer is after the mdl.lastEventBlock
	// or start block of range for check token tranafer
	if len(blockNums) == 0 || blockNums[0] > mdl.lastEventBlock {
		mdl.FetchFromDCForChangedSessions(mdl.lastEventBlock)
	}
	for _, blockNum := range blockNums {
		mdl.ProcessDirectTransfersOnBlock(blockNum, data[blockNum])
		// if there are direct token tranfer on the start block of range then
		// use changed sessions
		if blockNum == mdl.lastEventBlock {
			mdl.FetchFromDCForChangedSessions(mdl.lastEventBlock)
		}
	}
}

func (mdl *CreditManager) ProcessDirectTransfersOnBlock(blockNum int64, sessionIDToTxs map[string][]*schemas.TokenTransfer) {
	for sessionID, txs := range sessionIDToTxs {
		session := mdl.Repo.GetCreditSession(sessionID)
		txsList := schemas.TokenTransferList(txs)
		sort.Sort(txsList)
		for _, tx := range txsList {
			var amount *big.Int
			switch session.Account {
			case tx.From:
				amount = new(big.Int).Neg(tx.Amount.Convert())
				mdl.Repo.RecentEventMsg(tx.BlockNum, "Withdrawn(%s): %s", sessionID, tx)
				log.Fatalf("Token withdrawn directly from account %v", tx)
			case tx.To:
				amount = tx.Amount.Convert()
				mdl.AddCollateralToSession(tx.BlockNum, sessionID, tx.Token, amount)
				mdl.Repo.RecentEventMsg(tx.BlockNum, "Deposit: %s", tx)
			}
			if blockNum == mdl.lastEventBlock {
				mdl.setUpdateSession(sessionID)
			}
			mdl.Repo.AddAccountOperation(&schemas.AccountOperation{
				TxHash:      tx.TxHash,
				BlockNumber: tx.BlockNum,
				LogId:       tx.LogID,
				Borrower:    session.Borrower,
				SessionId:   sessionID,
				Dapp:        tx.Token,
				Action:      "DirectTokenTransfer",
				Args:        &core.Json{"amount": amount, "to": tx.To, "from": tx.From},
				AdapterCall: false,
				Transfers:   &core.Transfers{tx.Token: amount},
			})
		}
		// for blocks without credit manager events, update session
		if blockNum != mdl.lastEventBlock {
			mdl.updateSession(sessionID, blockNum)
		}
	}
}

// works for newBlockNum > mdl.lastEventBlock
func (mdl *CreditManager) onBlockChange(newBlockNum int64) {
	// on each new block
	mdl.ProcessAccountEvents(newBlockNum)
	// datacompressor works for cm address only after the address is registered with contractregister
	// i.e. discoveredAt
	// only after each event block.
	if mdl.lastEventBlock != 0 && mdl.lastEventBlock >= mdl.DiscoveredAt {
		mdl.calculateCMStat(mdl.lastEventBlock)
		mdl.lastEventBlock = 0
	}
}

func (mdl *CreditManager) OnLog(txLog types.Log) {
	// creditConfigurator events for test
	// we only require CreditFacadeUpgraded so that we can update the details for credit manager and
	if mdl.GetDetailsByKey("configurator") == txLog.Address.Hex() {
		switch txLog.Topics[0] {
		case core.Topic("CreditFacadeUpgraded(address)"):
			facade := utils.ChecksumAddr(txLog.Topics[1].Hex())
			mdl.SetCreditFacadeContract(common.HexToAddress(facade))
		case core.Topic("LimitsUpdated(uint256,uint256)"): // SET min/max BorrowAmount
			minAmount := new(big.Int).SetBytes(txLog.Data[:32])
			maxAmount := new(big.Int).SetBytes(txLog.Data[32:])
			mdl.State.MinAmount = (*core.BigInt)(minAmount)
			mdl.State.MaxAmount = (*core.BigInt)(maxAmount)
		}
		return
	}
	// on txHash
	switch mdl.GetVersion() {
	case 1:
		// storing execute order in a single tx and processing them in a single go on next tx
		// for credit session stats
		//
		// execute events are matched with tenderly response to get transfers for each events
		if mdl.LastTxHash != txLog.TxHash.Hex() {
			mdl.processExecuteEvents()
			mdl.LastTxHash = txLog.TxHash.Hex()
		}
	case 2:
		mdl.onNewTxHashV2(txLog.TxHash.Hex())
	}
	// on new block
	// for credit manager stats
	blockNum := int64(txLog.BlockNumber)
	if mdl.lastEventBlock != blockNum {
		mdl.onBlockChange(blockNum)
	}
	mdl.lastEventBlock = blockNum
	//
	mdl.Repo.GetAccountManager().DeleteTxHash(blockNum, txLog.TxHash.Hex())
	switch mdl.GetVersion() {
	case 1:
		mdl.checkLogV1(txLog)
	case 2:
		mdl.checkLogV2(txLog)
	}
}
