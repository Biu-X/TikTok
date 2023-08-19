package favorite

import (
	"biu-x.org/TikTok/dao"
	"biu-x.org/TikTok/module/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

// List /douyin/favorite/list/ - 喜欢列表
func List(c *gin.Context) {
	// 从 RequireAuth 处读取 user_id
	userId, _ := strconv.ParseInt(c.GetString("user_id"), 10, 64)

	// 根据 user_id 查询喜欢列表
	favorites, err := dao.GetFavoriteByUserID(userId)
	if err != nil {
		response.ErrRespWithMsg(c, err.Error())
		return
	}

	var videoList []response.VideoResponse

	for _, favorite := range favorites {
		video, err := dao.GetVideoByID(favorite.VideoID) // 根据 video_id 查询视频信息
		if err != nil {
			response.ErrRespWithMsg(c, err.Error())
			return
		}

		user, err := dao.GetUserByID(video.AuthorID) // 根据 author_id 查询作者信息
		if err != nil {
			response.ErrRespWithMsg(c, err.Error())
			return
		}

		followCount, err := dao.GetFollowingCountByFollowerID(userId) // 查询关注数
		if err != nil {
			response.ErrRespWithMsg(c, err.Error())
			return
		}

		followerCount, err := dao.GetFollowerCountByUserID(userId) // 查询粉丝数
		if err != nil {
			response.ErrRespWithMsg(c, err.Error())
			return
		}

		// 待实现...

		resUser := response.UserResponse{
			UserID:         user.ID,
			Username:       user.Name,
			FollowCount:    followCount,
			FollowerCount:  followerCount,
			IsFollow:       false,
			Avatar:         user.Avatar,
			BackGroudImage: user.BackgroundImage,
			Signature:      user.Signature,
			TotalFavorite:  1,
			WorkCount:      1,
			FavoriteCount:  1,
		}

		resVideo := response.VideoResponse{
			VideoID:       video.ID,
			Author:        resUser,
			PlayURL:       video.PlayURL,
			CoverURL:      video.CoverURL,
			FavoriteCount: 1,
			CommentCount:  1,
			IsFavorite:    true, // 已点赞
			Title:         video.Title,
		}

		videoList = append(videoList, resVideo)
	}

	response.OKRespWithData(c, map[string]interface{}{
		"video_list": videoList,
	})
}
