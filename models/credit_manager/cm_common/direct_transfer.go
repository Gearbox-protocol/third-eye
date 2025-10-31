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
func (mdl *CommonCMAdapter) UpdateSessionWithDirectTokenTransferBefore(tillBlock int64) {
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

// nikitale account
var gearbox_addr_so_collateral_is_not_from_user = []string{"0x7b065Fcb0760dF0CEA8CFd144e08554F3CeA73D1",
	"0x6f378f36899cEB7C6fB7D293aAE1ca86B0Edbf6D",
	"0xFC85C07C3e0D497d97F287a70C6b2fA5CD5fdBE0",
	"0x23e85dB353bFDa25329c4e12b98Aa991F541eA2d"}

// if blockNum is lasteventblock then set session is Updated
// if blockNum is not equal to lasteventblock then fetch details for that session
func (mdl *CommonCMAdapter) processDirectTransfersOnBlock(blockNum int64, sessionIDToTxs map[string][]*schemas.TokenTransfer) {
	for sessionID, txs := range sessionIDToTxs {
		session := mdl.Repo.GetCreditSession(sessionID)
		txsList := schemas.TokenTransferList(txs)
		sort.Sort(txsList)
		for _, tx := range txsList {
			if tx.Token == "0x9D65fF81a3c488d585bBfb0Bfe3c7707c7917f54" ||
				tx.TxHash == "0x3eb5d6c93d73517cf1b927e998900fc066dc2912f744178bcb3557ad9b7e5526" { // uni transfer on v1.
				continue
			}
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
			if tx.To == session.Account && !utils.Contains(gearbox_addr_so_collateral_is_not_from_user, tx.From) {
				// add transfer as collateral for rewardClaimed too
				// reward token is enabled to account, and will be counted as user fund
				// https://github.com/Gearbox-protocol/integrations-v2/blob/main/contracts/adapters/convex/ConvexV1_BaseRewardPool.sol#L292-L298
				log.Infof("in tx %s , transferred to %s", tx.TxHash, sessionID)
				mdl.AddCollateralToSession(tx.BlockNum, sessionID, tx.Token, tx.Amount.Convert())
			}
			mdl.onDirectTokenTransfer(mdl.Repo, tx, session)
		}
		mdl.onDirectTokenTransfer(mdl.Repo, nil, nil)
	}
}
