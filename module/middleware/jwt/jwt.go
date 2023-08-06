package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var (
	key []byte
	t   *jwt.Token
	s   string
)

type Claims struct {
	jwt.RegisteredClaims
	CustomClaims map[string]interface{} `json:"custom_claims"`
}

var claims = Claims{
	RegisteredClaims: jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(6 * time.Hour)),
		Subject:   "Tiktok",
	},
	//CustomClaims: map[string]interface{}{
	//	"uid":      1,
	//	"username": "hiifong",
	//},
}

func CreateToken() string {

	key = []byte("hiifong")

	t = jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)
	s, err := t.SignedString(key)
	if err != nil {
		fmt.Printf("%#v\n", err)
	}

	fmt.Printf("%#v\n", s)
	return s
}

func VerifyToken(token string) (bool, error) {
	if token == "" {
		return false, nil
	}
	verified, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return true, nil
	}
	return verified.Valid, nil
}

func ParsToken(token string) {
	withClaims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		fmt.Printf("%#v\n", err)
	}
	fmt.Printf("%#v\n", withClaims.Valid)
	time.Sleep(6 * time.Second)
	cla, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	fmt.Printf("%#v\n", cla.Valid)
}
