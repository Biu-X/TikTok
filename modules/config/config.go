package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Server struct {
	Port int
	Mode string
}

type Log struct {
	Level string
	Mode  []string
	Path  string
}

type MySQL struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type DataBase struct {
	*MySQL
}

type Redis struct {
	Host     string
	Port     int
	Password string
	DB       int
}

type Cache struct {
	*Redis
}

func InitConfig() {
	viper.SetDefault("server", map[string]interface{}{
		"port": 8080,
		"mode": "prod",
	})

	viper.SetDefault("log", map[string]interface{}{
		"level": "debug",
		"mode":  []string{"console", "file"},
		"path":  "./log",
	})

	viper.SetDefault("mysql", map[string]interface{}{
		"host":     "127.0.0.1",
		"port":     3306,
		"username": "root",
		"password": "123456",
		"database": "demo",
	})

	viper.SetDefault("redis", map[string]interface{}{
		"host":     "127.0.0.1",
		"port":     6379,
		"password": "123456",
		"database": 0,
	})

	//viper.SetConfigFile("./conf/tiktok.yml")
	viper.SetConfigName("tiktok")
	viper.AddConfigPath("./conf/")
	viper.AddConfigPath("/etc/tiktok/")
	viper.AddConfigPath("$HOME/.tiktok/")
	viper.AddConfigPath("./")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误
			fmt.Println("Config file not found!")
		} else {
			// 配置文件被找到，但产生了另外的错误
			fmt.Printf("some error: %v\n", err)
		}
	}

	viper.WatchConfig()
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})

	fmt.Println(viper.Get("database.mysql.host"))
}
