package core

type TokenOracle struct {
	BlockNumber int64  `gorm:"primaryKey;column:block_num"`
	Token       string `gorm:"primaryKey;column:token"`
	Oracle      string `gorm:"column:oracle"`
	Feed        string `gorm:"column:feed"`
}

func (TokenOracle) TableName() string {
	return "token_oracle"
}

type PriceFeed struct {
	ID                 int64   `gorm:"primaryKey;column:id;autoIncrement:true"`
	BlockNumber        int64   `gorm:"column:block_num"`
	Token              string  `gorm:"column:token"`
	Feed               string  `gorm:"column:feed"`
	RoundId            int64   `gorm:"column:round_id"`
	PriceETHBI         *BigInt `gorm:"column:price_eth_bi"`
	Uniswapv2Price     float64 `gorm:"column:uniswapv2_price"`
	Uniswapv3Twap      float64 `gorm:"column:uniswapv3_twap"`
	Uniswapv3Price     float64 `gorm:"column:uniswapv3_price"`
	UniPriceFetchBlock int64   `gorm:"column:uni_price_fetch_block"`
	// PriceETHBI *BigInt `gorm:"column:price_eth_bi"`
	PriceETH float64 `gorm:"column:price_eth"`
}

func (PriceFeed) TableName() string {
	return "price_feeds"
}

type SortedPriceFeed []*PriceFeed

func (ts SortedPriceFeed) Len() int {
	return len(ts)
}
func (ts SortedPriceFeed) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}

// sort in increasing order by blockNumber,index
func (ts SortedPriceFeed) Less(i, j int) bool {
	return ts[i].BlockNumber < ts[i].BlockNumber
}
