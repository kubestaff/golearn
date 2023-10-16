package db

import (
	"github.com/kubestaff/golearn/setting"
	"gorm.io/gorm"
)

func Migrate(dbConn *gorm.DB) {
	dbConn.AutoMigrate(&setting.Settings{})
}
