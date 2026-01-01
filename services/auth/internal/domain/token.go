package domain

import "github.com/golang-jwt/jwt/v5"

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

type AccessTokenClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}
