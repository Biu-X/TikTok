package relation

import (
	"biu-x.org/TikTok/dal/query"
	"biu-x.org/TikTok/dao"
	"biu-x.org/TikTok/model"
	"biu-x.org/TikTok/module/response"
	user_service "biu-x.org/TikTok/service/user"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Action /douyin/relation/action/ - 关系操作
func Action(c *gin.Context) {
	// 从 RequireAuth 处读取 user_id
	userId, _ := strconv.ParseInt(c.GetString("user_id"), 10, 64)
	// 从 request 中查询
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)    // 对方用户id
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64) // 1-关注，2-取消关注

	// 根据 action_type 执行不同的操作
	if actionType == 1 {
		err := dao.CreateFollow(&model.Follow{
			UserID:     toUserId,
			FollowerID: userId,
		})
		if err != nil {
			response.ErrRespWithMsg(c, err.Error())
			return
		}
	} else {
		err := dao.SetFollowCancelByBoth(toUserId, userId)
		if err != nil {
			response.ErrRespWithMsg(c, err.Error())
			return
		}
	}

	response.OKResp(c)
}

// FollowerInfoResponse 返回格式
type FollowerInfoResponse struct {
	StatusCode int                         `json:"status_code"`
	Message    string                      `json:"status_msg"`
	UserList   []user_service.UserResponse `json:"user_list"`
}

// FollowList 关注列表
func FollowList(c *gin.Context) {
	var follow model.Follow
	//// 从 RequireAuth 处读取 user_id
	follow.UserID, _ = strconv.ParseInt(c.GetString("user_id"), 10, 64)

	// 查询关注列表
	followListInfo, err := query.Follow.
		Select(query.Follow.UserID).
		Where(query.Follow.FollowerID.Eq(follow.UserID), query.Follow.Cancel.Eq(0)).Find()
	if err != nil {
		response.ErrRespWithMsg(c, err.Error())
		return
	}
	//var followerIDs []int64
	//for _, follow := range followListInfo {
	//	followerIDs = append(followerIDs, follow.UserID)
	//}
	//log.Logger.Debugf("这是关注列表的id：%v", followerIDs)

	// 遍历函数
	var followInfo []user_service.UserResponse

	for _, follow := range followListInfo {
		followerID := follow.UserID
		userInfo, err := user_service.GetUserInfoByID(followerID)
		if err != nil {
			// 处理错误，例如日志记录或其他操作
			continue // 继续下一个迭代
		}

		// 进行类型转换
		userResponse := user_service.UserResponse{
			UserID:         userInfo.UserID,
			Username:       userInfo.Username,
			FollowCount:    userInfo.FollowCount,
			FollowerCount:  userInfo.FollowerCount,
			IsFollow:       true,
			Avatar:         userInfo.Avatar,
			BackGroudImage: userInfo.BackGroudImage,
			Signature:      userInfo.Signature,
			TotalFavorite:  userInfo.TotalFavorite,
			WorkCount:      userInfo.WorkCount,
			FavoriteCount:  userInfo.FavoriteCount,
		}

		followInfo = append(followInfo, userResponse)
	}

	if followInfo == nil {
		// 根据数组followerIDs里的数据，进行推演，获取需要返回的json
		response.OKRespWithData(c, map[string]interface{}{
			"user_list": nil,
		})
	} else {
		// 根据数组followerIDs里的数据，进行推演，获取需要返回的json
		response.OKRespWithData(c, map[string]interface{}{
			"user_list": followInfo,
		})
	}

}
