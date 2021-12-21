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

func (repo *Repository) ConvertToBalanceWithMask(balances []mainnet.DataTypesTokenBalance, mask *big.Int) (*core.JsonBalance, error) {
	jsonBalance := core.JsonBalance{}
	maskInBits := fmt.Sprintf("%b", mask)
	maskLen := len(maskInBits)
	for i, token := range balances {
		tokenAddr := token.Token.Hex()
		if token.Balance.Sign() != 0 {
			tokenObj, err := repo.getTokenWithError(tokenAddr)
			if err != nil {
				return nil, err
			}
			jsonBalance[tokenAddr] = &core.BalanceType{
				BI:     (*core.BigInt)(token.Balance),
				F:      utils.GetFloat64Decimal(token.Balance, tokenObj.Decimals),
			}
			if maskLen-i-1 >= 0 {
				jsonBalance[tokenAddr].Linked = maskInBits[maskLen-i-1] == '1'
			}
		}
	}
	return &jsonBalance, nil
}
