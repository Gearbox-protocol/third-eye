/*
 * Gearbox monitoring
 * Copyright (c) 2021. Mikael Lazarev
 *
 */

package config

import (
	"os"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func NewConfig() *Config {

	var config Config

	filenames := []string{".env", "../.env", "../../.env"}

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
			log.Infof("Cant read .env config file %s: %s", filename, err)
		} else {
			log.Info("Getting configuration from " + filename)
		}
	}

	utils.ReadFromEnv(&config)
	utils.ReadFromEnv(&config.DebtConfig)
	utils.ReadFromEnv(&config.ReduntantConfig)
	utils.ReadFromEnv(&config.CommonEnvs)

	validate(&config)
	return &config

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
