package gear_token

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

func (mdl *GearToken) OnLog(txLog types.Log) {
	// blockNum := int64(txLog.BlockNumber)
	switch txLog.Topics[0] {
	case core.Topic("Transfer(address,address,uint256)"):
		transferEvent, err := mdl.contractETH.ParseTransfer(txLog)
		log.CheckFatal(err)
		mdl.UpdateStateBalances(transferEvent.From.Hex(), transferEvent.To.Hex(), transferEvent.Value)
	}
}

func (mdl *GearToken) UpdateStateBalances(from, to string, value *big.Int) {
	fromBalance := mdl.State[from]
	toBalance := mdl.State[to]
	if fromBalance == nil {
		fromBalance = &core.GearBalance{
			Balance: (*core.BigInt)(big.NewInt(0)),
			User:    from,
		}
	}
	if toBalance == nil {
		toBalance = &core.GearBalance{
			Balance: (*core.BigInt)(big.NewInt(0)),
			User:    to,
		}
	}
	fromBalance.Balance = core.AddCoreAndInt(fromBalance.Balance, new(big.Int).Neg(value))
	fromBalance.Updated = true
	mdl.State[from] = fromBalance
	toBalance.Balance = core.AddCoreAndInt(toBalance.Balance, value)
	toBalance.Updated = true
	mdl.State[to] = toBalance
}
