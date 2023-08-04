package jwt

import (
	"testing"
)

func TestSingToken(t *testing.T) {
	token := CreateToken()
	ParsToken(token)
}
