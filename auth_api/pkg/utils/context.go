package utils

import (
	"context"
	"errors"

	"github.com/emurmotol/project/auth_api/pkg/grpc/pb"
	"github.com/go-kit/kit/auth/jwt"
)

type contextKey string

const (
	UserApiContextKey contextKey = "user_api"
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

func GetUserApi(ctx context.Context) pb.UserApiClient {
	return ctx.Value(UserApiContextKey).(pb.UserApiClient)
}

func MustGetUserApi(ctx context.Context) (pb.UserApiClient, error) {
	client, ok := ctx.Value(UserApiContextKey).(pb.UserApiClient)
	if !ok {
		return nil, errors.New("failed to get user_api from context")
	}
	return client, nil
}
