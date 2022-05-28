package initialize

import (
	"DuDao/global"
	"DuDao/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

func RegisterTables (db *gorm.DB){
	err := db.AutoMigrate(
		models.User{},
		models.Article{},
		models.Img{},
		)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
