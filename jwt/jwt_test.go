package jwt

import (
	"testing"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
)

func TestExample(t *testing.T) {
	type CustomClaims struct {
		Role string
		jwtgo.StandardClaims
	}

	tokenString, err := Sign(&CustomClaims{
		Role: "NORMAL",
		StandardClaims: jwtgo.StandardClaims{
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
	if err := Verify(tokenString, c); err != nil {
		t.Fatal(err)
	}
}
