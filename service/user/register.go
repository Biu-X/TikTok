package user

import (
	"github.com/Biu-X/TikTok/dal/model"
	"github.com/Biu-X/TikTok/dal/query"
	"github.com/Biu-X/TikTok/module/random"
	"github.com/Biu-X/TikTok/module/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Register /douyin/user/register/ - 用户注册接口
func Register(c *gin.Context) {
	// 优先从 url 中获取参数
	username := c.Query("username")
	password := c.Query("password")

	if len(username) == 0 && len(password) == 0 {
		username = c.Request.PostFormValue("username")
		password = c.Request.PostFormValue("password")
	}

	if len(username) == 0 || len(password) == 0 {
		c.Abort()
		response.ErrRespWithMsg(c, "username or password is required")
		return
	}

	u := query.User
	// 使用 First 查询时，如果查询不到结果默认报错，因此使用计数方法
	count, _ := u.Where(u.Name.Eq(username)).Count()
	if count > 0 {
		c.Abort()
		response.ErrRespWithMsg(c, "user already exist")
		return
	}

	// 生成密码的 hash 值
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		c.Abort()
		response.ErrRespWithMsg(c, "generate hash from password failed")
		return
	}

	avatar, bgIMG, signature := random.GetAvatarAndBGIMGAndSignature()
	newuser := model.User{
		Name:            username,
		Password:        string(hash),
		Avatar:          avatar,
		BackgroundImage: bgIMG,
		Signature:       signature,
	}
	// pass pointer of data to Create
	err = u.Create(&newuser)
	if err != nil {
		response.ErrRespWithMsg(c, "singup: create new user failed")
		return
	}

	user, err := u.Where(u.Name.Eq(username)).First()
	// 数据库查询出现错误，服务端错误
	if err != nil {
		response.ErrRespWithMsg(c, "signup: insert user success but search failed")
		return
	}

	// 打印注册信息
	// 注册成功时并不生成 token，第一次登录成功时才会生成
	response.OKRespWithData(c, map[string]interface{}{
		"user_id": user.ID,
		"token":   "",
	})
}
