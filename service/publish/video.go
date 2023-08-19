package publish

import (
	"biu-x.org/TikTok/dao"
	"biu-x.org/TikTok/model"
	"biu-x.org/TikTok/module/config"
	"biu-x.org/TikTok/module/ffmpeg"
	"biu-x.org/TikTok/module/log"
	"biu-x.org/TikTok/module/response"
	"biu-x.org/TikTok/module/s3"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"sync"
)

// Action 投稿操作 /douyin/publish/action/
func Action(c *gin.Context) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	// 在完成上传视频后把临时文件都删除
	defer func(wg *sync.WaitGroup) {
		wg.Wait()
		path := c.GetString("user_id")
		err := os.RemoveAll(path)
		if err != nil {
			log.Logger.Error(err)
			return
		}
	}(&wg)

	// 获取视频
	file, err := c.FormFile("data")
	if err != nil {
		log.Logger.Error(err)
		response.ErrRespWithMsg(c, err.Error())
		return
	}

	log.Logger.Infof("file name: %v", file.Filename)

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

	// fileName 即是保存临时文件的路径与文件名，也是上传到对象存储的路径也文件名
	fileName := fmt.Sprintf("%v/%v", userID, file.Filename)
	// 上传文件至指定的完整文件路径
	err = c.SaveUploadedFile(file, fileName)
	if err != nil {
		log.Logger.Error(err)
		response.ErrRespWithMsg(c, err.Error())
		return
	}

	// 获取视频的第十帧截图
	image, err := ffmpeg.GetCoverFromVideo(fileName, 10)
	if err != nil {
		log.Logger.Error(err)
	}

	img, err := imaging.Decode(image)
	if err != nil {
		log.Logger.Error(err)
	}

	cover := fmt.Sprintf("%v/%v-cover.jpeg", aid, file.Filename)

	// 保存截图到临时文件
	err = imaging.Save(img, cover)
	if err != nil {
		log.Logger.Error(err)
	}

	// 上传视频到对象存储
	err = s3.PutFromFile(fileName, fileName)
	if err != nil {
		log.Logger.Error(err)
		response.ErrRespWithMsg(c, "upload to s3 field")
		return
	}

	// 上传封面到对象存储
	err = s3.PutFromFile(cover, cover)
	if err != nil {
		log.Logger.Error(err)
	}

	err = dao.CreateVideo(&model.Video{
		AuthorID: int64(aid),
		PlayURL:  fmt.Sprintf("https://%v.%v/%v", config.S3Config.Bucket, config.S3Config.Endpoint, fileName),
		CoverURL: fmt.Sprintf("https://%v.%v/%v", config.S3Config.Bucket, config.S3Config.Endpoint, cover),
		Title:    c.PostForm("title"),
	})
	if err != nil {
		response.ErrResp(c)
		return
	}
	response.OKResp(c)
	wg.Done()
}
