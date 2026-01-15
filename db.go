package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func openDB(filename string) (*gorm.DB, error) {
	// SQLite is great for beginner demos because it has no server.
	// GORM uses the driver to build the connection.
	return gorm.Open(sqlite.Open(filename), &gorm.Config{
		// Logger configuration
		Logger: logger.Default.LogMode(logger.Info),

		// Name
	})
}
