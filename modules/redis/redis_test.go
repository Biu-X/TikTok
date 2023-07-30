package redis

import (
	"fmt"
	"testing"
	"time"
)

func TestNewRedisClient(t *testing.T) {
	c := NewRedisClient()
	fmt.Printf("%#v\n", c.Set("name", "hiifong"))
	fmt.Printf("%#v\n", c.Get("name").Val())
	time.Sleep(12 * time.Second)
	fmt.Printf("%#v\n", c.Get("name").Val())
}
