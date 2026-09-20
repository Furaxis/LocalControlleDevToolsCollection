package config

import (
	"fmt"
)

var (
	logger *Logger

// db     *gorm.DB
)

func Init() error {
	var err error

	if err != nil {
		return fmt.Errorf("error initializing sqlite: %v", err)
	}

	return nil
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
