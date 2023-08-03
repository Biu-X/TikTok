package config_test

import (
	"biu-x.org/TikTok/modules/config"
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	config.Init()
	fmt.Println(config.GetString("mysql.host"))
}
