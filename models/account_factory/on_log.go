package account_factory

import (
	// "fmt"
	// "math/big"

	// "github.com/Gearbox-protocol/third-eye/artifacts/priceFeed"
	// "github.com/Gearbox-protocol/third-eye/artifacts/yearnPriceFeed"
	// "github.com/Gearbox-protocol/third-eye/models/chainlink_price_feed"
	// "github.com/Gearbox-protocol/third-eye/models/yearn_price_feed"
	// "github.com/Gearbox-protocol/third-eye/utils"
	// "github.com/ethereum/go-ethereum/accounts/abi/bind"
	// "github.com/ethereum/go-ethereum/common"
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
