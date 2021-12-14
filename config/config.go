/*
 * Gearbox monitoring
 * Copyright (c) 2021. Mikael Lazarev
 *
 */

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

	Port           string `env:"PORT" default:"8080" validate:"required"`
	AMPQUrl        string `env:"CLOUDAMQP_URL" validate:"required"`
	AMPQEnable     string `env:"AMPQ_ENABLE" validate:"required"`
	WETHAddr       string `env:"WETH_ADDRESS" validate:"required"`
	DebtDCMatching string `env:"DEBT_DC_MATCHING" validate:"required"`
	DebtCheck      bool

	MiningAddr string `env:"MINING_ADDR"`
}
