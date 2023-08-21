package random

import (
	"biu-x.org/TikTok/module/govatar"
	"biu-x.org/TikTok/module/log"
	"fmt"
	"image/png"
	"os"
	"time"
)

func GenAvatar(username string) (string, error) {
	avatar, err := govatar.CreateAvatarWithDefault(username)
	if err != nil {
		return
	}
	fileName := fmt.Sprintf("avatar/%v.png", time.Now().UnixMilli())

	//defer func() {
	//	err := os.RemoveAll(fileName)
	//	if err != nil {
	//		log.Logger.Error(err)
	//		return
	//	}
	//}()

	file, err := os.Create(fileName)
	if err != nil {
		log.Logger.Error(err)

	}
	err = png.Encode(file, avatar)
	if err != nil {
		log.Logger.Error(err)
	}
}
