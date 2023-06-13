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
// - gets all the directokentransfer by block
func (mdl *CreditManager) UpdateSessionWithDirectTokenTransferBefore(tillBlock int64) {
	data := mdl.Repo.GetAccountManager().CheckTokenTransfer(mdl.GetAddress(), 0, tillBlock)
	blockNums := []int64{}
	for blockNum := range data {
		blockNums = append(blockNums, blockNum)
	}
	sort.Slice(blockNums, func(i, j int) bool { return blockNums[i] < blockNums[j] })
	for _, blockNum := range blockNums {
		mdl.processDirectTransfersOnBlock(blockNum, data[blockNum])
		calls, processFn := mdl.FetchFromDCForChangedSessions(blockNum)
		results := core.MakeMultiCall(mdl.Client, blockNum, false, calls)
		for i, result := range results {
			processFn[i](result)
		}
	}
}

func (mdl CreditManager) DirecTokenTransferString(tx *schemas.TokenTransfer) string {
	msg := fmt.Sprintf("DirectTokenTransfer(%s) %f %s at %d from %s to %s",
		tx.TxHash,
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
func (mdl *CreditManager) processDirectTransfersOnBlock(blockNum int64, sessionIDToTxs map[string][]*schemas.TokenTransfer) {
	for sessionID, txs := range sessionIDToTxs {
		session := mdl.Repo.GetCreditSession(sessionID)
		txsList := schemas.TokenTransferList(txs)
		sort.Sort(txsList)
		currentRewardClaim := newRewardClaimDetails()
		for _, tx := range txsList {
			if session.Account == tx.From {
				// USDT in transferFrom emits event even if the amount is zero
				if tx.Amount.Convert().Cmp(big.NewInt(0)) == 0 {
					continue
				}
				log.Fatalf("Token withdrawn directly from account %v", mdl.DirecTokenTransferString(tx))
			}
			//
			mdl.setUpdateSession(sessionID)
			//
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
				mdl.Repo.RecentMsgf(log.RiskHeader{
					BlockNumber: tx.BlockNum,
					EventCode:   "AMQP",
				}, "Deposit: %s", mdl.DirecTokenTransferString(tx))
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
	}
}
