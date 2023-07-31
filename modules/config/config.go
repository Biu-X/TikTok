package config

import (
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var config *viper.Viper

func InitConfig() {
	config = viper.New()
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

	config.SetConfigName("tiktok")
	config.AddConfigPath("./conf/")
	config.AddConfigPath("./")
	config.AddConfigPath("$HOME/.tiktok/")
	config.AddConfigPath("/etc/tiktok/")
	config.SetConfigType("yml")

	if err := config.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			// 配置文件未找到错误
			fmt.Println("Config file not found!")
		}
	}

	config.WatchConfig()
	config.WatchConfig()
	config.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})

	fmt.Println(config.Get("mysql"))
}

func Get(key string) interface{} {
	return config.Get(key)
}
