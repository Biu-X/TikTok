package config

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	Init()
	fmt.Println(GetString("mysql.host"))
	fmt.Println(MySQLDSN())
	fmt.Printf("%#v\n", OSSConfig)
	fmt.Printf("pool size: %v\n", GetString("redis.poolSize"))
}
