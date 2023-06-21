package cm_common

import (
	"fmt"
	"math/big"
	"sort"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
)

// range is [from, to)  or just before newBlockNum
// - gets all the directokentransfer by block
func (mdl *CMCommon) UpdateSessionWithDirectTokenTransferBefore(tillBlock int64) {
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

func DirecTokenTransferString(repo ds.RepositoryI, tx *schemas.TokenTransfer) string {
	msg := fmt.Sprintf("DirectTokenTransfer(%s) %f %s at %d from %s to %s",
		tx.TxHash,
		utils.GetFloat64Decimal(tx.Amount.Convert(), repo.GetToken(tx.Token).Decimals),
		repo.GetToken(tx.Token).Symbol,
		tx.BlockNum, tx.From, tx.To,
	)
	return msg
}

// if blockNum is lasteventblock then set session is Updated
// if blockNum is not equal to lasteventblock then fetch details for that session
func (mdl *CMCommon) processDirectTransfersOnBlock(blockNum int64, sessionIDToTxs map[string][]*schemas.TokenTransfer) {
	for sessionID, txs := range sessionIDToTxs {
		session := mdl.Repo.GetCreditSession(sessionID)
		txsList := schemas.TokenTransferList(txs)
		sort.Sort(txsList)
		for _, tx := range txsList {
			if session.Account == tx.From {
				// USDT in transferFrom emits event even if the amount is zero
				if tx.Amount.Convert().Cmp(big.NewInt(0)) == 0 {
					continue
				}
				log.Fatalf("Token withdrawn directly from account %v", DirecTokenTransferString(mdl.Repo, tx))
			}
			//
			mdl.SetSessionIsUpdated(sessionID)
			//
			if tx.To == session.Account {
				// add transfer as collateral for rewardClaimed too
				// reward token is enabled to account, and will be counted as user fund
				// https://github.com/Gearbox-protocol/integrations-v2/blob/main/contracts/adapters/convex/ConvexV1_BaseRewardPool.sol#L292-L298
				mdl.AddCollateralToSession(tx.BlockNum, sessionID, tx.Token, tx.Amount.Convert())
			}
			mdl.onDirectTokenTransfer(mdl.Repo, tx, session)
		}
		mdl.onDirectTokenTransfer(mdl.Repo, nil, nil)
	}
}
