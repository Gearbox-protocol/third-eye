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

	// Environment
	Env string `env:"ENV" default:"development" validate:"required"`

	// Ethereum
	EthProvider        string `validate:"required"`
	EthProviderMainnet string `env:"ETH_PROVIDER_MAINNET"`
	EthProviderKovan   string `env:"ETH_PROVIDER_KOVAN"`
	EthProviderFork    string `env:"ETH_PROVIDER_FORK"`

	ChainId   uint   `validate:"required"`
	NetworkId string `env:"REACT_APP_CHAIN_ID" validate:"required"`

	Port    string `env:"PORT" default:"8080" validate:"required"`
	AMPQUrl string `env:"CLOUDAMQP_URL" validate:"required"`
}
