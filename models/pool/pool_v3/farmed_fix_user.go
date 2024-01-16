package pool_v3

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"
)

type UpdatePoolLedger struct {
	Zapper   string
	Pool     string
	User     string
	BlockNum int64
	TxHash   string
	Amount   *big.Int
	Type     string
}

func (mdl *Poolv3) updateFarmedv3(txLog types.Log) {
	from := common.BytesToAddress(txLog.Topics[1][:]).Hex()
	to := common.BytesToAddress(txLog.Topics[2][:]).Hex()
	amount := new(big.Int).SetBytes(txLog.Data[:])
	txHash := txLog.TxHash.Hex()
	blockNum := int64(txLog.BlockNumber)

	if from == core.NULL_ADDR.Hex() || to == core.NULL_ADDR.Hex() {
		return
	}

	if from == mdl.getZapPoolv2() {
		mdl.updatesForPoolv2 = append(mdl.updatesForPoolv2, UpdatePoolLedger{
			Zapper:   from,
			User:     to,
			Pool:     mdl.getPoolv2(),
			BlockNum: blockNum,
			TxHash:   txHash,
			Type:     "AddLiquidity",
		})
	}

	if from == mdl.getZapUnderlying() || from == mdl.getZapPoolv2() { // usdc-farmedUSDCv3, dUSDC-farmedUSDCv3
		if txHash != mdl.addLiquidityEvent.TxHash ||
			blockNum != mdl.addLiquidityEvent.BlockNumber ||
			amount.Cmp(mdl.addLiquidityEvent.AmountBI.Convert()) != 0 ||
			from != mdl.addLiquidityEvent.User {
			log.Fatal(utils.ToJson(mdl.addLiquidityEvent), "addLiquidityEvent", "txHash", txHash, "blockNum", blockNum, "amount", amount)
		}
		mdl.addLiquidityEvent.User = to
		mdl.Repo.AddPoolLedger(mdl.addLiquidityEvent)
		mdl.addLiquidityEvent = nil
	}

	if to == mdl.getZapPoolv2() || to == mdl.getZapPoolv2() {
		mdl.removeLiqUpdate = &UpdatePoolLedger{
			Zapper:   to,
			User:     from,
			Pool:     mdl.Address,
			BlockNum: blockNum,
			TxHash:   txHash,
			Amount:   amount,
		}
	}
	if to == mdl.getZapPoolv2() {
		mdl.updatesForPoolv2 = append(mdl.updatesForPoolv2, UpdatePoolLedger{
			Zapper:   to,
			User:     from,
			Pool:     mdl.getPoolv2(),
			BlockNum: blockNum,
			TxHash:   txHash,
			Type:     "RemoveLiquidity",
		})
	}
}

func (mdl *Poolv3) changeAddressOnAddLiq(event *schemas.PoolLedger) {
	if mdl.removeLiqUpdate != nil {
		log.Fatal(utils.ToJson(mdl.removeLiqUpdate), "removeLiqUpdate is not nil on add Liq", utils.ToJson(event))
	}
	if mdl.addLiquidityEvent != nil {
		log.Fatal(utils.ToJson(mdl.addLiquidityEvent), "addLiquidityEvent is not nil", "pool", mdl.Address, "event", event)
	}
	mdl.addLiquidityEvent = event
}
func (mdl *Poolv3) changeAddressOnRemoveLiq(event *schemas.PoolLedger) {
	if mdl.removeLiqUpdate == nil {
		log.Fatal(utils.ToJson(event))
	}
	txHash := mdl.removeLiqUpdate.TxHash
	blockNum := mdl.removeLiqUpdate.BlockNum
	amount := mdl.removeLiqUpdate.Amount
	if txHash != event.TxHash ||
		blockNum != event.BlockNumber ||
		amount.Cmp(event.AmountBI.Convert()) != 0 ||
		mdl.removeLiqUpdate.Zapper != event.User {
		log.Fatal(utils.ToJson(mdl.removeLiqUpdate), "removeLiqUpdate", "txHash", txHash, "blockNum", blockNum, "amount", amount)
	}
	if mdl.addLiquidityEvent != nil {
		log.Fatal(utils.ToJson(mdl.addLiquidityEvent), "addLiquidityEvent is not nil", "pool", mdl.Address, "event", event)
	}
	//
	//
	event.User = mdl.removeLiqUpdate.User
	mdl.Repo.AddPoolLedger(mdl.addLiquidityEvent)
	mdl.removeLiqUpdate = nil
}

func (mdl *Poolv3) UpdatePoolv2Ledger(tx *gorm.DB) {
	for _, update := range mdl.updatesForPoolv2 {
		if update.Type == "" {
			continue
		}
		x := tx.Exec(`UPDATE pool_ledger set user_address=? WHERE block_num=? AND type = ? AND user in (?, ?)`,
			update.User, update.BlockNum, update.Type, update.User, update.Zapper)
		if x.RowsAffected != 1 {
			log.Fatal("Can't update for", utils.ToJson(update))
		}
		log.CheckFatal(x.Error)
	}
}
