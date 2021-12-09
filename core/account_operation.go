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
	AdapterCall bool      `gorm:"column:adapter_call"`
	Action      string    `gorm:"column:action"`
	Args        string    `gorm:"column:args"`
	Transfers   Transfers `gorm:"column:transfers"`
	// extras
	Depth uint8 `gorm:"column:depth"`
}

func (AccountOperation) TableName() string {
	return "account_operations"
}

const (
	EventType = iota
	CallType
)

type AccountOperationState struct {
	ID               int64       `gorm:"primaryKey;autoincrement:true" json:"id"`
	BlockNum         int64       `gorm:"column:block_num"`
	LogId            int64       `gorm:"column:log_id"`
	SessionId        string      `gorm:"column:session_id"`
	BorrowedAmountBI *BigInt     `gorm:"column:borrowed_amount_bi"`
	BorrowedAmount   float64     `gorm:"column:borrowed_amount"`
	Balances         JsonBalance `gorm:"column:balances"`
}

func (AccountOperationState) TableName() string {
	return "account_operation_states"
}
