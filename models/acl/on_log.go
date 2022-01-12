package acl

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/ethereum/go-ethereum/core/types"
)

func (mdl *ACL) OnLog(txLog types.Log) {
	blockNum := int64(txLog.BlockNumber)
	switch txLog.Topics[0] {
	case core.Topic("PausableAdminAdded(address)"):
		pausableAdminAddedEvent, err := mdl.contractETH.ParsePausableAdminAdded(txLog)
		log.CheckFatal(err)
		mdl.Repo.AddDAOOperation(&core.DAOOperation{
			BlockNumber: blockNum,
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    mdl.Address,
			Type:        core.PausableAdminAdded,
			Args:        &core.Json{"newAdmin": pausableAdminAddedEvent.NewAdmin.Hex()},
		})
	case core.Topic("PausableAdminRemoved(address)"):
		pausableAdminRemovedEvent, err := mdl.contractETH.ParsePausableAdminRemoved(txLog)
		log.CheckFatal(err)
		mdl.Repo.AddDAOOperation(&core.DAOOperation{
			BlockNumber: blockNum,
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    mdl.Address,
			Type:        core.PausableAdminRemoved,
			Args:        &core.Json{"admin": pausableAdminRemovedEvent.Admin.Hex()},
		})
	case core.Topic("UnpausableAdminAdded(address)"):
		unpausableAdminAddedEvent, err := mdl.contractETH.ParseUnpausableAdminAdded(txLog)
		log.CheckFatal(err)
		mdl.Repo.AddDAOOperation(&core.DAOOperation{
			BlockNumber: blockNum,
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    mdl.Address,
			Type:        core.UnpausableAdminAdded,
			Args:        &core.Json{"newAdmin": unpausableAdminAddedEvent.NewAdmin.Hex()},
		})
	case core.Topic("UnpausableAdminRemoved(address)"):
		unpausableAdminRemovedEvent, err := mdl.contractETH.ParseUnpausableAdminRemoved(txLog)
		log.CheckFatal(err)
		mdl.Repo.AddDAOOperation(&core.DAOOperation{
			BlockNumber: blockNum,
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    mdl.Address,
			Type:        core.PausableAdminRemoved,
			Args:        &core.Json{"admin": unpausableAdminRemovedEvent.Admin.Hex()},
		})
	case core.Topic("OwnershipTransferred(address,address)"):
		transferEvent, err := mdl.contractETH.ParseOwnershipTransferred(txLog)
		log.CheckFatal(err)
		mdl.Repo.AddDAOOperation(&core.DAOOperation{
			BlockNumber: blockNum,
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    mdl.Address,
			Type:        core.PausableAdminRemoved,
			Args: &core.Json{
				"oldOwner": transferEvent.PreviousOwner.Hex(),
				"newOwner": transferEvent.NewOwner.Hex(),
			},
		})
	}
}
