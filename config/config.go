package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	var err error

	db, err = InitDatabase()
	if err != nil {
		return fmt.Errorf("Error at initialize database: %v", err)
	}

	return nil
}

func GetDatabase() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger := NewLogger(p)
	return logger
}
