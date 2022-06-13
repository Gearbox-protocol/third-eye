package config

type Config struct {
	// Authentication
	AddressProviderAddress string `env:"REACT_APP_ADDRESS_PROVIDER" validate:"required"`

	// Database
	DatabaseUrl string `env:"DATABASE_URL" validate:"required"`
	Domain      string `env:"DOMAIN"`

	// Environment
	Env string `env:"ENV" default:"development" validate:"required"`

	// Ethereum
	EthProvider        string `validate:"required"`
	EthProviderMainnet string `env:"ETH_PROVIDER_MAINNET" validate:"required"`
	EthProviderKovan   string `env:"ETH_PROVIDER_KOVAN" validate:"required"`
	EthProviderFork    string `env:"ETH_PROVIDER_FORK"`

	ChainId   uint   `validate:"required"`
	NetworkId string `env:"REACT_APP_CHAIN_ID" validate:"required"`

	Port                 string `env:"PORT" default:"8080" validate:"required"`
	AMPQUrl              string `env:"CLOUDAMQP_URL" validate:"required"`
	AMPQEnable           string `env:"AMPQ_ENABLE" validate:"required"`
	DebtDCMatchingStr    string `env:"DEBT_DC_MATCHING" validate:"required"`
	DebtDCMatching       bool
	DisableDebtEngineStr string `env:"DISABLE_DEBT_ENGINE" validate:"required"`
	DisableDebtEngine    bool
	ThrottleDebtCalStr   string `env:"THROTTLE_DEBT_CAL" validate:"required"`
	ThrottleDebtCal      bool
	ThrottleByHrsStr     string `env:"THROTTLE_HRS"`
	ThrottleByHrs        int64
	MiningAddr           string `env:"MINING_ADDR"`
	Rollback             string `env:"ROLLBACK"`
	Uniswapv2Router      string `env:"UNISWAPV2_ROUTER"`
	Uniswapv3Router      string `env:"UNISWAPV3_ROUTER"`
	Interval             int64
	IntervalStr          string `env:"INTERVAL" default:"25"`
}
