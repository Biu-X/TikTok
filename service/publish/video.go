package publish

import (
	"biu-x.org/TikTok/dal/query"
	"biu-x.org/TikTok/model"
	"biu-x.org/TikTok/module/config"
	"biu-x.org/TikTok/module/log"
	"biu-x.org/TikTok/module/response"
	"biu-x.org/TikTok/module/s3"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func PublishVideo(c *gin.Context) {
	file, err := c.FormFile("data")
	if err != nil {
		log.Logger.Error(err)
		response.ErrRespWithMsg(c, err.Error())
		return
	}

	log.Logger.Infof("file name: %v", file.Filename)

	dst := "./" + file.Filename
	// 上传文件至指定的完整文件路径
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		log.Logger.Error(err)
		response.ErrRespWithMsg(c, err.Error())
	}

	userId, exists := c.Get("user_id")
	if !exists {
		response.ErrRespWithMsg(c, "user id is null")
	}
	fileName := fmt.Sprintf("%v/%v", userId, file.Filename)
	err = s3.PutFromFile(fileName, dst)
	if err != nil {
		log.Logger.Error(err)
		response.ErrRespWithMsg(c, "upload to s3 field")
	}

	v := query.Video
	aid, err := strconv.Atoi(fmt.Sprintf("%v", userId))
	if err != nil {
		response.ErrResp(c)
	}

	err = v.Create(&model.Video{
		AuthorID: int64(aid),
		PlayURL:  fmt.Sprintf("https://%v.%v/%v", config.S3Config.Bucket, config.S3Config.Endpoint, fileName),
		Title:    c.Query("title"),
	})
	if err != nil {
		response.ErrResp(c)
	}
	response.OKResp(c)
}
