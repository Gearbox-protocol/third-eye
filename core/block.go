package core

import (
	"sort"
)

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
		allowedTokens     []*AllowedToken          `gorm:"foreignKey:block_num"`
		eventBalances     SortedEventbalances      `gorm:"-"`
		debts             []*Debt                  `gorm:"foreignKey:block_num"`
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

func (b *Block) AddAllowedToken(atoken *AllowedToken) {
	b.allowedTokens = append(b.allowedTokens, atoken)
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

func (b *Block) GetAllowedTokens() []*AllowedToken {
	return b.allowedTokens
}

func (b *Block) GetPriceFeeds() []*PriceFeed {
	return b.PriceFeeds
}

func (b *Block) AddEventBalance(eb *EventBalance) {
	b.eventBalances = append(b.eventBalances, eb)
}

func (b *Block) GetEventBalances() []*EventBalance {
	sort.Sort(b.eventBalances)
	return b.eventBalances
}

func (b *Block) GetCSS() []*CreditSessionSnapshot {
	return b.CSS
}

func (b *Block) GetPoolStats() []*PoolStat {
	return b.PoolStats
}

func (b *Block) AddDebt(debt *Debt) {
	b.debts = append(b.debts, debt)
}
