package core

type TokenOracle struct {
	BlockNumber int64  `gorm:"primaryKey;column:block_num" json:"blockNum"`
	Token       string `gorm:"primaryKey;column:token" json:"token"`
	Oracle      string `gorm:"column:oracle" json:"oracle"`
	Feed        string `gorm:"column:feed" json:"feed"`
	Version     int16  `gorm:"column:version" json:"version"`
}

func (TokenOracle) TableName() string {
	return "token_oracle"
}

type PriceFeed struct {
	ID                 int64   `gorm:"primaryKey;column:id;autoIncrement:true" json:"-"`
	BlockNumber        int64   `gorm:"column:block_num" json:"blockNum"`
	Token              string  `gorm:"column:token" json:"token"`
	Feed               string  `gorm:"column:feed" json:"feed"`
	RoundId            int64   `gorm:"column:round_id" json:"roundId"`
	PriceBI            *BigInt `gorm:"column:price_bi" json:"priceBI"`
	IsPriceInUSD       bool    `gorm:"column:price_in_usd" json:"isPriceInUSD"`
	Uniswapv2Price     float64 `gorm:"column:uniswapv2_price" json:"uniswapv2Price"`
	Uniswapv3Twap      float64 `gorm:"column:uniswapv3_twap" json:"uniswapv3Twap"`
	Uniswapv3Price     float64 `gorm:"column:uniswapv3_price" json:"uniswapv3Price"`
	UniPriceFetchBlock int64   `gorm:"column:uni_price_fetch_block" json:"uniPriceFetchBlock"`
	// PriceETHBI *BigInt `gorm:"column:price_eth_bi"`
	Price float64 `gorm:"column:price" json:"price"`
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
	return ts[i].BlockNumber < ts[j].BlockNumber
}
