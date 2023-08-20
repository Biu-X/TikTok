package user

import (
	"errors"
	"strconv"

	"biu-x.org/TikTok/dal/query"
	"biu-x.org/TikTok/model"
	"biu-x.org/TikTok/module/middleware/jwt"
	"biu-x.org/TikTok/module/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

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

	newuser := model.User{
		Name:     username,
		Password: string(hash),
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
		"UserId": user.ID,
		"Token":  "",
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
		response.ErrRespWithData(c, "username and password is required...", map[string]interface{}{
			"UserId": 0,
			"Token":  "",
		})
		return
	}

	user, err := u.Where(u.Name.Eq(username)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.ErrRespWithMsg(c, "You have not signup")
		return
	}
	// verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err == nil {
		// 注册之后的下次登录成功，才会为其生成 token
		token := jwt.GenerateToken(username)
		// 打印相应信息和用户信息以及生成的 token 值
		response.OKRespWithData(c, map[string]interface{}{
			"UserId": user.ID,
			"Token":  token,
		})
	} else {
		response.ErrRespWithData(c, "Invalid Username or Password", map[string]interface{}{
			"UserId": 0,
			"Token":  "",
		})
	}
}

// token 验证通过后，可以根据用户 id 查询用户的信息
func UserInfo(c *gin.Context) {
	idStr := c.GetString("user_id")
	id, _ := strconv.Atoi(idStr)
	userinfo, err := response.GetUserResponseByUserId(int64(id))
	if err != nil {
		response.ErrRespWithMsg(c, "User not found")
		return
	}

	response.OKRespWithData(c, map[string]interface{}{
		"user": *userinfo,
	})
}
