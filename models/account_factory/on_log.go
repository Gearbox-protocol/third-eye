package account_factory

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/common"
)

func (mdl *AccountFactory) OnLog(txLog types.Log) {
	blockNum := int64(txLog.BlockNumber)
	switch txLog.Topics[0] {
	case core.Topic("TakeForever(address,address)"):
		takeForeverEvent, err := mdl.contractETH.ParseTakeForever(txLog)
		log.CheckFatal(err)
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: blockNum,
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    mdl.Address,
			Type:        schemas.TakeForever,
			Args: &core.Json{
				"creditAccount": takeForeverEvent.CreditAccount.Hex(),
				"to":            takeForeverEvent.To.Hex(),
			},
		})
	case core.Topic("NewCreditAccount(address)"):
		accountAddr := common.HexToAddress(txLog.Topics[1].Hex()).Hex()
		mdl.Repo.AddAccountAddr(accountAddr)
	case core.Topic("DeployCreditAccount(address,address)"):
		accountAddr := common.HexToAddress(txLog.Topics[1].Hex()).Hex()
		mdl.Repo.AddAccountAddr(accountAddr)
	}
}
