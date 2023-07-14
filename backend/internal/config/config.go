package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBConfig struct {
	DBPath string
}

func NewDBConnect(config *DBConfig) (*gorm.DB, error) {

	db, err := gorm.Open(sqlite.Open(config.DBPath), &gorm.Config{})

	if err != nil {

		return nil, err

	}

	return db, nil

}
