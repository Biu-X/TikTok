package random

import (
	"github.com/Biu-X/TikTok/module/config"
	"github.com/Biu-X/TikTok/module/db"
	"github.com/Biu-X/TikTok/module/log"
	"github.com/Biu-X/TikTok/module/oss"
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
