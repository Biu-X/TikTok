package publish

import (
	"biu-x.org/TikTok/dal/query"
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

func PublishVideo(c *gin.Context) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	defer func(wg *sync.WaitGroup) {
		wg.Wait()
		path := c.GetString("user_id")
		err := os.RemoveAll(path)
		if err != nil {
			log.Logger.Error(err)
			return
		}
	}(&wg)

	file, err := c.FormFile("data")
	if err != nil {
		log.Logger.Error(err)
		response.ErrRespWithMsg(c, err.Error())
		return
	}

	log.Logger.Infof("file name: %v", file.Filename)

	userId, exists := c.Get("user_id")
	if !exists {
		response.ErrRespWithMsg(c, "user id is null")
		return
	}
	aid, err := strconv.Atoi(fmt.Sprintf("%v", userId))
	fileName := fmt.Sprintf("%v/%v", userId, file.Filename)
	// 上传文件至指定的完整文件路径
	err = c.SaveUploadedFile(file, fileName)
	if err != nil {
		log.Logger.Error(err)
		response.ErrRespWithMsg(c, err.Error())
		return
	}

	image, err := ffmpeg.GetCoverFromVideo(fileName, 10)
	if err != nil {
		log.Logger.Error(err)
	}

	img, err := imaging.Decode(image)
	if err != nil {
		log.Logger.Error(err)
	}

	cover := fmt.Sprintf("%v/%v-cover.jpeg", aid, file.Filename)
	err = imaging.Save(img, cover)
	if err != nil {
		log.Logger.Error(err)
	}

	err = s3.PutFromFile(fileName, fileName)
	if err != nil {
		log.Logger.Error(err)
		response.ErrRespWithMsg(c, "upload to s3 field")
		return
	}

	err = s3.PutFromFile(cover, cover)
	if err != nil {
		log.Logger.Error(err)
	}

	v := query.Video
	if err != nil {
		response.ErrResp(c)
		return
	}

	err = v.Create(&model.Video{
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
