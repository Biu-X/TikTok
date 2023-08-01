package db

import (
	"biu-x.org/TikTok/modules/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	DB = ConnectDB(config.MySQLDSN())
}

func ConnectDB(dsn string) (db *gorm.DB) {
	var err error

	fmt.Println(config.MySQLDSN())

	db, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}

	return db
}
