package rebase_token

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/core/types"
)

func (mdl *RebaseToken) OnLog(txLog types.Log) {
}
func (mdl *RebaseToken) Query(queryTill int64) {
	if queryTill-mdl.LastSync <= 0 {
		return
	}
	//
	newDetails := mdl.GetstETHDetails(queryTill)
	prevDetails := mdl.previousDetails()
	//
	mdl.cmpRatiosAndAdd(prevDetails, newDetails)
	mdl.save(newDetails)
}

// prevDetails is already added.
func (mdl *RebaseToken) cmpRatiosAndAdd(prevDetails, newDetails stETHValues) {
	if newDetails.blockNum-prevDetails.blockNum <= 1 {
		if newDetails.ratio().Cmp(prevDetails.ratio()) != 0 {
			mdl.Repo.AddRebaseDetailsForDB(newDetails.ToDB())
		}
		return
	}
	prevRatio := prevDetails.ratio()
	newRatio := newDetails.ratio()
	if prevRatio.Cmp(newRatio) != 0 {
		blockNum := (prevDetails.blockNum + newDetails.blockNum) / 2
		midDetails := mdl.GetstETHDetails(blockNum)
		midRatio := midDetails.ratio()
		if midRatio.Cmp(prevRatio) != 0 {
			mdl.cmpRatiosAndAdd(prevDetails, midDetails)
		}
		if newRatio.Cmp(midRatio) != 0 {
			mdl.cmpRatiosAndAdd(midDetails, newDetails)
		}
	}
}

func (mdl *RebaseToken) getBigIntFromDetails(field string) *big.Int {
	totalSharesStr := mdl.Details[field]
	if totalSharesStr == nil {
		return new(big.Int)
	}
	return utils.StringToInt(totalSharesStr.(string))
}

func (mdl *RebaseToken) AfterSyncHook(block int64) {
	mdl.SyncAdapter.AfterSyncHook(block)
}
