package utils

import (
	stdjwt "github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	stdjwt.StandardClaims
}
