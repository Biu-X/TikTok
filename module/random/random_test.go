package random

import (
	"biu-x.org/TikTok/module/config"
	"biu-x.org/TikTok/module/db"
	"biu-x.org/TikTok/module/log"
	"biu-x.org/TikTok/module/oss"
	"testing"
)

func TestRandomAny(t *testing.T) {
	config.Init()
	log.Init()
	db.Init()
	oss.Init()
	r, err := Random(Avatar)
	if err != nil {
		log.Logger.Error(err)
		return
	}
	log.Logger.Infof("res: %v", r)

	r, err = Random(BackgroundIMG)
	if err != nil {
		log.Logger.Error(err)
		return
	}
	log.Logger.Infof("res: %v", r)

	r, err = Random(Signature)
	if err != nil {
		log.Logger.Error(err)
		return
	}
	log.Logger.Infof("res: %v", r)
}
