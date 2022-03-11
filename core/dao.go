package core

type DAOOperation struct {
	LogID       uint   `gorm:"column:log_id;primaryKey"`
	TxHash      string `gorm:"column:tx_hash"`
	BlockNumber int64  `gorm:"column:block_num;primaryKey"`
	Contract    string `gorm:"column:contract"`
	Type        uint   `gorm:"column:type"`
	Args        *Json  `gorm:"column:args"`
}

type TreasuryTransfer struct {
	Amount   *BigInt `gorm:"column:amount"`
	Token    string  `gorm:"column:token"`
	LogID    uint    `gorm:"column:log_id;primaryKey"`
	BlockNum int64   `gorm:"column:block_num;primaryKey"`
}

type BlockDate struct {
	Date      string `gorm:"column:date"`
	BlockNum  int64  `gorm:"column:block_num"`
	Timestamp int64  `gorm:"column:timestamp"`
}

type TreasurySnapshot struct {
	BlockNum    int64         `gorm:"column:block_num"`
	Date        string        `gorm:"column:date_str"`
	PricesInUSD *JsonFloatMap `gorm:"column:prices_in_usd"`
	Balances    *JsonFloatMap `gorm:"column:balances"`
	ValueInUSD  float64       `gorm:"column:value_in_usd"`
}

type TreasurySnapshotModel2 struct {
	BlockNum    int64         `gorm:"column:block_num;primaryKey"`
	Date        string        `gorm:"column:date_str"`
	PricesInUSD *JsonFloatMap `gorm:"column:prices_in_usd"`
	Balances    *JsonFloatMap `gorm:"column:balances"`
	ValueInUSD  float64       `gorm:"column:value_in_usd"`
}

func (TreasurySnapshotModel2) TableName() string {
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
	// for every contract
	Paused
	UnPaused
	EventNewParameters
	///////////
	// v2 events
	///////////
	TokenAllowedV2
	LimitsUpdated
	FeesUpdated
	CreditFacadeUpgraded
	CreditConfiguratorUpgraded
	LTUpdated
)
