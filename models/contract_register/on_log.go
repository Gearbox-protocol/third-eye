package contract_register

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/Gearbox-protocol/third-eye/models/credit_manager"
	"github.com/Gearbox-protocol/third-eye/models/pool"
)

func (mdl *ContractRegister) OnLog(txLog types.Log) {
	blockNum := int64(txLog.BlockNumber)
	if core.GetChainId(mdl.Client) != 1337 && (core.Topic("NewPoolAdded(address)") == txLog.Topics[0] ||
		core.Topic("NewCreditManagerAdded(address)") == txLog.Topics[0]) {

		address := common.HexToAddress(txLog.Topics[1].Hex())
		version := core.FetchVersionOptimized(address, 0, mdl.Client)
		log.Info(version)
		if version.MoreThanEq(core.NewVersion(300)) {
			return
		}
	}
	switch txLog.Topics[0] {
	case core.Topic("NewPoolAdded(address)"):
		address := common.HexToAddress(txLog.Topics[1].Hex()).Hex()
		obj := pool.NewPool(address, mdl.SyncAdapter.Client, mdl.Repo, blockNum)
		mdl.Repo.AddSyncAdapter(obj)
	case core.Topic("NewCreditManagerAdded(address)"):
		address := common.HexToAddress(txLog.Topics[1].Hex()).Hex()
		cm := credit_manager.NewCreditManager(address, mdl.SyncAdapter.Client, mdl.Repo, blockNum)
		mdl.Repo.AddSyncAdapter(cm)
	}
}
