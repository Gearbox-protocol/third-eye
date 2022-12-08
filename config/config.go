package config

type Config struct {
	AppName string `env:"APP_NAME" default:"Third-eye"`
	// Authentication
	AddressProviderAddress string `env:"REACT_APP_ADDRESS_PROVIDER" validate:"required"`

	// Database
	DatabaseUrl string `env:"DATABASE_URL" validate:"required"`
	Domain      string `env:"DOMAIN"`

	// Environment
	Env string `env:"ENV" default:"development" validate:"required"`

	// Ethereum
	EthProvider string `env:"ETH_PROVIDER" validate:"required"`

	ChainId int64

	Port                   string `env:"PORT" default:"0" validate:"required"`
	AMQPUrl                string `env:"CLOUDAMQP_URL" validate:"required"`
	AMQPEnable             string `env:"AMPQ_ENABLE" validate:"required"`
	DebtDCMatchingStr      string `env:"DEBT_DC_MATCHING" validate:"required"`
	DebtDCMatching         bool
	DisableDebtEngineStr   string `env:"DISABLE_DEBT_ENGINE" validate:"required"`
	DisableDebtEngine      bool
	ThrottleDebtCalStr     string `env:"THROTTLE_DEBT_CAL" validate:"required"`
	ThrottleDebtCal        bool
	ThrottleByHrsStr       string `env:"THROTTLE_HRS"`
	ThrottleByHrs          int64
	MiningAddr             string `env:"MINING_ADDR"`
	Rollback               string `env:"ROLLBACK"`
	Uniswapv2Router        string `env:"UNISWAPV2_ROUTER"`
	Uniswapv3Router        string `env:"UNISWAPV3_ROUTER"`
	Interval               int64
	IntervalStr            string `env:"INTERVAL" default:"25"`
	BatchSizeForHistoryStr string `env:"BATCH_SIZE_FOR_HISTORY" default:"5000"`
	BatchSizeForHistory    int64
}
