package debts

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
)

type sessionDetailsForCalc struct {
	*schemas.CreditSessionSnapshot
	CM            string
	rebaseDetails *schemas.RebaseDetailsForDB
	stETH         string
}

func (s sessionDetailsForCalc) GetCM() string {
	return s.CM
}
func (s sessionDetailsForCalc) GetBalances() map[string]core.BalanceType {
	balances := s.Balances.ToBalanceType()
	schemas.AdjustRebaseToken(balances, s.stETH, s.rebaseDetails)
	return balances
}

func (s sessionDetailsForCalc) GetBorrowedAmount() *big.Int {
	return s.BorrowedAmountBI.Convert()
}
func (s sessionDetailsForCalc) GetCumulativeIndex() *big.Int {
	return s.СumulativeIndexAtOpen.Convert()
}

type storeForCalc struct {
	inner *DebtEngine
}

func (s storeForCalc) GetToken(token string) *schemas.Token {
	return s.inner.repo.GetToken(token)
}
func (s storeForCalc) GetPrices(token string, version core.VersionType, blockNums ...int64) *big.Int {
	return s.inner.GetTokenLastPrice(token, version)
}

func (s storeForCalc) GetLiqThreshold(cm, token string) *big.Int {
	return s.inner.allowedTokensThreshold[cm][token].Convert()
}
