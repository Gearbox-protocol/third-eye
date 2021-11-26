

package core


type TokenOracle struct {
	BlockNumber int64   `gorm:"primaryKey;column:block_num"`
	Token string   `gorm:"primaryKey;column:token"`
	Oracle string   `gorm:"column:oracle"`
}

func (TokenOracle) TableName() string {
	return "token_oracle"
}

type PriceFeed struct {
	BlockNumber int64   `gorm:"column:block_num"`
	Token string   `gorm:"column:token"`
	Feed string `gorm:"primaryKey;column:feed"`
	RoundId int64 `gorm:"primaryKey;column:round_id"`
	PriceETHBI string `gorm:"column:price_eth_bi"`
	// PriceETHBI *BigInt `gorm:"column:price_eth_bi"`
	PriceETH float64   `gorm:"column:price_eth"`
}

func (PriceFeed) TableName() string {
	return "price_feeds"
}

