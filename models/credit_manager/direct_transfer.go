package credit_manager

import (
	"fmt"
	"math/big"
	"sort"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

// range is [from, to)  or just before newBlockNum
//
// - gets all the directokentransfer by block
// - if no transfer or first directokentransfer is after lastEventBlock, fetch account details from datacompressor
// - if mdl.lastEventBlock== first directtokentransfer block
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

func (mdl CreditManager) DirecTokenTransferString(tx *schemas.TokenTransfer) string {
	msg := fmt.Sprintf("DirectTokenTransfer %f %s at %d from %s to %s",
		utils.GetFloat64Decimal(tx.Amount.Convert(), mdl.Repo.GetToken(tx.Token).Decimals),
		mdl.Repo.GetToken(tx.Token).Symbol,
		tx.BlockNum, tx.From, tx.To,
	)
	return msg
}

type RewardClaimDetails struct {
	Args      core.Json
	TxHash    string
	LogID     uint
	Transfers core.Transfers
}

func (mdl *CreditManager) addRewardClaimAccountOperation(blockNum int64, sessionId, borrower string, claim RewardClaimDetails) {
	mdl.Repo.AddAccountOperation(&schemas.AccountOperation{
		TxHash:      claim.TxHash,
		BlockNumber: blockNum,
		LogId:       claim.LogID,
		Borrower:    borrower,
		SessionId:   sessionId,
		Dapp:        core.NULL_ADDR.Hex(),
		Action:      "RewardClaimed",
		Args:        &claim.Args,
		AdapterCall: false,
		Transfers:   &claim.Transfers,
	})
}

func newRewardClaimDetails() RewardClaimDetails {
	return RewardClaimDetails{
		Args:      core.Json{},
		Transfers: core.Transfers{},
	}
}

// if blockNum is lasteventblock then set session is Updated
// if blockNum is not equal to lasteventblock then fetch details for that session
func (mdl *CreditManager) ProcessDirectTransfersOnBlock(blockNum int64, sessionIDToTxs map[string][]*schemas.TokenTransfer) {
	for sessionID, txs := range sessionIDToTxs {
		session := mdl.Repo.GetCreditSession(sessionID)
		txsList := schemas.TokenTransferList(txs)
		sort.Sort(txsList)
		currentRewardClaim := newRewardClaimDetails()
		for _, tx := range txsList {
			if session.Account == tx.From {
				// withdrawAmount := new(big.Int).Neg(tx.Amount.Convert())
				// mdl.Repo.RecentEventMsg(tx.BlockNum, "Withdrawn(%s): %s", sessionID, tx)
				log.Fatalf("Token withdrawn directly from account %v", mdl.DirecTokenTransferString(tx))
			}
			if blockNum == mdl.lastEventBlock {
				mdl.setUpdateSession(sessionID)
			}
			var amount *big.Int
			if tx.To == session.Account {
				amount = tx.Amount.Convert()
				// add transfer as collateral for rewardClaimed too
				// reward token is enabled to account, and will be counted as user fund
				// https://github.com/Gearbox-protocol/integrations-v2/blob/main/contracts/adapters/convex/ConvexV1_BaseRewardPool.sol#L292-L298
				mdl.AddCollateralToSession(tx.BlockNum, sessionID, tx.Token, amount)
			}
			// rewardPaid doesn't emitted executeOrder or any other gearbox event
			if tx.From == core.NULL_ADDR.Hex() || mdl.allowedProtocols[tx.From] {
				// different rewardclaim
				if currentRewardClaim.TxHash != "" && currentRewardClaim.TxHash != tx.TxHash {
					mdl.addRewardClaimAccountOperation(blockNum, sessionID, session.Borrower, currentRewardClaim)
					currentRewardClaim = newRewardClaimDetails()
				}
				currentRewardClaim.Args[tx.Token] = tx.From
				currentRewardClaim.Transfers[tx.Token] = amount
				currentRewardClaim.TxHash = tx.TxHash
				currentRewardClaim.LogID = tx.LogID
			} else {
				mdl.Repo.RecentEventMsg(tx.BlockNum, "Deposit: %s", mdl.DirecTokenTransferString(tx))
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
		}
		if currentRewardClaim.TxHash != "" {
			mdl.addRewardClaimAccountOperation(blockNum, sessionID, session.Borrower, currentRewardClaim)
		}
		// for blocks without credit manager events, update session
		if blockNum != mdl.lastEventBlock {
			//  works similar to FetchFromDCForChangedSessions, only for single session
			mdl.updateSession(sessionID, blockNum)
		}
	}
}
