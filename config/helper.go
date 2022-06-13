/*
 * Gearbox monitoring
 * Copyright (c) 2021. Mikael Lazarev
 *
 */

package config

import (
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func NewConfig() *Config {

	var config Config

	filenames := []string{".env", "../.env"}

	for _, filename := range filenames {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal("Cant get working dir")
		}

		if strings.Contains(cwd, "/server/") {
			serverDir := "/server/"
			lastIndex := strings.Index(cwd, serverDir) + len(serverDir)
			filename = cwd[:lastIndex] + strings.TrimPrefix(filename, "./")
		}

		err = godotenv.Load(filename)
		if err != nil {
			log.Infof("Cant read .env config file %s: %s\n", filename, err)
		} else {
			log.Info("Getting configuration from " + filename)
		}
	}

	rv := reflect.ValueOf(&config).Elem()
	num := rv.NumField()
	for i := 0; i < num; i++ {
		envValue := rv.Type().Field(i).Tag.Get("env")
		defaultValue := rv.Type().Field(i).Tag.Get("default")
		if envValue != "" {
			value := strings.Replace(GetEnv(envValue, defaultValue), "\\n", "\n", -1)
			rv.Field(i).SetString(value)
		}
	}

	chainId, err := strconv.Atoi(config.NetworkId)
	if err != nil {
		log.Fatal("Cant get interval")
	}

	interval, err := strconv.Atoi(config.IntervalStr)
	if err != nil {
		log.Fatal("Cant get chain id")
	}
	config.Interval = int64(interval)

	if config.DebtDCMatchingStr == "1" {
		config.DebtDCMatching = true
	} else {
		config.DebtDCMatching = false
	}

	if config.DisableDebtEngineStr == "1" {
		config.DisableDebtEngine = true
	} else {
		config.DisableDebtEngine = false
	}

	if config.ThrottleDebtCalStr == "1" {
		config.ThrottleDebtCal = true
	} else {
		config.ThrottleDebtCal = false
	}
	throttleHr, err := strconv.Atoi(config.ThrottleByHrsStr)
	if err != nil {
		config.ThrottleByHrs = 1
	} else {
		config.ThrottleByHrs = int64(throttleHr)
	}

	config.ChainId = uint(chainId)
	switch config.ChainId {
	case 1:
		config.EthProvider = config.EthProviderMainnet
	case 42:
		config.EthProvider = config.EthProviderKovan
	case 1337:
		config.EthProvider = config.EthProviderFork
	}

	validate(&config)
	return &config

}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// Validate config structures. If errors found, it break program
func validate(config interface{}) {

	// config validation
	validate := validator.New()
	if err := validate.Struct(config); err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Fatalf("Validation error in file %s", err)
		}

		for _, err := range err.(validator.ValidationErrors) {
			log.Warnf("Configuration problem: %s doesn't set\n", err.Namespace())
		}

		// from here you can create your own error messages in whatever language you wish
		log.Fatal("Cant continue")
	}

}
