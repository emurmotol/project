package utils

import (
	"context"
	"errors"

	"github.com/go-kit/kit/auth/jwt"
)

func GetClaims(ctx context.Context) *JWTClaims {
	return ctx.Value(jwt.JWTClaimsContextKey).(*JWTClaims)
}

func MustGetClaims(ctx context.Context) (*JWTClaims, error) {
	db, ok := ctx.Value(jwt.JWTClaimsContextKey).(*JWTClaims)
	if !ok {
		return nil, errors.New("failed to get claims from context")
	}
	return db, nil
}
