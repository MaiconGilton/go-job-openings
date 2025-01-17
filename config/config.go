package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
}

func GetLogger(p string) *Logger {
	logger := NewLogger(p)
	return logger
}
