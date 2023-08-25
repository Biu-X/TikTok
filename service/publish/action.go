package publish

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Biu-X/TikTok/dal/model"
	"github.com/Biu-X/TikTok/dao"
	"github.com/Biu-X/TikTok/module/ffmpeg"
	"github.com/Biu-X/TikTok/module/log"
	"github.com/Biu-X/TikTok/module/oss"
	"github.com/Biu-X/TikTok/module/response"
	"github.com/gin-gonic/gin"
)

// Action /douyin/publish/action/ - 视频投稿
func Action(c *gin.Context) {
	tempfolder := "TikTokOSS"

	// 在完成上传视频后把临时文件都删除
	defer func() {
		path := c.GetString("user_id")
		err := os.RemoveAll(fmt.Sprintf("%v/%v", tempfolder, path))
		if err != nil {
			log.Logger.Error(err)
			return
		}
	}()

	// 获取视频
	file, err := c.FormFile("data")
	if err != nil {
		log.Logger.Error(err)
		response.ErrRespWithMsg(c, err.Error())
		return
	}

	log.Logger.Infof("file name: %v, size: %v", file.Filename, file.Size)

	// 上传文件大小限制为 30M
	if file.Size > 31457280 {
		log.Logger.Infof("file is too large, size limit is 30M, you file size is: %v", file.Size)
		response.ErrRespWithMsg(c, "file is too large, size limit 30M")
		return
	}

	// 接收完视频后，提前返回，防止超时, 实际存入数据库还是等待所有文件上传完成才写入
	response.OKResp(c)

	userID, exists := c.Get("user_id")
	if !exists {
		response.ErrRespWithMsg(c, "user id is null")
		return
	}
	aid, err := strconv.Atoi(fmt.Sprintf("%v", userID))
	if err != nil {
		log.Logger.Error(err)
		response.ErrRespWithMsg(c, err.Error())
		return
	}

	ts := time.Now().Local().UnixNano()
	log.Logger.Infof("timestamp: %v", ts)

	str := strings.Split(file.Filename, ".")
	log.Logger.Infof("str: %v", str)
	// fileName 即是保存临时文件的路径与文件名，也是上传到对象存储的路径也文件名
	fileName := fmt.Sprintf("%v/%v/%v.%v", tempfolder, userID, ts, str[len(str)-1])
	cover := fmt.Sprintf("%v/%v/%v-cover.jpeg", tempfolder, aid, ts)

	// 上传文件至指定的完整文件路径
	err = c.SaveUploadedFile(file, fileName)
	if err != nil {
		log.Logger.Error(err)
		return
	}

	err = ffmpeg.CoverSnap(fileName, cover)
	if err != nil {
		log.Logger.Error(err)
	}

	// 上传视频到对象存储
	err = oss.PutFromFile(fileName, fileName)
	if err != nil {
		log.Logger.Error(err)
		return
	}

	// 上传封面到对象存储
	err = oss.PutFromFile(cover, cover)
	if err != nil {
		log.Logger.Error(err)
		return
	}

	err = dao.CreateVideo(&model.Video{
		AuthorID: int64(aid),
		PlayURL:  fmt.Sprintf("%v", fileName),
		CoverURL: fmt.Sprintf("%v", cover),
		Title:    c.PostForm("title"),
	})
	if err != nil {
		log.Logger.Error(err)
		return
	}
}
