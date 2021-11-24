package core

import "github.com/ethereum/go-ethereum/core/types"

type (
	Block struct {
		BlockNumber         int64                     `json:"id" gorm:"primaryKey;column:id"` // Block Number
		Timestamp           int64                   `json:"timestamp" gorm:"column:timestamp"`
		Hash                string                  `json:"hash" gorm:"-"`
		TransactionsCount   int                     `json:"transactions_count" gorm:"-"`
		Transactions        []Transaction           `json:"transactions" gorm:"-"`
		Deployments         map[string]Transaction  `json:"deployments" gorm:"-"`
		Logs                []types.Log             `json:"logs" gorm:"-"`
		// PoolStat            []PoolStat              `gorm:"foreignKey:BlockNum"`
		// CreditManagerStat   []CreditManagerStat     `gorm:"foreignKey:BlockNum"`
		// Operations          []Operation             `gorm:"foreignKey:BlockNum"`
		// PriceFeed           []PriceItem             `gorm:"foreignKey:BlockNum"`
		// CreditOperation     []CreditOperation       `gorm:"foreignKey:BlockNum"`
		// CSSnapshots         []CreditSessionSnapshot `gorm:"foreignKey:BlockNum"`
		// AccountOperations          []AccountOperation              `gorm:"foreignKey:BlockNum"`
	}
)