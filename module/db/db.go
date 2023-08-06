package db

import (
	"biu-x.org/TikTok/dal/query"
	"biu-x.org/TikTok/module/config"
	"biu-x.org/TikTok/module/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	DB = ConnectDB(config.MySQLDSN())
	query.SetDefault(DB)
	log.Logger.Debugf("Set query default database")
}

func ConnectDB(dsn string) (db *gorm.DB) {
	var err error

	log.Logger.Debugf("MySQL DSN: %v", config.MySQLDSN())

	db, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		log.Logger.Fatalf("connect db fail: %w", err)
	}

	return db
}
