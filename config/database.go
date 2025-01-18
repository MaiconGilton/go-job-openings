package config

import (
	"job-openings/schemas"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	logger := GetLogger("Database")
	dbPath := "./db/main.db"

	// Check if database exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		// Create database file
		logger.Info("Database not found, creating...")

		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.ErrorF("Error at creating database directory: %v", err)
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			logger.ErrorF("Error at creating database file: %v", err)
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.ErrorF("Error at init database: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.ErrorF("Error at init auto migrate: %v", err)
		return nil, err
	}

	return db, nil
}
