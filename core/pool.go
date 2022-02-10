package core

type PoolState struct {
	Address                string  `gorm:"primaryKey" json:"address"`
	UnderlyingToken        string  `gorm:"column:underlying_token" json:"underlyingToken"`
	DieselToken            string  `gorm:"column:diesel_token" json:"dieselToken"`
	IsWETH                 bool    `gorm:"column:is_weth"`
	ExpectedLiquidityLimit *BigInt `gorm:"column:expected_liq_limit"`
	WithdrawFee            *BigInt `gorm:"column:withdraw_fee"`
}

func (PoolState) TableName() string {
	return "pools"
}

type UTokenAndPool struct {
	Pool   string
	UToken string
}
type PoolStat struct {
	ID                       int64   `gorm:"primaryKey" json:"-"`
	BlockNum                 int64   `gorm:"column:block_num" json:"blockNum"`
	Address                  string  `gorm:"column:pool" json:"pool"`
	UniqueUsers              int     `gorm:"column:unique_users" json:"uniqueUsers"`
	DepositAPY               float64 `gorm:"column:deposit_apy" json:"depositAPY"`
	DepositAPYBI             *BigInt `gorm:"column:deposit_apy_bi" json:"depositAPYBI"`
	BorrowAPY                float64 `gorm:"column:borrow_apy" json:"borrowAPY"`
	BorrowAPYBI              *BigInt `gorm:"column:borrow_apy_bi" json:"borrowAPYBI"`
	ExpectedLiquidity        float64 `gorm:"column:expected_liquidity" json:"expectedLiquidity"`
	ExpectedLiquidityBI      *BigInt `gorm:"column:expected_liquidity_bi" json:"expectedLiquidityBI"`
	ExpectedLiquidityLimitBI *BigInt `gorm:"column:expected_liquidity_limit_bi" json:"expectedLiquidityLimitBI"`
	AvailableLiquidity       float64 `gorm:"column:available_liquidity" json:"availableLiquidity"`
	AvailableLiquidityBI     *BigInt `gorm:"column:available_liquidity_bi" json:"availableLiquidityBI"`
	TotalBorrowed            float64 `gorm:"column:total_borrowed" json:"totalBorrowed"`
	TotalBorrowedBI          *BigInt `gorm:"column:total_borrowed_bi" json:"totalBorrowedBI"`
	TotalProfit              float64 `gorm:"column:total_profit" json:"totalProfit"`
	TotalProfitBI            *BigInt `gorm:"column:total_profit_bi" json:"totalProfitBI"`
	TotalLosses              float64 `gorm:"column:total_losses" json:"totalLosses"`
	TotalLossesBI            *BigInt `gorm:"column:total_losses_bi" json:"totalLossesBI"`
	DieselRate               float64 `gorm:"column:diesel_rate" json:"dieselRate"`
	DieselRateBI             *BigInt `gorm:"column:diesel_rate_bi" json:"dieselRateBI"`
	WithdrawFee              int     `gorm:"column:withdraw_fee" json:"withdrawFee"`
	CumulativeIndexRAY       *BigInt `gorm:"column:cumulative_index_ray" json:"cumulativeIndexRAY"`
}

type PoolInterestData struct {
	BorrowAPYBI        *BigInt `gorm:"column:borrow_apy_bi"`
	CumulativeIndexRAY *BigInt `gorm:"column:cumulative_index_ray"`
	BlockNum           int64   `gorm:"column:block_num"`
	Address            string  `gorm:"column:pool"`
	Timestamp          uint64  `gorm:"column:timestamp"`
}

type PoolLedger struct {
	Id          int64   `gorm:"primaryKey;autoincrement:true" json:"-"`
	BlockNumber int64   `gorm:"column:block_num" json:"blockNum"`
	Pool        string  `gorm:"column:pool" json:"pool"`
	User        string  `gorm:"column:user_address" json:"user"`
	TxHash      string  `gorm:"column:tx_hash" json:"txHash"`
	SessionId   string  `gorm:"column:session_id" json:"sessionId"`
	LogId       uint    `gorm:"column:log_id" json:"logId"`
	Event       string  `gorm:"column:event" json:"event"`
	AmountBI    *BigInt `gorm:"column:amount_bi" json:"-"`
	Amount      float64 `gorm:"column:amount" json:"amount"`
}

func (PoolLedger) TableName() string {
	return "pool_ledger"
}
