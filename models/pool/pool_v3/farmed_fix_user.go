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
	Shares   *big.Int
	Type     string
}

func (mdl *Poolv3) updateFarmedv3(txLog types.Log) {
	from := common.BytesToAddress(txLog.Topics[1][:]).Hex()
	to := common.BytesToAddress(txLog.Topics[2][:]).Hex()
	shares := new(big.Int).SetBytes(txLog.Data[:])
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
			Type:     "RemoveLiquidity",
		})
	}
	if mdl.checkIfZapAddr(from) { // usdc-farmedUSDCv3, dUSDC-farmedUSDCv3, ETH-farmedETHv3
		if txHash != mdl.addLiquidityEvent.TxHash ||
			blockNum != mdl.addLiquidityEvent.BlockNumber ||
			shares.Cmp(mdl.addLiquidityEvent.SharesBI.Convert()) != 0 ||
			from != mdl.addLiquidityEvent.User {
			log.Fatal(utils.ToJson(mdl.addLiquidityEvent), "addLiquidityEvent", "txHash", txHash, "blockNum", blockNum, "shares", shares)
		}
		mdl.addLiquidityEvent.User = to
		mdl.Repo.AddPoolLedger(mdl.addLiquidityEvent)
		mdl.addLiquidityEvent = nil
	}

	if mdl.checkIfZapAddr(to) {
		mdl.removeLiqUpdate = &UpdatePoolLedger{
			Zapper:   to,
			User:     from,
			Pool:     mdl.Address,
			BlockNum: blockNum,
			TxHash:   txHash,
			Shares:   shares,
		}
	}
	if to == mdl.getZapPoolv2() {
		mdl.updatesForPoolv2 = append(mdl.updatesForPoolv2, UpdatePoolLedger{
			Zapper:   to,
			User:     from,
			Pool:     mdl.getPoolv2(),
			BlockNum: blockNum,
			TxHash:   txHash,
			Type:     "AddLiquidity",
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
	shares := mdl.removeLiqUpdate.Shares
	if txHash != event.TxHash ||
		blockNum != event.BlockNumber ||
		shares.Cmp(event.SharesBI.Convert()) != 0 ||
		mdl.removeLiqUpdate.Zapper != event.User {
		log.Fatal(utils.ToJson(mdl.removeLiqUpdate), "removeLiqUpdate", "txHash", event.TxHash, "blockNum", event.BlockNumber, "shares", event.SharesBI, "user", event.User)
	}
	if mdl.addLiquidityEvent != nil {
		log.Fatal(utils.ToJson(mdl.addLiquidityEvent), "addLiquidityEvent is not nil", "pool", mdl.Address, "event", event)
	}
	//
	//
	event.User = mdl.removeLiqUpdate.User
	mdl.Repo.AddPoolLedger(event)
	mdl.removeLiqUpdate = nil
}

func (mdl *Poolv3) UpdatePoolv2Ledger(tx *gorm.DB) {
	for _, update := range mdl.updatesForPoolv2 {
		if update.Type == "" {
			continue
		}
		// log.Info(utils.ToJson(update))
		x := tx.Exec(`UPDATE pool_ledger set user_address=? WHERE block_num=? AND event = ? AND tx_hash=? AND user_address in (?, ?)`,
			update.User, update.BlockNum, update.Type, update.TxHash, update.User, update.Zapper)
		log.CheckFatal(x.Error)
		//

		if x.RowsAffected != 1 {
			log.Infof("SEELCT * from pool_ledger WHERE block_num=%d AND event = '%s' AND tx_hash='%s' AND user_address in ('%s', '%s')", update.BlockNum, update.Type, update.TxHash, update.User, update.Zapper)
			log.Fatal("Can't update for", utils.ToJson(update))
		}
	}
}
