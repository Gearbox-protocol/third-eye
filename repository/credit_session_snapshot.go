package repository

import (
	"github.com/Gearbox-protocol/third-eye/artifacts/dataCompressor/mainnet"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/utils"
)

func (repo *Repository) AddCreditSessionSnapshot(css *core.CreditSessionSnapshot) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	css.ID = 0
	repo.blocks[css.BlockNum].AddCreditSessionSnapshot(css)
}

func (repo *Repository) ConvertToBalance(balances []mainnet.DataTypesTokenBalance) *core.JsonBalance {
	jsonBalance := core.JsonBalance{}
	for _, token := range balances {
		tokenAddr := token.Token.Hex()
		if token.Balance.Sign() != 0 {
			jsonBalance[tokenAddr] = &core.BalanceType{
				BI: (*core.BigInt)(token.Balance),
				F:  utils.GetFloat64Decimal(token.Balance, repo.GetToken(tokenAddr).Decimals),
			}
		}
	}
	return &jsonBalance
}
