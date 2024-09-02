/*
 * Gearbox monitoring
 * Copyright (c) 2021. Mikael Lazarev
 *
 */

package repository

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/third-eye/config"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connects to MongoDB using config credentials
func NewDBClient(config *config.Config) *gorm.DB {
	// Getting database settings
	return core.NewDBClient(config.DatabaseUrl, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
}
