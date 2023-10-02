package config

import "gorm.io/gorm"

var (
	db *gorm.DB
)

func Init() error {
	return nil
}

func GetLogger(msg string) *Logger {

	logger := NewLogger(msg)
	return logger
}
