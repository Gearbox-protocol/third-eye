package core

type DAOOperation struct {
	LogID       uint   `gorm:"column:log_id;primaryKey"`
	TxHash      string `gorm:"column:tx_hash"`
	BlockNumber int64  `gorm:"column:block_num;primaryKey"`
	Contract    string `gorm:"column:contract"`
	Type        uint   `gorm:"column:"`
	Args        *Json  `gorm:"column:args"`
}

type TreasuryTransfer struct {
	Amount   *BigInt `gorm:"column:amount"`
	Token    string  `gorm:"column:token"`
	LogID    uint    `gorm:"column:log_id;primaryKey"`
	BlockNum int64   `gorm:"column:block_num;primaryKey"`
}

type BlockDate struct {
	BlockNum  int64 `gorm:"column:block_num"`
	Timestamp int64 `gorm:"column:timestamp"`
}

type TreasurySnapshot struct {
	Date        string         `gorm:"column:date_str"`
	Timestamp   int64          `gorm:"primaryKey;column:timestamp"`
	PricesInUSD *JsonBigIntMap `gorm:"column:prices_in_usd"`
	Balances    *JsonBigIntMap `gorm:"column:balances"`
	ValueInUSD  float64        `gorm:"column:value_in_usd"`
}

func (TreasurySnapshot) TableName() string {
	return "treasury_snapshots"
}

const (
	// credit filter
	TokenAllowed = iota
	TokenForbidden
	ContractAllowed
	ContractForbidden
	NewFastCheckParameters
	TransferPluginAllowed
	PriceOracleUpdated
	// pools
	NewInterestRateModel
	NewCreditManagerConnected
	NewExpectedLiquidityLimit
	BorrowForbidden
	NewWithdrawFee
	// price oracle
	NewPriceFeed
	// account factory
	TakeForever
	// acl
	PausableAdminAdded
	PausableAdminRemoved
	UnpausableAdminAdded
	UnpausableAdminRemoved
	OwnershipTransferred
)
