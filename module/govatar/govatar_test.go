package govatar

import (
	"image/png"
	"os"
	"testing"
)

func TestCreateAvatar(t *testing.T) {
	avatar, err := CreateAvatarWithDefault("hiifong@qq.com")
	if err != nil {
		return
	}
	file, err := os.Create("hiifong.png")
	err = png.Encode(file, avatar)
	if err != nil {
		panic(err)
	}
}
