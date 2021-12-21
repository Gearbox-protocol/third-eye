package repository

import (
	"fmt"
	"github.com/Gearbox-protocol/third-eye/artifacts/dataCompressor/mainnet"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/utils"
	"math/big"
)

func (repo *Repository) AddCreditSessionSnapshot(css *core.CreditSessionSnapshot) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	css.ID = 0
	repo.blocks[css.BlockNum].AddCreditSessionSnapshot(css)
}

func (repo *Repository) ConvertToBalanceWithMask(balances []mainnet.DataTypesTokenBalance, mask *big.Int) *core.JsonBalance {
	jsonBalance := core.JsonBalance{}
	maskInBits := fmt.Sprintf("%b", mask)
	for i, token := range balances {
		tokenAddr := token.Token.Hex()
		if token.Balance.Sign() != 0 {
			jsonBalance[tokenAddr] = &core.BalanceType{
				BI:     (*core.BigInt)(token.Balance),
				F:      utils.GetFloat64Decimal(token.Balance, repo.GetToken(tokenAddr).Decimals),
				Linked: maskInBits[i] == '1',
			}
		}
	}
	return &jsonBalance
}
