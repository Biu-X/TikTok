package user

import (
	"errors"
	"net/http"
	"strconv"

	"biu-x.org/TikTok/dal/query"
	"biu-x.org/TikTok/dao"
	"biu-x.org/TikTok/model"
	"biu-x.org/TikTok/module/log"
	"biu-x.org/TikTok/module/middleware/jwt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//包含响应状态码和响应信息
//field：StatusCode 状态码 ，值为 0（正常） 或者 1（异常）
//field：Message 状态信息，描述响应

type Response struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"status_message"`
}

type UserSignupAndLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserInfoResponse struct {
	Response
	UserResponse `json:"user"`
}

type UserResponse struct {
	UserID         int64  `json:"id"`               // 用户ID
	Username       string `json:"name"`             // 用户名
	FollowCount    int64  `json:"follow_count"`     // 该用户关注了多少个其他用户
	FollowerCount  int64  `json:"follower_count"`   // 该用户粉丝总数
	IsFollow       bool   `json:"is_follow"`        // true: 已关注 false: 未关注
	Avatar         string `json:"avatar"`           // 头像
	BackGroudImage string `json:"background_image"` // 背景大图
	Signature      string `json:"signature"`        // 个人简介
	TotalFavorite  int64  `json:"total_favorite"`   // 该用户获赞总量
	WorkCount      int64  `json:"work_count"`       // 作品数量
	FavoriteCount  int64  `json:"favorite_count"`   // 喜欢的作品数量
}

// Signup 用户注册 /douyin/user/signup/
func Signup(c *gin.Context) {
	// 优先从 url 中获取参数
	username := c.Query("username")
	password := c.Query("password")

	if len(username) == 0 && len(password) == 0 {
		username = c.Request.PostFormValue("username")
		password = c.Request.PostFormValue("password")
	}

	if len(username) == 0 || len(password) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{
			StatusCode: -1,
			Message:    "username or password is required",
		})
		return
	}

	u := query.User
	// 使用 First 查询时，如果查询不到结果默认报错，因此使用计数方法
	count, _ := u.Where(u.Name.Eq(username)).Count()
	if count > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{
			StatusCode: -1,
			Message:    "user already exist",
		})
		return
	}

	// 生成密码的 hash 值
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Logger.Errorf("signup: get password's hash failed...., error: %v", err)
		println()
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	newuser := model.User{
		Name:     username,
		Password: string(hash),
	}
	// pass pointer of data to Create
	err = u.Create(&newuser)
	if err != nil {
		log.Logger.Errorf("singup: create new user failed, err: %v", err)
		c.JSON(http.StatusServiceUnavailable, gin.H{})
		return
	}

	user, err := u.Where(u.Name.Eq(username)).First()
	// 数据库查询出现错误，服务端错误
	if err != nil {
		log.Logger.Errorf("signup: insert user success but search failed, err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	// 打印注册信息
	c.JSON(http.StatusOK, UserSignupAndLoginResponse{
		Response: Response{
			StatusCode: 0,
			Message:    "newuser signup success...",
		},
		UserId: user.ID,
		Token:  "", // 注册成功时并不生成 token，第一次登录成功时才会生成
	})
}

// Login Post /douyin/user/login/ 用户登录
func Login(c *gin.Context) {
	u := query.User
	username := c.Query("username")
	password := c.Query("password")

	if len(username) == 0 && len(password) == 0 {
		username = c.Request.PostFormValue("username")
		password = c.Request.PostFormValue("password")
	}

	if len(username) == 0 || len(password) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, UserSignupAndLoginResponse{
			Response: Response{
				StatusCode: -1,
				Message:    "username or password is required",
			},
			UserId: 0,
			Token:  "",
		})
		return
	}

	user, err := u.Where(u.Name.Eq(username)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{
			StatusCode: -1,
			Message:    "You have not signup",
		})
		return
	}
	// verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err == nil {
		// 注册之后的下次登录成功，才会为其生成 token
		token := jwt.GenerateToken(username)
		// 打印相应信息和用户信息以及生成的 token 值
		c.JSON(http.StatusOK, UserSignupAndLoginResponse{
			Response: Response{
				StatusCode: 0,
				Message:    "Login Success",
			},
			UserId: user.ID,
			Token:  token,
		})
	} else {
		c.JSON(http.StatusBadRequest, UserSignupAndLoginResponse{
			Response: Response{
				StatusCode: 1,
				Message:    "Invalid Username or Password",
			},
			UserId: 0,
			Token:  "",
		})
	}
}

// token 验证通过后，可以根据用户 id 查询用户的信息
func UserInfo(c *gin.Context) {
	// 从 RequireAuth 处读取 user_id
	userId := c.GetString("user_id")
	id, _ := strconv.ParseInt(userId, 10, 64)
	user, err := dao.GetUserByID(id)
	// 错误处理和错误日志输出统一处理，本层只负责错误提前返回
	if checkError(c, err) {
		return
	}

	// 求用户关注了多少个用户，即求表中关注者 ID 为 userId 的列数
	followCount, err := dao.GetFollowingCountByUserID(id)
	if checkError(c, err) {
		return
	}

	// 求用户的关注者数量，即求表中用户 id 等于 userId 的列数
	followerCount, err := dao.GetFollowerCountByUserID(id)
	if checkError(c, err) {
		return
	}

	// 作品获赞数量（需要去 Video 表中查询该用户所有的 Video_ID，然后再去 Favorite 表中查询每一个 Video_ID 的获赞数）
	videoIds, err := dao.GetVideoIDByAuthorID(id)
	if checkError(c, err) {
		return
	}

	acquireFavoriteTotal := int64(0)
	for _, videoId := range videoIds {
		// 收集每一条视频的获赞量
		count, err := dao.GetFavoriteCountByVideoID(videoId)
		if checkError(c, err) {
			return
		}

		acquireFavoriteTotal += count
	}

	// 总的作品数量
	totalWork, err := dao.GetVideoCountByAuthorID(id)
	if checkError(c, err) {
		return
	}

	// 总的喜欢作品量
	totalFavorite, err := dao.GetFavoriteCountByUserID(id)
	if checkError(c, err) {
		return
	}

	c.JSON(http.StatusOK, UserInfoResponse{
		Response: Response{
			StatusCode: 0,
			Message:    "query success",
		},
		UserResponse: UserResponse{
			UserID:         user.ID,
			Username:       user.Name,
			FollowCount:    followCount,
			FollowerCount:  followerCount,
			IsFollow:       false,
			Avatar:         user.Avatar,
			BackGroudImage: user.BackgroundImage,
			Signature:      user.Signature,
			TotalFavorite:  totalFavorite,
			WorkCount:      totalWork,
			FavoriteCount:  totalFavorite,
		},
	})
}

func checkError(c *gin.Context, err error) bool {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{}) // 查询失败时允许返回 null
		log.Logger.Error(err.Error())
		return true
	}
	return false
}
