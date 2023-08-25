package dao

import (
	"testing"

	"github.com/Biu-X/TikTok/module/config"
	"github.com/Biu-X/TikTok/module/db"
	"github.com/Biu-X/TikTok/module/log"
	"github.com/stretchr/testify/assert"
)

// 在go test指令启动的测试中，各个文件之间是并发的，但每个文件中的TestXxx函数是串行的
func init() {
	config.Init()
	log.Init()
	db.Init()
}

var userID, videoID int64 = 114, 514

func setup() {
	err := SetFavoriteByUserIDAndVideoID(userID, videoID)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
}

func cleanup() {
	err := SetFavoriteCancelByUserIDAndVideoID(userID, videoID)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
}

func Test_GetFavoriteCountByVideoID(t *testing.T) {
	setup()
	defer cleanup()
	a := assert.New(t)

	expectFavoriteNumber := int64(1)
	favoriteNumber, err := GetFavoriteCountByVideoID(videoID)

	a.Nil(err)
	a.Equal(expectFavoriteNumber, favoriteNumber)
}

func Test_GetUserIsFavoriteVideo(t *testing.T) {
	setup()
	defer cleanup()
	a := assert.New(t)

	isLove := GetUserIsFavoriteVideo(userID, videoID)

	a.Equal(true, isLove)
}

func Test_GetFavoriteByUserID(t *testing.T) {
	setup()
	defer cleanup()
	a := assert.New(t)

	favorites, err := GetFavoriteByUserID(userID)

	a.Nil(err)
	a.Equal(1, len(favorites))
}

func Test_SetFavoriteCancelByUserIDAndVideoID(t *testing.T) {
	setup()
	defer cleanup()
	a := assert.New(t)

	err := SetFavoriteCancelByUserIDAndVideoID(userID, videoID)

	a.Nil(err)
}
