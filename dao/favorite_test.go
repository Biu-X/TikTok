package dao

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Biu-X/TikTok/dal/model"
	"github.com/Biu-X/TikTok/module/config"
	"github.com/Biu-X/TikTok/module/db"
	"github.com/Biu-X/TikTok/module/log"
	"gorm.io/gorm"
)

func init() {
	config.Init()
	log.Init()
	db.Init()
}

func Test_FavoriteDAO(t *testing.T) {
	// ----------------------------
	// Test for CreateFavorite
	// ----------------------------
	f := &model.Favorite{
		UserID:  5,
		VideoID: 102,
	}

	err := CreateFavorite(f)
	if err != nil {
		t.Error("CreateFavorite fail", err)
		return
	}

	expectFavoriteNumber := int64(1)
	favoriteNumber, err := GetFavoriteCountByVideoID(102)
	if err != nil {
		t.Fatal(err.Error())
	}
	if favoriteNumber != expectFavoriteNumber {
		t.Fatalf("expect favorite number is %v, but got %v", expectFavoriteNumber, favoriteNumber)
	}
	// ----------------------------
	// Test for GetFavoriteByBoth
	// ----------------------------

	isLove := GetUserIsFavoriteVideo(f.UserID, f.VideoID)
	expect := true
	if isLove != expect {
		t.Fatalf("expect %v, but got %v", expect, isLove)
	}

	// ----------------------------
	// Test for GetFavoriteByID
	// ----------------------------
	favorites, err := GetFavoriteByUserID(f.UserID)
	if err != nil {
		t.Error("GetFavoriteByUserID fail", err)
		return
	}
	for _, favorite := range favorites {
		t.Log(favorite)
	}

	// ----------------------------
	// Test for GetFavoriteByID
	// ----------------------------
	favorite, err := GetFavoriteByID(f.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("GetFavoriteByID fail", err)
		return
	}
	if !reflect.DeepEqual(favorite, f) {
		t.Error("GetFavoriteByID result error")
		t.Error(favorite)
	}

	// ----------------------------
	// Test for SetFavoriteCancelByID
	// ----------------------------
	err = SetFavoriteCancelByID(f.ID, 1)
	if err != nil {
		t.Error("SetFavoriteCancelByID fail", err)
		return
	}
	test_cancel_f, _ := GetFavoriteByID(f.ID)
	if test_cancel_f.Cancel == 0 {
		t.Error("SetFavoriteCancelByID result wrong")
		t.Error(test_cancel_f)
	}
}
