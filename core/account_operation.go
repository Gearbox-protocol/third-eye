package core

type AccountOperation struct {
	// Input string
	ID int64 `gorm:"primaryKey;autoIncrement:true" json:"-"`
	// ordering data
	TxHash      string `gorm:"column:tx_hash" json:"txHash"`
	BlockNumber int64  `gorm:"column:block_num" json:"blockNum"`
	LogId       uint   `gorm:"column:log_id" json:"logId"`
	// owner/account data
	Borrower  string `json:"borrower"`
	SessionId string `gorm:"column:session_id" json:"sessionId"`
	// application
	Dapp string `gorm:"column:dapp" json:"dapp"`
	// call/events data
	AdapterCall bool       `gorm:"column:adapter_call" json:"adapterCall"`
	Action      string     `gorm:"column:action" json:"action"`
	Args        *Json      `gorm:"column:args" json:"args"`
	Transfers   *Transfers `gorm:"column:transfers" json:"transfers"`
	// extras
	Depth      uint8               `gorm:"column:depth" json:"depth"`
	MainAction *int64              `gorm:"column:main_action"`
	MultiCall  []*AccountOperation `gorm:"foreignkey:MainAction" json:"multicalls"`
}

func (AccountOperation) TableName() string {
	return "account_operations"
}

const (
	EventType = iota
	CallType
)

type AccountOperationState struct {
	ID               int64        `gorm:"primaryKey;autoincrement:true" json:"id"`
	BlockNum         int64        `gorm:"column:block_num"`
	LogId            int64        `gorm:"column:log_id"`
	SessionId        string       `gorm:"column:session_id"`
	BorrowedAmountBI *BigInt      `gorm:"column:borrowed_amount_bi"`
	BorrowedAmount   float64      `gorm:"column:borrowed_amount"`
	Balances         *JsonBalance `gorm:"column:balances"`
}

func (AccountOperationState) TableName() string {
	return "account_operation_states"
}
