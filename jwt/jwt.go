package jwt

import (
	jwtgo "github.com/dgrijalva/jwt-go"
)

func keyfn(t *jwtgo.Token) (interface{}, error) {
	return publicKey, nil
}

func Verify(tokenString string, claims jwtgo.Claims) error {
	token, err := jwtgo.ParseWithClaims(tokenString, claims, keyfn)
	if err != nil {
		return err
	}
	return token.Claims.Valid()
}

func Sign(claims jwtgo.Claims) (string, error) {
	return jwtgo.NewWithClaims(signingMethodED25519, claims).SignedString(privateKey)
}
