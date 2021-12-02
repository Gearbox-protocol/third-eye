package core

// import "github.com/ethereum/go-ethereum/core/types"

type (
	Block struct {
		BlockNumber       int64                    `gorm:"primaryKey;column:id"` // Block Number
		Timestamp         uint64                   `gorm:"column:timestamp"`
		AccountOperations []*AccountOperation      `gorm:"foreignKey:block_num"`
		TokenOracles      []*TokenOracle           `gorm:"foreignKey:block_num"`
		PriceFeeds        []*PriceFeed             `gorm:"foreignKey:block_num"`
		Protocols         []*Protocol              `gorm:"foreignKey:block_num"`
		CSS               []*CreditSessionSnapshot `gorm:"foreignKey:block_num"`
		PoolStats         []*PoolStat              `gorm:"foreignKey:block_num"`
		PoolLedgers       []*PoolLedger            `gorm:"foreignKey:block_num"`
		CMStats           []*CreditManagerStat     `gorm:"foreignKey:block_num"`
	}
)

func (Block) TableName() string {
	return "blocks"
}

func (b *Block) AddAccountOperation(accountOperation *AccountOperation) {
	b.AccountOperations = append(b.AccountOperations, accountOperation)
}
func (b *Block) AddTokenOracle(tokenOracle *TokenOracle) {
	b.TokenOracles = append(b.TokenOracles, tokenOracle)
}
func (b *Block) AddPriceFeed(pf *PriceFeed) {
	b.PriceFeeds = append(b.PriceFeeds, pf)
}
func (b *Block) AddAllowedProtocol(p *Protocol) {
	b.Protocols = append(b.Protocols, p)
}

func (b *Block) AddCreditSessionSnapshot(css *CreditSessionSnapshot) {
	b.CSS = append(b.CSS, css)
}

func (b *Block) AddPoolStat(ps *PoolStat) {
	b.PoolStats = append(b.PoolStats, ps)
}
func (b *Block) AddPoolLedger(pl *PoolLedger) {
	b.PoolLedgers = append(b.PoolLedgers, pl)
}

func (b *Block) AddCreditManagerStats(cms *CreditManagerStat) {
	b.CMStats = append(b.CMStats, cms)
}
