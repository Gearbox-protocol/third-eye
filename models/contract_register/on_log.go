package contract_register

import (
	"log"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager/cm_v1"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager/cm_v2"
	"github.com/Gearbox-protocol/third-eye/models/pool"
)

func (mdl *ContractRegister) OnLog(txLog types.Log) {
	blockNum := int64(txLog.BlockNumber)
	switch txLog.Topics[0] {
	case core.Topic("NewPoolAdded(address)"):
		address := common.HexToAddress(txLog.Topics[1].Hex()).Hex()
		obj := pool.NewPool(address, mdl.SyncAdapter.Client, mdl.Repo, blockNum)
		mdl.Repo.AddSyncAdapter(obj)
	case core.Topic("NewCreditManagerAdded(address)"):
		address := common.HexToAddress(txLog.Topics[1].Hex()).Hex()
		cm := NewCM(address, mdl.SyncAdapter.Client, mdl.Repo, blockNum)
		mdl.Repo.AddSyncAdapter(cm)
	}
}

func NewCM(addr string, client core.ClientI, repo ds.RepositoryI, blockNum int64) ds.SyncAdapterI {
	version := core.FetchVersion(addr, blockNum, client)
	switch version {
	case 1:
		return cm_v1.NewCMv1(addr, client, repo, blockNum)
	case 2:
		return cm_v2.NewCMv2(addr, client, repo, blockNum)
	default:
		log.Fatalf("Version(%d) of cm can't be created.", version)
	}
	return nil
}
