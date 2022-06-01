package repository

import (
	"fmt"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/mainnet"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

func (repo *Repository) AddCreditSessionSnapshot(css *schemas.CreditSessionSnapshot) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	css.ID = 0
	repo.SetAndGetBlock(css.BlockNum).AddCreditSessionSnapshot(css)
}

func (repo *Repository) ConvertToBalanceWithMask(balances []mainnet.DataTypesTokenBalance, mask *big.Int) (*core.JsonBalance, error) {
	jsonBalance := core.JsonBalance{}
	maskInBits := fmt.Sprintf("%b", mask)
	maskLen := len(maskInBits)
	for i, token := range balances {
		tokenAddr := token.Token.Hex()
		if token.Balance.Sign() != 0 {
			tokenObj := repo.GetToken(tokenAddr)
			jsonBalance[tokenAddr] = &core.BalanceType{
				BI: (*core.BigInt)(token.Balance),
				F:  utils.GetFloat64Decimal(token.Balance, tokenObj.Decimals),
			}
			if maskLen-i-1 >= 0 {
				jsonBalance[tokenAddr].Linked = maskInBits[maskLen-i-1] == '1'
			}
		}
	}
	return &jsonBalance, nil
}
