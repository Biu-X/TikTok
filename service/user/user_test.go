package user

import (
	"biu-x.org/TikTok/module/config"
	"biu-x.org/TikTok/module/db"
	"testing"
)

func TestSaveUser(t *testing.T) {
	config.Init()
	db.Init()
	SaveUser()
}
