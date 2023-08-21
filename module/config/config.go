package config

import (
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"strings"
)

var config *viper.Viper

type MySQL struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Database int    `yaml:"database"`
}

type OSS struct {
	Endpoint  string `yaml:"endpoint"`
	AccessKey string `yaml:"accessKey"`
	SecretKey string `yaml:"secretKey"`
	Region    string `yaml:"region"`
	Bucket    string `yaml:"bucket"`

	// 如果是使用 minio，并且没有使用 https，需要设置为 true
	UseSsl *bool `yaml:"useSsl"`
	// 如果是使用 minio，需要设置为 true
	HostnameImmutable *bool `yaml:"hostnameImmutable"`
}

var (
	mysql      MySQL
	redis      Redis
	OSSConfig  OSS
	OSS_PREFIX string
)

func Init() {
	config = viper.New()

	config.SetConfigName("tiktok")
	config.AddConfigPath("./conf/")
	config.AddConfigPath("./")
	config.AddConfigPath("$HOME/.tiktok/")
	config.AddConfigPath("/etc/tiktok/")
	config.SetConfigType("yml")

	config.AutomaticEnv()
	config.SetEnvPrefix("TK")
	replacer := strings.NewReplacer(".", "_")
	config.SetEnvKeyReplacer(replacer)

	config.WatchConfig()
	config.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})

	if err := config.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			// 配置文件未找到错误
			fmt.Println("config file not found use default config")
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
				"database": "tiktok",
			})

			config.SetDefault("redis", map[string]interface{}{
				"host":     "127.0.0.1",
				"port":     6379,
				"password": "123456",
				"database": 0,
			})
		}
	}
	err := config.UnmarshalKey("mysql", &mysql)
	if err != nil {
		log.Fatalf("unable to decode into mysql struct, %v", err)
	}
	err = config.UnmarshalKey("redis", &redis)
	if err != nil {
		log.Fatalf("unable to decode into redis struct, %v", err)
	}
	err = config.UnmarshalKey("oss", &OSSConfig)
	if err != nil {
		log.Fatalf("unable to decode into oss struct, %v", err)
	}

	// use env var to set oss config when some field is nil
	if OSSConfig.Endpoint == "" {
		OSSConfig.Endpoint = config.GetString("oss.endpoint")
	}
	if OSSConfig.AccessKey == "" {
		OSSConfig.AccessKey = config.GetString("oss.accesskey")
	}
	if OSSConfig.SecretKey == "" {
		OSSConfig.SecretKey = config.GetString("oss.secretkey")
	}
	if OSSConfig.Region == "" {
		OSSConfig.Region = config.GetString("oss.region")
	}
	if OSSConfig.Bucket == "" {
		OSSConfig.Bucket = config.GetString("oss.bucket")
	}

	ossType := config.GetString("oss.type")
	fmt.Printf("oss type: %v", ossType)
	switch ossType {
	case "minio":
		OSS_PREFIX = fmt.Sprintf("http://%v/%v/", OSSConfig.Endpoint, OSSConfig.Bucket)
	case "cos":
		OSS_PREFIX = fmt.Sprintf("https://%v.%v/", OSSConfig.Bucket, OSSConfig.Endpoint)
	default:
		OSS_PREFIX = fmt.Sprintf("http://%v/%v/", OSSConfig.Endpoint, OSSConfig.Bucket)
	}
	fmt.Printf("OSS PREFIX: %v", OSS_PREFIX)
}

func Get(key string) interface{} {
	return config.Get(key)
}

func GetString(key string) string {
	return config.GetString(key)
}

func MySQLDSN() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v"+
		"?charset=utf8mb4&parseTime=True&loc=Local",
		mysql.Username, mysql.Password, mysql.Host,
		mysql.Port, mysql.Database)
}
