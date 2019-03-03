package utils

import (
	stdjwt "github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	UserID int64 `json:"user_id"`
	stdjwt.StandardClaims
}
