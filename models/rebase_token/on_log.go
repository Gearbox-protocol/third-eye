package rebase_token

import (
	"encoding/hex"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// func (mdl *RebaseToken) saveUpdatesFromValidatorsJson(blockNum int64, saveLast bool) {
// for _, update := range mdl.validatorHandler.GetValuesBefore(blockNum) {
// 	mdl.state.DepositValidators = update.Validator
// 	if update.Block != blockNum || saveLast {
// 		//TODO: save
// 		mdl.Repo.AddRebaseDetailsForDB(mdl.state.GetDataForDB(update.Block))
// 	}
// }
// }

func (mdl *RebaseToken) save(blockNum int64) {
	if blockNum == 0 {
		return
	}
	dataToSave := mdl.state.GetDataForDB(blockNum)
	newRatio := getETHToSharesRatio(dataToSave)
	if utils.DiffMoreThanFraction(newRatio, mdl.prevRatio, big.NewFloat(.001)) { // .1%
		mdl.Repo.AddRebaseDetailsForDB(dataToSave)
		mdl.prevRatio = newRatio
	}
}
func getETHToSharesRatio(dataToSave *schemas.RebaseDetailsForDB) *big.Int {
	return new(big.Int).Quo(utils.GetInt64(dataToSave.TotalETH.Convert(), -18), dataToSave.TotalShares.Convert())
}
func (mdl *RebaseToken) OnLog(txLog types.Log) {
	blockNum := int64(txLog.BlockNumber)
	if mdl.lastBlockNum != blockNum {
		// TODO:save
		mdl.save(mdl.lastBlockNum)
	}
	// mdl.saveUpdatesFromValidatorsJson(blockNum, false)
	mdl.lastBlockNum = blockNum
	switch txLog.Topics[0] {
	case core.Topic("SetApp(bytes32,bytes32,address)"):
		namespace := hex.EncodeToString(txLog.Topics[1][:])
		appId := hex.EncodeToString(txLog.Topics[2][:])
		if appId == "3ca7c3e38968823ccb4c78ea688df41356f182ae1d159e4ee608d30d68cef320" &&
			namespace == "f1f3eb40f5bc1ad1344716ced8b8a0431d840b5783aea1fd01786bc26f35ac0f" {
			implementationAddr := common.BytesToAddress(txLog.Data)
			mdl.checkImplementationAddr(implementationAddr.Hex(), txLog.TxHash.Hex())
		}
	///////////////
	// for position: DEPOSITED_VALIDATORS_POSITION
	case core.Topic("DepositedValidatorsChanged(uint256)"):
		event, err := mdl.contract.ParseDepositedValidatorsChanged(txLog)
		log.CheckFatal(err)
		mdl.state.DepositValidators = event.DepositedValidators.Int64()

	///////////////
	// for position: BUFFERED_ETHER_POSITION
	case core.Topic("Unbuffered(uint256)"):
		delta := new(big.Int).SetBytes(txLog.Data)
		if mdl.state.DepositBalance == nil {
			mdl.state.DepositBalance = core.NewBigInt(nil)
		}
		mdl.state.DepositBalance = (*core.BigInt)(
			new(big.Int).Add(mdl.state.DepositBalance.Convert(), delta))
	case core.Topic("Submitted(address,uint256,address)"):
		event, err := mdl.contract.ParseSubmitted(txLog)
		log.CheckFatal(err)
		mdl.state.DepositBalance = (*core.BigInt)(
			new(big.Int).Add(mdl.state.DepositBalance.Convert(), event.Amount))

	case core.Topic("ETHDistributed(uint256,uint256,uint256,uint256,uint256,uint256)"):
		event, err := mdl.contract.ParseETHDistributed(txLog)
		log.CheckFatal(err)
		mdl.state.DepositBalance = (*core.BigInt)(event.PostBufferedEther)
		mdl.state.CLBalance = (*core.BigInt)(event.PostCLBalance)

	//////////////
	// for position: CL_VALIDATORS_POSITION
	case core.Topic("CLValidatorsUpdated(uint256,uint256,uint256)"):
		event, err := mdl.contract.ParseCLValidatorsUpdated(txLog)
		log.CheckFatal(err)
		mdl.state.CLValidators = event.PostCLValidators.Int64()

		///////////
		//
	case core.Topic("TokenRebased(uint256,uint256,uint256,uint256,uint256,uint256,uint256)"):
		// log.Info("check at ", txLog.BlockNumber)
		// log.Info(utils.ToJson(mdl.state))
		// log.Info(utils.ToJson(mdl.getStateAt(int64(txLog.BlockNumber))))
		//
		event, err := mdl.contract.ParseTokenRebased(txLog)
		log.CheckFatal(err)
		if mdl.state.GetPostTotalEther().Cmp(event.PostTotalEther) != 0 {
			log.Fatalf("TotalETH(%d) calculated is different from tokenRebased event reported value(%d) at txHash(%s)",
				mdl.state.GetPostTotalEther(), event.PostTotalEther, txLog.TxHash)
		}
		if mdl.state.TotalShares.Convert().Cmp(event.PostTotalShares) != 0 {
			log.Fatalf("TotalShares(%d) calculated is different from tokenRebase event reported value(%d) at txHash(%s)",
				mdl.state.TotalShares.Convert(), event.PostTotalShares, txLog.TxHash)
		}
	// totalShares
	case core.Topic("SharesBurnt(address,uint256,uint256,uint256)"):
		event, err := mdl.contract.ParseSharesBurnt(txLog)
		log.CheckFatal(err)
		mdl.state.TotalShares = (*core.BigInt)(new(big.Int).Sub(mdl.state.TotalShares.Convert(), event.SharesAmount))
	case core.Topic("TransferShares(address,address,uint256)"):
		event, err := mdl.contract.ParseTransferShares(txLog)
		log.CheckFatal(err)
		if event.From == core.NULL_ADDR {
			mdl.state.TotalShares = (*core.BigInt)(new(big.Int).Add(mdl.state.TotalShares.Convert(), event.SharesValue))
		}
	}
}

func (mdl *RebaseToken) AfterSyncHook(block int64) {
	// TODO:save
	mdl.save(mdl.lastBlockNum)
	mdl.lastBlockNum = 0
	//
	// mdl.saveUpdatesFromValidatorsJson(block, true)
	//
	mdl.Details = mdl.state.Serialize()
	//
	mdl.SyncAdapter.AfterSyncHook(block)
}
