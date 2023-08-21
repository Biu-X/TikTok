package random

import (
	"biu-x.org/TikTok/module/config"
	"biu-x.org/TikTok/module/db"
	"biu-x.org/TikTok/module/log"
	"biu-x.org/TikTok/module/oss"
	"testing"
)

func TestGetRandomSignature(t *testing.T) {
	config.Init()
	log.Init()
	db.Init()
	oss.Init()
	_, _ = GetRandomSignature()
}
