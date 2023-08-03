package config

import (
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var config *viper.Viper

func Init() {
	config = viper.New()

	config.SetConfigName("tiktok")
	config.AddConfigPath("./conf/")
	config.AddConfigPath("./")
	config.AddConfigPath("../../conf")
	config.AddConfigPath("$HOME/.tiktok/")
	config.AddConfigPath("/etc/tiktok/")
	config.SetConfigType("yml")

	config.SetDefault("server", map[string]interface{}{
		"port": 8080,
		"mode": "prod",
	})

	config.SetDefault("log", map[string]interface{}{
		"level": "debug",
		"mode":  []string{"console", "file"},
		"path":  "./log",
	})

	config.SetDefault("mysql", map[string]interface{}{
		"host":     "127.0.0.1",
		"port":     3306,
		"username": "root",
		"password": "123456",
		"database": "demo",
	})

	config.SetDefault("redis", map[string]interface{}{
		"host":     "127.0.0.1",
		"port":     6379,
		"password": "123456",
		"database": 0,
	})

	if err := config.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			// 配置文件未找到错误
			fmt.Println("Config file not found!")
		}
	}

	config.WatchConfig()
	config.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})
}

func Get(key string) interface{} {
	return config.Get(key)
}

func GetString(key string) string {
	return config.GetString(key)
}

func Set(key string, value interface{}) {
	config.Set(key, value)
}

func MySQLDSN() string {
	return GetString("mysql.username") + ":" +
		GetString("mysql.password") + "@tcp(" +
		GetString("mysql.host") + ":" +
		GetString("mysql.port") + ")/" +
		GetString("mysql.database") +
		"?charset=utf8mb4&parseTime=True&loc=Local"
}
