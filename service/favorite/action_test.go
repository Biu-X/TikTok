package favorite

import (
	"testing"

	"github.com/Biu-X/TikTok/dal/model"
	"github.com/Biu-X/TikTok/dal/query"
	"github.com/Biu-X/TikTok/dao"
	"github.com/Biu-X/TikTok/module/config"
	"github.com/Biu-X/TikTok/module/db"
	"github.com/Biu-X/TikTok/module/log"
)

func TestActionFavorite(t *testing.T) {
	config.Init()
	log.Init()
	db.Init()
	f := query.Favorite

	favorite := &model.Favorite{
		VideoID: 6,
		UserID:  2,
	}

	f.Create(favorite)

	// 查看点赞数
	count, _ := f.Where(f.VideoID.Eq(6), f.Cancel.Eq(0)).Count()
	expect := 1
	if count != int64(expect) {
		t.Fatalf("expect %v but got %v", expect, count)
	}

	// 取消点赞
	err := dao.SetFavoriteCancelByUserIDAndVideoID(2, 6)
	if err != nil {
		t.Fatal(err.Error())
	}

	// 查看点赞数
	count, _ = f.Where(f.VideoID.Eq(6), f.Cancel.Eq(0)).Count()
	expect = 0
	if count != int64(expect) {
		t.Fatalf("expect %v but got %v", expect, count)
	}
}
