/*
 * Gearbox monitoring
 * Copyright (c) 2021. Mikael Lazarev
 *
 */

 package repository

 import (
	"github.com/Gearbox-protocol/gearscan/config"
	"github.com/Gearbox-protocol/gearscan/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)
 
 // Connects to MongoDB using config credentials
 func NewDBClient(config *config.Config) *gorm.DB {
	// Getting database settings
	gormDB, err := gorm.Open(postgres.Open(config.DatabaseUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err)
	}
 
	return gormDB
}
 