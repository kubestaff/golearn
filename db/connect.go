package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

const dbName = "golearning.db"

func ConnectToDb() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
