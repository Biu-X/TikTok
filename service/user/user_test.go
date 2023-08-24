package user

import (
	"github.com/Biu-X/TikTok/dal/model"
	"github.com/Biu-X/TikTok/dal/query"
	"github.com/Biu-X/TikTok/module/config"
	"github.com/Biu-X/TikTok/module/db"
	"github.com/Biu-X/TikTok/module/log"
	"testing"
)

func TestSaveUser(t *testing.T) {
	config.Init()
	log.Init()
	db.Init()
	err := query.Comment.Create(&model.Comment{Content: "hello"})
	if err != nil {
		log.Logger.Info(err)
	}
	info, err := query.Comment.Where(query.Comment.ID.Eq(1)).Delete()
	if err != nil {
		log.Logger.Info(err)
	}
	log.Logger.Infof("info: %v", info)
}
