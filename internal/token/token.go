package token

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// create function to createAccessToken, createRefreshToken, verifyAccessToken

type TokenMaker struct {
	JWTSecret     string
	RefreshSecret string
}

func (t *TokenMaker) CreateAccessToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(t.JWTSecret))
}

func (t *TokenMaker) CreateRefreshToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(t.RefreshSecret))
}

func (t *TokenMaker) VerifyAccessToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.JWTSecret), nil
	})
}
