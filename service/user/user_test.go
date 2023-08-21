package user

import (
	"biu-x.org/TikTok/dal/query"
	"biu-x.org/TikTok/dal/model"
	"biu-x.org/TikTok/module/config"
	"biu-x.org/TikTok/module/db"
	"biu-x.org/TikTok/module/log"
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
