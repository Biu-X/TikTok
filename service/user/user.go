package user

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"biu-x.org/TikTok/dal/query"
	"biu-x.org/TikTok/model"
	"biu-x.org/TikTok/module/db"
	"biu-x.org/TikTok/module/middleware/jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SaveUser() {
	q := query.Use(db.DB)
	if !q.Available() {
		fmt.Println("query Use(db) fail: query.Available() == false")
	}
	do := q.WithContext(context.Background())

	u := &model.User{
		Name:     "i@hiif.ong",
		Password: "123456",
	}

	err := do.User.Create(u)
	if err != nil {
		fmt.Println(err)
		return
	}
}

/**
*
* 包含响应状态码和响应信息
* field：StatusCode 状态码 ，值为 0（正常） 或者 1（异常）
* field：Message 状态信息，描述响应
*
 */
type Response struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"status_message"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User *model.User
}

// Signup 用户注册 /douyin/user/signup/
func Signup(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	// ....
	u := query.User
	user, err := query.User.Where(u.Name.Eq("username")).First()
	// 数据库查询出现错误，服务端错误
	if err != nil {
		println("signup: database search failed, err:", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	// 用户已存在
	if user.Name == username {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{
				StatusCode: 1,
				Message:    "User Already Exists",
			},
		})
		return
	}

	// 生成密码的 hash 值
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		println("signup: get password's hash failed....")
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	newuser := model.User{
		Name:     username,
		Password: string(hash),
	}
	// err := u.WithContext(ctx).Create(&user) // pass pointer of data to Create
	err = query.User.Create(&newuser)
	if err != nil {
		println("singup: create newuser failed, err: ", err)
		c.JSON(http.StatusServiceUnavailable, gin.H{})
		return
	}

	user, err = query.User.Where(u.Name.Eq("username")).First()
	// 数据库查询出现错误，服务端错误
	if err != nil {
		println("signup: insert user success but search failed, err: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	// 打印用户信息
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{
			StatusCode: 0,
			Message:    "signup success...",
		},
		UserId: user.ID,
	})
}

// Login Post /douyin/user/login/ 用户登录
func Login(c *gin.Context) {
	u := query.User
	username := c.Query("username")
	password := c.Query("password")
	// 生成密码的 hash 值
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		println("login: get password's hash failed....")
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	user, err := query.User.Where(u.Name.Eq(username)).First()
	if err != nil {
		println("login: database search failed, err:", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	// verify pass
	if user.Password == string(hashedPassword) {
		// 注册之后的下次登录成功，才会为其生成 token
		token := jwt.GenerateToken(username)
		// 打印相应信息和用户信息以及生成的 token 值
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{
				StatusCode: 0,
				Message:    "Login Success",
			},
			UserId: user.ID,
			Token:  token,
		})
	} else {
		c.JSON(http.StatusBadRequest, UserLoginResponse{
			Response: Response{
				StatusCode: 1,
				Message:    "Username or Password Error",
			},
		})
	}
}

// token 验证通过后，可以根据用户 id 查询用户的信息
func UserInfo(c *gin.Context) {
	u := query.User

	userId := c.Query("user_id")
	id, _ := strconv.ParseInt(userId, 10, 64)
	if user, err := query.User.Where(u.ID.Eq(id)).First(); err != nil {
		c.JSON(http.StatusInternalServerError, UserResponse{
			Response: Response{
				StatusCode: 1,
				Message:    "UserInfo: query user info by user id failed...",
			},
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{
				StatusCode: 0,
				Message:    "query success",
			},
			User: user,
		})
	}
}
