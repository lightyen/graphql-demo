package auth

import (
	"github.com/dgrijalva/jwt-go"
)

func keyfn(t *jwt.Token) (interface{}, error) {
	return publicKey, nil
}

func VerifyJwt(tokenString string, claims jwt.Claims) error {
	token, err := jwt.ParseWithClaims(tokenString, claims, keyfn)
	if err != nil {
		return err
	}
	return token.Claims.Valid()
}

func SignJwt(claims jwt.Claims) (string, error) {
	return jwt.NewWithClaims(signingMethodED25519, claims).SignedString(privateKey)
}
