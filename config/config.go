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
	EthProviderFork    string `env:"ETH_PROVIDER_FORK" validate:"required"`

	ChainId   uint   `validate:"required"`
	NetworkId string `env:"REACT_APP_CHAIN_ID" validate:"required"`

	//TelegramBotToken string `env:"TG_BOT_TOKEN" validate:"required"`

	SentryDSN string `env:"SENTRY_DSN" validate:"required"`

	Port    string `env:"PORT" default:"8080" validate:"required"`
	AMPQUrl string `env:"CLOUDAMQP_URL" validate:"required"`
}
