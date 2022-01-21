package account_factory

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/common"
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
	case core.Topic("NewCreditAccount(address)"):
		accountAddr := common.HexToAddress(txLog.Topics[1].Hex()).Hex()
		mdl.Repo.AddAccountAddr(accountAddr)
	}
}
