package account_factory

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/ethereum/go-ethereum/core/types"
)

func (mdl *AccountFactory) OnLog(txLog types.Log) {
	blockNum := int64(txLog.BlockNumber)
	switch txLog.Topics[0] {
	case core.Topic("TakeForever(address,address)"):
		takeForeverEvent, err := mdl.contractETH.ParseTakeForever(txLog)
		log.CheckFatal(err)
		mdl.Repo.AddDAOOperation(&core.DAOOperation{
			BlockNumber: blockNum,
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    mdl.Address,
			Type:        core.TakeForever,
			Args: &core.Json{
				"creditAccount": takeForeverEvent.CreditAccount.Hex(),
				"to":            takeForeverEvent.To.Hex(),
			},
		})
	}
}
