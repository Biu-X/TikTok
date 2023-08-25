package jwt

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/Biu-X/TikTok/module/log"

	"github.com/Biu-X/TikTok/dal/model"
	"github.com/Biu-X/TikTok/dal/query"
	"github.com/golang-jwt/jwt/v5"
)

//- jwt 最终表现形式 {
//  base64Url(json.marshal(header)).
//  base64Url(json.marshal(claim)).
//  signature
//  }
//- Claim Set 用于在 token 中记录用户的一些信息供服务端使用，是 jwt 的第二部分；
//- 第一部分： Header{Alg, TokenType} 表示使用的签名算法和使用的 Token 类型（比如 jwt）
//- 第三部分： Header 和 Cliam 分别序列化加 base64Url 编码组合后再使用服务器私钥
//- 签名后得到的签名值（也是经过了 base64Url 编码）

// 根据使用情况调整 jwt 过期时间
const (
	TokenExpiredDuration = time.Hour * 24
)

var mySigningKey = []byte("tiktok")

// GenerateTokne 用户登录成功后，根据 username 查询到用户后生成一个 token
func GenerateToken(username string) string {
	user, err := query.User.Where(query.User.Name.Eq(username)).First()
	if err != nil {
		log.Logger.Info(err)
		return ""
	}
	return GenToken(user)
}

// GenToken 生成 jwt(json web token)
func GenToken(u *model.User) string {
	userId := strconv.FormatInt(u.ID, 10)
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpiredDuration)),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    "gopher-dance",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   "token",
		ID:        userId, // jwt 中保存合法用户的 ID
	}

	// 使用指定的签名算法创建用于签名的字符串对象（使用 json 序列化和 base64Url 编码生成 jwt 的 1、2 部分））
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 以上面生成 token 作为签名值，使用 secret 进行签名获取签名值
	// 将 token 和生成的签名值使用 '.' 拼接后就生成了 jwt；
	// 这里有个坑：参数可以是任意类型，但是你传 string 类型就会失败，这里一定要使用字节切片
	tokenStr, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Logger.Info(err)
		return ""
	}
	return tokenStr
}

// ParseToken 负责解析客户端 Header 中包含的 jwt，解析成功返回用户的 Claims（包含了用户的信息）
func ParseToken(tokenString string) (*jwt.RegisteredClaims, error) {
	// 使用匿名函数先去查询服务器签名时使用的私钥，然后调用签名的验证算法进行验证
	// 验证通过后，将 tokenString 进行反编码并反序列化到 jwt.Token 结构体相应字段
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})
	if err != nil {
		log.Logger.Info(err)
	}

	// 对空接口类型值进行类型断言
	// 如果类型断言成功并且 token 的有效位为 true（ParseWithClaims 方法调用成功后会将 Vaild 设置为 true）
	if cliams, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return cliams, nil
	}

	return nil, errors.New("invalid token")
}
