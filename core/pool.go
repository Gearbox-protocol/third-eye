package core

type PoolState struct {
	Address         string `gorm:"primaryKey"`
	UnderlyingToken string `gorm:"column:underlying_token"`
	DieselToken     string `gorm:"column:diesel_token"`
	IsWETH          bool   `gorm:"column:is_weth"`
}

func (PoolState) TableName() string {
	return "pools"
}

type PoolStat struct {
	ID                   int64   `gorm:"primaryKey"`
	BlockNum             int64   `gorm:"column:block_num"`
	Address              string  `gorm:"column:pool"`
	UniqueUsers          int     `gorm:"column:unique_users"`
	DepositAPY           float64 `gorm:"column:deposit_apy"`
	DepositAPYBI         *BigInt `gorm:"column:deposit_apy_bi"`
	BorrowAPY            float64 `gorm:"column:borrow_apy"`
	BorrowAPYBI          *BigInt `gorm:"column:borrow_apy_bi"`
	ExpectedLiquidity    float64 `gorm:"column:expected_liquidity"`
	ExpectedLiquidityBI  *BigInt `gorm:"column:expected_liquidity_bi"`
	AvailableLiquidity   float64 `gorm:"column:available_liquidity"`
	AvailableLiquidityBI *BigInt `gorm:"column:available_liquidity_bi"`
	TotalBorrowed        float64 `gorm:"column:total_borrowed"`
	TotalBorrowedBI      *BigInt `gorm:"column:total_borrowed_bi"`
	TotalProfit          float64 `gorm:"column:total_profit"`
	TotalProfitBI        *BigInt `gorm:"column:total_profit_bi"`
	TotalLosses          float64 `gorm:"column:total_losses"`
	TotalLossesBI        *BigInt `gorm:"column:total_losses_bi"`
	DieselRate           float64 `gorm:"column:diesel_rate"`
	DieselRateBI         *BigInt `gorm:"column:diesel_rate_bi"`
	WithdrawFee          int     `gorm:"column:withdraw_fee"`
	CumulativeIndexRAY   *BigInt `gorm:"column:cumulative_index_ray"`
}

type PoolInterestData struct {
	BorrowAPYBI        *BigInt `gorm:"column:borrow_apy_bi"`
	CumulativeIndexRAY *BigInt `gorm:"column:cumulative_index_ray"`
	BlockNum           int64   `gorm:"column:block_num"`
	Address            string  `gorm:"column:pool"`
	Timestamp          uint64  `gorm:"column:timestamp"`
}

type PoolLedger struct {
	Id          int64   `gorm:"primaryKey;autoincrement:true"`
	BlockNumber int64   `gorm:"column:block_num"`
	Pool        string  `gorm:"column:pool"`
	User        string  `gorm:"column:address"`
	LogId       int64   `gorm:"column:log_id"`
	Event       string  `gorm:"column:event"`
	Liquidity   *BigInt `gorm:"column:liquidity"`
}

func (PoolLedger) TableName() string {
	return "pool_ledger"
}
