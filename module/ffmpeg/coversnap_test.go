package ffmpeg

import (
	"github.com/Biu-X/TikTok/module/config"
	"github.com/Biu-X/TikTok/module/db"
	"github.com/Biu-X/TikTok/module/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	config.Init()
	log.Init()
	db.Init()
}

func TestCoverSnap(t *testing.T) {
	a := assert.New(t)

	err := CoverSnap("video.mp4", "test.jpg")

	a.Nil(err)
}
