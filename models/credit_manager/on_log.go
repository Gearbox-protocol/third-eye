package credit_manager

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"sort"
)

func (mdl *CreditManager) processExecuteEvents() {
	if len(mdl.executeParams) > 0 {
		mdl.handleExecuteEvents()
		mdl.executeParams = []core.ExecuteParams{}
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

func (mdl *CreditManager) ProcessDirectTransfersOnBlock(blockNum int64, sessionIDToTxs map[string][]*core.TokenTransfer) {
	for sessionID, txs := range sessionIDToTxs {
		session := mdl.Repo.GetCreditSession(sessionID)
		for _, tx := range txs {
			var amount *big.Int
			switch session.Account {
			case tx.From:
				amount = new(big.Int).Neg(tx.Amount.Convert())
				mdl.Repo.RecentEventMsg(tx.BlockNum, "Direct Token Withdrawn %v, id: %s", tx, sessionID)
				log.Fatalf("Token withdrawn directly from account %v", tx)
			case tx.To:
				amount = tx.Amount.Convert()
				mdl.AddCollateralToSession(tx.BlockNum, sessionID, tx.Token, amount)
				mdl.Repo.RecentEventMsg(tx.BlockNum, "Direct Token Deposit %v", tx)
			}
			if blockNum == mdl.lastEventBlock {
				mdl.UpdatedSessions[sessionID]++
			}
			mdl.Repo.AddAccountOperation(&core.AccountOperation{
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
	// storing execute order in a single tx and processing them in a single go on next tx
	// for credit session stats
	//
	// execute events are matched with tenderly response to get transfers for each events
	if mdl.LastTxHash != txLog.TxHash.Hex() {
		mdl.processExecuteEvents()
		mdl.LastTxHash = txLog.TxHash.Hex()
	}
	// for credit manager stats
	blockNum := int64(txLog.BlockNumber)
	if mdl.lastEventBlock != blockNum {
		mdl.onBlockChange(blockNum)
	}
	mdl.lastEventBlock = blockNum
	mdl.Repo.GetAccountManager().DeleteTxHash(blockNum, txLog.TxHash.Hex())
	//-- for credit manager stats
	switch txLog.Topics[0] {
	case core.Topic("OpenCreditAccount(address,address,address,uint256,uint256,uint256)"):
		openCreditAccountEvent, err := mdl.contractETH.ParseOpenCreditAccount(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack OpenCreditAccount event", err)
		}

		mdl.onOpenCreditAccount(&txLog, openCreditAccountEvent.Sender.Hex(),
			openCreditAccountEvent.OnBehalfOf.Hex(),
			openCreditAccountEvent.CreditAccount.Hex(),
			openCreditAccountEvent.Amount,
			openCreditAccountEvent.BorrowAmount,
			openCreditAccountEvent.ReferralCode)
	case core.Topic("CloseCreditAccount(address,address,uint256)"):
		closeCreditAccountEvent, err := mdl.contractETH.ParseCloseCreditAccount(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack CloseCreditAccount event", err)
		}

		mdl.onCloseCreditAccount(&txLog,
			closeCreditAccountEvent.Owner.Hex(),
			closeCreditAccountEvent.To.Hex(),
			closeCreditAccountEvent.RemainingFunds)
	case core.Topic("LiquidateCreditAccount(address,address,uint256)"):
		liquidateCreditAccountEvent, err := mdl.contractETH.ParseLiquidateCreditAccount(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack LiquidateCreditAccount event", err)
		}

		mdl.onLiquidateCreditAccount(&txLog,
			liquidateCreditAccountEvent.Owner.Hex(),
			liquidateCreditAccountEvent.Liquidator.Hex(),
			liquidateCreditAccountEvent.RemainingFunds)
	case core.Topic("RepayCreditAccount(address,address)"):
		repayCreditAccountEvent, err := mdl.contractETH.ParseRepayCreditAccount(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack RepayCreditAccount event", err)
		}

		mdl.onRepayCreditAccount(&txLog,
			repayCreditAccountEvent.Owner.Hex(),
			repayCreditAccountEvent.To.Hex())
	case core.Topic("IncreaseBorrowedAmount(address,uint256)"):
		increaseBorrowEvent, err := mdl.contractETH.ParseIncreaseBorrowedAmount(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack IncreaseBorrowedAmount event", err)
		}

		mdl.onIncreaseBorrowedAmount(&txLog,
			increaseBorrowEvent.Borrower.Hex(), increaseBorrowEvent.Amount)
	case core.Topic("AddCollateral(address,address,uint256)"):
		addCollateralEvent, err := mdl.contractETH.ParseAddCollateral(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack AddCollateral event", err)
		}

		mdl.onAddCollateral(&txLog, addCollateralEvent.OnBehalfOf.Hex(),
			addCollateralEvent.Token.Hex(),
			addCollateralEvent.Value)
	case core.Topic("ExecuteOrder(address,address)"):
		execute, err := mdl.contractETH.ParseExecuteOrder(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack ExecuteOrder event", err)
		}
		mdl.AddExecuteParams(&txLog, execute.Borrower, execute.Target)
	case core.Topic("NewParameters(uint256,uint256,uint256,uint256,uint256,uint256)"):
		params, err := mdl.contractETH.ParseNewParameters(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack NewParameters event", err)
		}
		mdl.State.MinAmount = (*core.BigInt)(params.MinAmount)
		mdl.State.MaxAmount = (*core.BigInt)(params.MaxAmount)
		mdl.State.MaxLeverageFactor = params.MaxLeverage.Int64()
		mdl.State.FeeInterest = params.FeeInterest.Int64()
		mdl.Repo.AddParameters(txLog.Index, txLog.TxHash.Hex(), &core.Parameters{
			BlockNum:            blockNum,
			CreditManager:       mdl.GetAddress(),
			MinAmount:           (*core.BigInt)(params.MinAmount),
			MaxAmount:           (*core.BigInt)(params.MaxAmount),
			MaxLeverage:         (*core.BigInt)(params.MaxLeverage),
			FeeInterest:         (*core.BigInt)(params.FeeInterest),
			FeeLiquidation:      (*core.BigInt)(params.FeeLiquidation),
			LiquidationDiscount: (*core.BigInt)(params.LiquidationDiscount),
		}, mdl.State.UnderlyingToken)
	case core.Topic("TransferAccount(address,address)"):
		if len(txLog.Data) == 0 { // oldowner and newowner are indexed
			transferAccount, err := mdl.contractETH.ParseTransferAccount(txLog)
			if err != nil {
				log.Fatal("[CreditManagerModel]: Cant unpack TransferAccount event", err)
			}
			mdl.onTransferAccount(&txLog, transferAccount.OldOwner.Hex(), transferAccount.NewOwner.Hex())
		} else {
			oldOwner := common.BytesToAddress(txLog.Data[:32])
			newOwner := common.BytesToAddress(txLog.Data[32:])
			mdl.onTransferAccount(&txLog, oldOwner.Hex(), newOwner.Hex())
		}
	}
}
