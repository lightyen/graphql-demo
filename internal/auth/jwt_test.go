package auth

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func TestExample(t *testing.T) {
	type CustomClaims struct {
		Role string
		jwt.StandardClaims
	}

	tokenString, err := SignJwt(&CustomClaims{
		Role: "NORMAL",
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Unix(),
			Issuer:    "lightyen",
			Subject:   "someone@example.com",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	c := &CustomClaims{}
	if err := VerifyJwt(tokenString, c); err != nil {
		t.Fatal(err)
	}
}
