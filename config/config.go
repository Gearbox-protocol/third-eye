package config

import "github.com/Gearbox-protocol/sdk-go/log"

type Config struct {
	// Authentication
	AddressProviderAddress string `env:"REACT_APP_ADDRESS_PROVIDER" validate:"required"`

	// Database
	DatabaseUrl string `env:"DATABASE_URL" validate:"required"`
	Domain      string `env:"DOMAIN"`

	// Ethereum
	EthProvider string `env:"ETH_PROVIDER" validate:"required"`

	// chainid is set in ethclient module while creating rpc client
	ChainId int64

	// port for health service
	Port string `env:"PORT" default:"0" validate:"required"`

	// set rollback if we are deleting some data in db and rerunning third-eye for getting that data again, this prevents adding some sync adapter again.
	Rollback string `env:"ROLLBACK"`

	// for aggregated block feed , at interval x it should have the prices again
	Interval    int64
	IntervalStr string `env:"INTERVAL" default:"25"`
	// the batch size for filter logs, if third-eye is far behind the latest block in blockchain
	BatchSizeForHistoryStr string `env:"BATCH_SIZE_FOR_HISTORY" default:"5000"`
	BatchSizeForHistory    int64
	//
	UseTenderlyTrace string `env:"TENDERLY_TRACE" default:"1"`
	//
	DebtConfig
	ReduntantConfig
	log.CommonEnvs
}

type ReduntantConfig struct {
	// mining address was contract which minted 5000 credit accounts
	MiningAddr string `env:"MINING_ADDR"`
	// uniswap v2/v3 router were used for getting quote and v3 oracle prices for token/eth pairs
	Uniswapv2Router string `env:"UNISWAPV2_ROUTER"`
	Uniswapv3Router string `env:"UNISWAPV3_ROUTER"`
}
type DebtConfig struct {
	DebtDCMatchingStr    string `env:"DEBT_DC_MATCHING" validate:"required"`
	DebtDCMatching       bool
	DisableDebtEngineStr string `env:"DISABLE_DEBT_ENGINE" validate:"required"`
	DisableDebtEngine    bool
	ThrottleDebtCalStr   string `env:"THROTTLE_DEBT_CAL" validate:"required"`
	ThrottleDebtCal      bool
	ThrottleByHrsStr     string `env:"THROTTLE_HRS"`
	ThrottleByHrs        int64
}
