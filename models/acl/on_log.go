package acl

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/core/types"
)

func (mdl *ACL) OnLog(txLog types.Log) {
	blockNum := int64(txLog.BlockNumber)
	switch txLog.Topics[0] {
	case core.Topic("PausableAdminAdded(address)"):
		pausableAdminAddedEvent, err := mdl.contractETH.ParsePausableAdminAdded(txLog)
		log.CheckFatal(err)
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: blockNum,
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    mdl.Address,
			Type:        schemas.PausableAdminAdded,
			Args:        &core.Json{"admin": pausableAdminAddedEvent.NewAdmin.Hex()},
		})
	case core.Topic("PausableAdminRemoved(address)"):
		pausableAdminRemovedEvent, err := mdl.contractETH.ParsePausableAdminRemoved(txLog)
		log.CheckFatal(err)
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: blockNum,
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    mdl.Address,
			Type:        schemas.PausableAdminRemoved,
			Args:        &core.Json{"admin": pausableAdminRemovedEvent.Admin.Hex()},
		})
	case core.Topic("UnpausableAdminAdded(address)"):
		unpausableAdminAddedEvent, err := mdl.contractETH.ParseUnpausableAdminAdded(txLog)
		log.CheckFatal(err)
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: blockNum,
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    mdl.Address,
			Type:        schemas.UnpausableAdminAdded,
			Args:        &core.Json{"admin": unpausableAdminAddedEvent.NewAdmin.Hex()},
		})
	case core.Topic("UnpausableAdminRemoved(address)"):
		unpausableAdminRemovedEvent, err := mdl.contractETH.ParseUnpausableAdminRemoved(txLog)
		log.CheckFatal(err)
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: blockNum,
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    mdl.Address,
			Type:        schemas.UnpausableAdminRemoved,
			Args:        &core.Json{"admin": unpausableAdminRemovedEvent.Admin.Hex()},
		})
	case core.Topic("OwnershipTransferred(address,address)"):
		transferEvent, err := mdl.contractETH.ParseOwnershipTransferred(txLog)
		log.CheckFatal(err)
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: blockNum,
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    mdl.Address,
			Type:        schemas.OwnershipTransferred,
			Args: &core.Json{
				"oldOwner": transferEvent.PreviousOwner.Hex(),
				"newOwner": transferEvent.NewOwner.Hex(),
			},
		})
	}
}
