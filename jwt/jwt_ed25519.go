package jwt

import (
	"crypto/ed25519"
	"errors"

	jwtgo "github.com/dgrijalva/jwt-go"
)

var (
	signingMethodED25519 *SigningMethodEdDSA
	publicKey            ed25519.PublicKey
	privateKey           ed25519.PrivateKey
	ErrEdDSAVerification = errors.New("crypto/ed25519: verification error")
)

type SigningMethodEdDSA struct {
	Name string
}

func init() {
	signingMethodED25519 = &SigningMethodEdDSA{Name: "EdDSA"}
	jwtgo.RegisterSigningMethod(signingMethodED25519.Alg(), func() jwtgo.SigningMethod {
		return signingMethodED25519
	})
	var err error
	publicKey, privateKey, err = ed25519.GenerateKey(nil)
	if err != nil {
		panic(err)
	}
}

func (m *SigningMethodEdDSA) Alg() string {
	return m.Name
}

func (m *SigningMethodEdDSA) Verify(signingString string, signature string, key interface{}) error {
	var err error

	// Decode the signature
	var sig []byte
	if sig, err = jwtgo.DecodeSegment(signature); err != nil {
		return err
	}

	// Get the key
	var publicKey ed25519.PublicKey
	var ok bool
	if publicKey, ok = key.(ed25519.PublicKey); !ok {
		return jwtgo.ErrInvalidKeyType
	}

	if len(publicKey) != ed25519.PublicKeySize {
		return jwtgo.ErrInvalidKey
	}

	if ok := ed25519.Verify(publicKey, []byte(signingString), sig); !ok {
		return ErrEdDSAVerification
	}

	return nil
}

func (m *SigningMethodEdDSA) Sign(signingString string, key interface{}) (str string, err error) {
	// Get the key
	var privateKey ed25519.PrivateKey
	var ok bool
	if privateKey, ok = key.(ed25519.PrivateKey); !ok {
		return "", jwtgo.ErrInvalidKeyType
	}

	if len(privateKey) != ed25519.PrivateKeySize {
		return "", jwtgo.ErrInvalidKey
	}

	// Sign the string and return the encoded result
	sig := ed25519.Sign(privateKey, []byte(signingString))
	return jwtgo.EncodeSegment(sig), nil
}
