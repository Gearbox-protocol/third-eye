package contract_register

import (
	"github.com/Gearbox-protocol/gearscan/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/Gearbox-protocol/gearscan/models/credit_manager"
	"github.com/Gearbox-protocol/gearscan/models/pool"
)

func (mdl *ContractRegister) OnLog(txLog types.Log) {
	blockNum := int64(txLog.BlockNumber)
	switch txLog.Topics[0] {
		case core.Topic("NewPoolAdded(address)"):
			address := common.HexToAddress(txLog.Topics[1].Hex()).Hex()
			obj := pool.NewPool(address, mdl.SyncAdapter.Client, mdl.State.Repo, blockNum)
			mdl.State.Repo.AddSyncAdapter(obj)
		case core.Topic("NewCreditManagerAdded(address)"):
			address := common.HexToAddress(txLog.Topics[1].Hex()).Hex()
			cm := credit_manager.NewCreditManager(address, mdl.SyncAdapter.Client, mdl.State.Repo, blockNum)
			mdl.State.Repo.AddSyncAdapter(cm)
	}
}