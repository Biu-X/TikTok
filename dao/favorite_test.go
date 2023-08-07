package dao

import (
	"reflect"
	"testing"

	"biu-x.org/TikTok/model"
	"biu-x.org/TikTok/module/config"
	"biu-x.org/TikTok/module/db"
	"biu-x.org/TikTok/module/log"
)

func init() {
	config.Init()
	log.Init()
	db.Init()
}

func Test_FavoriteDAO(t *testing.T) {
	f := &model.Favorite{
		UserID:  0,
		VideoID: 0,
	}

	// ----------------------------
	// Test for CreateFavorite
	// ----------------------------
	err := CreateFavorite(f)
	if err != nil {
		t.Error("CreateFavorite fail", err)
		return
	}

	// ----------------------------
	// Test for GetFavoriteByBoth
	// ----------------------------
	favorite, err := GetFavoriteByBoth(f.UserID, f.VideoID)
	if err != nil {
		t.Error("GetFavoriteByBoth fail", err)
		return
	}
	if !reflect.DeepEqual(favorite, f) {
		t.Error("GetFavoriteByBoth result error")
		t.Error(favorite)
	}

	// ----------------------------
	// Test for GetFavoriteByID
	// ----------------------------
	favorite, err = GetFavoriteByID(f.ID)
	if err != nil {
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
	err = SetFavoriteCancelByID(f.ID, true)
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
