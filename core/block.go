package core

// import "github.com/ethereum/go-ethereum/core/types"

type (
	Block struct {
		BlockNumber         int64                     `gorm:"primaryKey;column:id"` // Block Number
		Timestamp           uint64                   `gorm:"column:timestamp"`
	}
)

func (Block) TableName() string {
	return "blocks"
}