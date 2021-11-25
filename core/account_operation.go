package core

type AccountOperation struct {
	// Input string
	ID int64 `gorm:"primaryKey;autoIncrement:true"`
	// ordering data
	TxHash      string `gorm:"column:tx_hash"`
	BlockNumber int64  `gorm:"column:block_num"`
	LogId       uint   `gorm:"column:log_id`
	// owner/account data
	Borrower  string
	SessionId string `gorm:"column:session_id"`
	// application
	Dapp string `gorm:"column:dapp"`
	// call/events data
	AdapterCall bool   `gorm:"column:adapter_call"`
	Action      string `gorm:"column:action"`
	Args        string `gorm:"column:args"`
	Transfers   string `gorm:"column:transfers"`
	// extras
	Depth     uint8 `gorm:"column:depth"`
	Timestamp int64 `gorm:"column:timestamp"`
}

func (AccountOperation) TableName() string {
	return "account_operations"
}

const (
	EventType = iota
	CallType
)
