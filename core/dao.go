package core

type DAOOperation struct {
	LogID       uint   `gorm:"column:log_id;primaryKey"`
	TxHash      string `gorm:"column:tx_hash"`
	BlockNumber int64  `gorm:"column:block_num;primaryKey"`
	Contract    string `gorm:"column:contract"`
	Type        uint   `gorm:"column:"`
	Args        *Json  `gorm:"column:args"`
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
