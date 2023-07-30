package jwt_test

import (
	"biu-x.org/TikTok/modules/middleware/jwt"
	"testing"
)

func TestSingToken(t *testing.T) {
	token := jwt.CreateToken()
	jwt.ParsToken(token)
}
