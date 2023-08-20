package publish

import (
	"fmt"
	"os"
	"strconv"

	"biu-x.org/TikTok/dao"
	"biu-x.org/TikTok/model"
	"biu-x.org/TikTok/module/config"
	"biu-x.org/TikTok/module/ffmpeg"
	"biu-x.org/TikTok/module/log"
	"biu-x.org/TikTok/module/response"
	"biu-x.org/TikTok/module/s3"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"time"
)

// Action 投稿操作 /douyin/publish/action/
func Action(c *gin.Context) {
	// 在完成上传视频后把临时文件都删除
	defer func() {
		path := c.GetString("user_id")
		err := os.RemoveAll(path)
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
	// fileName 即是保存临时文件的路径与文件名，也是上传到对象存储的路径也文件名
	fileName := fmt.Sprintf("%v/%v-%v", userID, ts, file.Filename)
	cover := fmt.Sprintf("%v/%v-%v-cover.jpeg", aid, ts, file.Filename)

	// 上传文件至指定的完整文件路径
	err = c.SaveUploadedFile(file, fileName)
	if err != nil {
		log.Logger.Error(err)
		return
	}

	// 获取视频的第十帧截图
	image, err := ffmpeg.GetCoverFromVideo(fileName, 10)
	if err != nil {
		log.Logger.Error(err)
		return
	}

	img, err := imaging.Decode(image)
	if err != nil {
		log.Logger.Error(err)
		return
	}

	// 保存截图到临时文件
	err = imaging.Save(img, cover)
	if err != nil {
		log.Logger.Error(err)
		return
	}

	// 上传视频到对象存储
	err = s3.PutFromFile(fileName, fileName)
	if err != nil {
		log.Logger.Error(err)
		return
	}

	// 上传封面到对象存储
	err = s3.PutFromFile(cover, cover)
	if err != nil {
		log.Logger.Error(err)
		return
	}

	err = dao.CreateVideo(&model.Video{
		AuthorID: int64(aid),
		PlayURL:  fmt.Sprintf("https://%v.%v/%v", config.S3Config.Bucket, config.S3Config.Endpoint, fileName),
		CoverURL: fmt.Sprintf("https://%v.%v/%v", config.S3Config.Bucket, config.S3Config.Endpoint, cover),
		Title:    c.PostForm("title"),
	})
	if err != nil {
		log.Logger.Error(err)
		return
	}
}
