package contract_register

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager/cm_v1"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager/cm_v2"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager/cm_v3"
	"github.com/Gearbox-protocol/third-eye/models/pool/pool_v2"
	"github.com/Gearbox-protocol/third-eye/models/pool/pool_v3"
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
		// log.Info("new pool", address)
		obj := NewPool(address, mdl.SyncAdapter.Client, mdl.Repo, blockNum)
		mdl.Repo.AddSyncAdapter(obj)
	case core.Topic("NewCreditManagerAdded(address)"):
		address := common.HexToAddress(txLog.Topics[1].Hex()).Hex()
		// log.Info("new cm", address)
		cm := NewCM(address, mdl.SyncAdapter.Client, mdl.Repo, blockNum)
		mdl.Repo.AddSyncAdapter(cm)
	}
}

func NewCM(addr string, client core.ClientI, repo ds.RepositoryI, blockNum int64) ds.SyncAdapterI {
	version := core.FetchVersion(addr, blockNum, client)
	switch version {
	case core.NewVersion(1):
		return cm_v1.NewCMv1(addr, client, repo, blockNum)
	case core.NewVersion(2):
		return cm_v2.NewCMv2(addr, client, repo, blockNum)
	case core.NewVersion(300):
		return cm_v3.NewCMv3(addr, client, repo, blockNum)
	default:
		log.Fatalf("Version(%d) of cm can't be created.", version)
	}
	return nil
}

func NewPool(addr string, client core.ClientI, repo ds.RepositoryI, blockNum int64) ds.SyncAdapterI {
	version := core.FetchVersion(addr, blockNum, client)
	switch version {
	case core.NewVersion(1), core.NewVersion(2):
		return pool_v2.NewPool(addr, client, repo, blockNum)
	case core.NewVersion(300):
		return pool_v3.NewPool(addr, client, repo, blockNum)
	default:
		log.Fatalf("Version(%d) of pool can't be created.", version)
	}
	return nil
}
