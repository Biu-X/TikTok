package db

import (
	"github.com/Biu-X/TikTok/dal/model"
	"github.com/Biu-X/TikTok/dal/query"
	"github.com/Biu-X/TikTok/module/config"
	"github.com/Biu-X/TikTok/module/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	DB = ConnectDB(config.MySQLDSN())
	err := DB.AutoMigrate(
		model.Comment{},
		model.Favorite{},
		model.Follow{},
		model.Message{},
		model.User{},
		model.Video{},
	)
	if err != nil {
		log.Logger.Error(err)
		return
	}
	query.SetDefault(DB)
	log.Logger.Debugf("Set query default database")
}

func ConnectDB(dsn string) (db *gorm.DB) {
	var err error

	log.Logger.Debugf("MySQL DSN: %v", config.MySQLDSN())

	db, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		log.Logger.Fatalf("connect db fail: %v", err)
	}

	return db
}
