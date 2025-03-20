package token

import (
	"crypto/ed25519"
	"encoding/hex"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	Issuer  string
	PrivKey ed25519.PrivateKey
	Expires time.Duration
}

func NewTokenService(issuer string, strPrivKeyStr string, strExpires string) TokenService {
	// Parse the private key
	privKeyBytes, err := hex.DecodeString(strPrivKeyStr)
	if err != nil {
		panic(err)
	}

	// Convert the private key to ed25519.PrivateKey
	privKey := ed25519.PrivateKey(privKeyBytes)

	// Parse the expiry
	expires, err := time.ParseDuration(strExpires)
	if err != nil {
		panic(err)
	}

	return TokenService{
		Issuer:  issuer,
		PrivKey: privKey,
		Expires: expires,
	}
}

func (t TokenService) GenerateToken(uuid string) (string, error) {
	// Create a new token
	claims := jwt.RegisteredClaims{
		ID:        uuid,
		Issuer:    t.Issuer,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(t.Expires)),
	}

	token := jwt.NewWithClaims(&jwt.SigningMethodEd25519{}, claims)

	// Sign the token
	tokenString, err := token.SignedString(t.PrivKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (t TokenService) VerifyToken(strToken string) (jwt.RegisteredClaims, error) {
	claim := jwt.RegisteredClaims{}
	_, err := jwt.ParseWithClaims(strToken, &claim, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodEd25519); !ok {
			return nil, jwt.ErrInvalidKeyType
		}
		return t.PrivKey.Public(), nil
	})

	if err != nil {
		return jwt.RegisteredClaims{}, err
	}

	return claim, nil
}
