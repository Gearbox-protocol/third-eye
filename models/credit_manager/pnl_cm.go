package credit_manager

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
)

type PnlCM struct {
	curPnl map[int64]*schemas.PnlOnRepay
}

func NewPnlCM() *PnlCM {
	return &PnlCM{curPnl: map[int64]*schemas.PnlOnRepay{}}
}
func (mdl *PnlCM) Get(blockNum int64) *schemas.PnlOnRepay {
	pnlOnRepay := mdl.curPnl[blockNum]
	delete(mdl.curPnl, blockNum)
	return pnlOnRepay
}

func (mdl *PnlCM) Set(pnl *schemas.PnlOnRepay) {
	oldPnl := mdl.curPnl[pnl.BlockNum]
	if oldPnl != nil {
		if oldPnl.BlockNum != pnl.BlockNum {
			log.Fatalf("Pool repay event has different oldPnl:%v pnl:%v for same block",
				oldPnl, pnl)
		}
		pnl.Profit = new(big.Int).Add(oldPnl.Profit, pnl.Profit)
		pnl.Loss = new(big.Int).Add(oldPnl.Loss, pnl.Loss)
		pnl.BorrowedAmount = new(big.Int).Add(oldPnl.BorrowedAmount, pnl.BorrowedAmount)
	}
	mdl.curPnl[pnl.BlockNum] = pnl
}
