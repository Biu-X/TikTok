package user

import (
	"errors"

	"github.com/Biu-X/TikTok/dal/query"
	"github.com/Biu-X/TikTok/module/middleware/jwt"
	"github.com/Biu-X/TikTok/module/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Login /douyin/user/login/ - 用户登录
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
			"user_id": 0,
			"token":   "",
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
			"user_id": user.ID,
			"token":   token,
		})
	} else {
		response.ErrRespWithData(c, "Invalid Username or Password", map[string]interface{}{
			"user_id": 0,
			"token":   "",
		})
	}
}
