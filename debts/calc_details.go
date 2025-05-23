package debts

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/core/schemas/schemas_v3"
	"github.com/Gearbox-protocol/third-eye/ds"
)

type sessionDetailsForCalc struct {
	addr string
	*schemas.CreditSessionSnapshot
	CM            string
	rebaseDetails *schemas.RebaseDetailsForDB
	client        core.ClientI
	// for v3
	forQuotas v3DebtDetails
	version   core.VersionType
}

func (s sessionDetailsForCalc) GetAddr() string {
	return s.addr
}

func (s sessionDetailsForCalc) GetCM() string {
	return s.CM
}
func (s sessionDetailsForCalc) GetBalances() core.DBBalanceFormat {
	// credit session snapshot is not saved to db , so we can use it.
	stETH := core.NULL_ADDR
	if core.GetBaseChainId(s.client) == 1 {
		stETH = core.GetToken(1, "stETH")
	}
	schemas.AdjustRebaseToken(*s.Balances, stETH.Hex(), s.rebaseDetails)
	return *s.Balances
}

func (s sessionDetailsForCalc) GetBorrowedAmount() *big.Int {
	return s.BorrowedAmountBI.Convert()
}
func (s sessionDetailsForCalc) GetCumulativeIndex() *big.Int {
	return s.СumulativeIndexAtOpen.Convert()
}
func (s sessionDetailsForCalc) GetVersion() core.VersionType {
	return s.version
}

func (s sessionDetailsForCalc) GetQuotaCumInterestAndFees() (*big.Int, *big.Int) {
	return s.CumulativeQuotaInterest.Convert(), s.QuotaFees.Convert()
}

type storeForCalc struct {
	inner *DebtEngine
}

func (s storeForCalc) GetToken(token string) (*schemas.Token, error) {
	return s.inner.repo.GetToken(token), nil
}
func (s storeForCalc) GetDecimals(token string) int8 {
	return s.inner.repo.GetToken(token).Decimals
}
func (s storeForCalc) GetPriceOnBlock(cm, token string, version core.VersionType, blockNums ...int64) *big.Int {
	return s.inner.priceHandler.GetLastPrice(cm, token, version)
}

func (s storeForCalc) GetLiqThreshold(ts uint64, cm, token string) *big.Int {
	if ltRamp := s.inner.tokenLTRamp[cm][token]; ltRamp != nil {
		return ltRamp.GetLTForTs(ts)
	}
	return nil
}

type poolDetailsForCalc struct {
	cumIndexAndUToken *ds.CumIndexAndUToken
	forQuotas         v3DebtDetails
}

func (s poolDetailsForCalc) GetUnderlying() string {
	return s.cumIndexAndUToken.Token
}

func (s poolDetailsForCalc) getPool() string {
	return s.cumIndexAndUToken.PoolAddr
}
func (s poolDetailsForCalc) GetPoolQuotaDetails() map[string]*schemas_v3.QuotaDetails {
	return s.forQuotas.poolQuotaDetails[s.getPool()]
}
func (s poolDetailsForCalc) GetCumIndexNow() *big.Int {
	return s.cumIndexAndUToken.CumulativeIndex
}
