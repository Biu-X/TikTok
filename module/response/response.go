package response

import (
	"net/http"

	"biu-x.org/TikTok/dao"
	"github.com/gin-gonic/gin"
)

func OKResp(c *gin.Context) {
	resp := map[string]interface{}{
		"status_code": OK,
		"status_msg":  Msg(OK),
	}
	c.JSON(http.StatusOK, resp)
}

func OKRespWithData(c *gin.Context, data map[string]interface{}) {
	resp := map[string]interface{}{
		"status_code": OK,
		"status_msg":  Msg(OK),
	}
	for key, value := range data {
		resp[key] = value
	}
	c.JSON(http.StatusOK, resp)
}

func ErrResp(c *gin.Context) {
	resp := map[string]interface{}{
		"status_code": Error,
		"status_msg":  Msg(Error),
	}
	c.JSON(http.StatusOK, resp)
}

func ErrRespWithMsg(c *gin.Context, msg string) {
	resp := map[string]interface{}{
		"status_code": Error,
		"status_msg":  msg,
	}
	c.JSON(http.StatusOK, resp)
}

func ErrRespWithData(c *gin.Context, msg string, data map[string]interface{}) {
	resp := map[string]interface{}{
		"status_code": Error,
		"status_msg":  msg,
	}
	for key, value := range data {
		resp[key] = value
	}
	c.JSON(http.StatusOK, resp)
}

// 返回用户信息，id 为要获取的用户 id，userID 为当前登录用户 id
func GetUserResponseByID(id int64, userID int64) (*UserResponse, error) {
	isFollow, err := dao.GetIsFollowByBothID(id, userID)
	if err != nil { // 日志在上层调用时已经打印无需再打印
		return nil, err
	}

	user, err := dao.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	// 求用户关注了多少个用户，即求表中关注者 ID 为 userId 的列数
	followCount, err := dao.GetFollowingCountByUserID(id)
	if err != nil {
		return nil, err
	}

	// 求用户的关注者数量，即求表中用户 id 等于 userId 的列数
	followerCount, err := dao.GetFollowerCountByUserID(id)
	if err != nil {
		return nil, err
	}

	// 作品获赞数量（需要去 Video 表中查询该用户所有的 Video_ID，然后再去 Favorite 表中查询每一个 Video_ID 的获赞数）
	videoIDs, err := dao.GetVideoIDByAuthorID(id)
	if err != nil {
		return nil, err
	}

	acquireFavoriteTotal := int64(0)
	for _, videoID := range videoIDs {
		count, err := dao.GetFavoriteCountByVideoID(videoID)
		if err != nil {
			return nil, err
		}

		acquireFavoriteTotal += count
	}

	// 总的作品数量
	totalWork := int64(len(videoIDs))

	// 总的喜欢作品量
	totalFavorite, err := dao.GetFavoriteCountByUserID(id)
	if err != nil {
		return nil, err
	}

	userResponse := UserResponse{
		UserID:         user.ID,
		Username:       user.Name,
		FollowCount:    followCount,
		FollowerCount:  followerCount,
		IsFollow:       isFollow,
		Avatar:         user.Avatar,
		BackGroudImage: user.BackgroundImage,
		Signature:      user.Signature,
		TotalFavorite:  totalFavorite,
		WorkCount:      totalWork,
		FavoriteCount:  totalFavorite,
	}

	return &userResponse, nil
}

func GetUserResponseByUserId(id int64) (*UserResponse, error) {
	return GetUserResponseByID(id, id)
}
