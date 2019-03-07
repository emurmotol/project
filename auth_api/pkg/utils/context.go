package utils

import (
	"context"
	"fmt"

	"github.com/emurmotol/project/auth_api/pkg/grpc/pb"
	"github.com/go-kit/kit/auth/jwt"
)

type contextKey string

const (
	UserApiClientContextKey contextKey = "user_api"
)

func GetClaims(ctx context.Context) *JWTClaims {
	return ctx.Value(jwt.JWTClaimsContextKey).(*JWTClaims)
}

func MustGetClaims(ctx context.Context) (*JWTClaims, error) {
	db, ok := ctx.Value(jwt.JWTClaimsContextKey).(*JWTClaims)
	if !ok {
		return nil, fmt.Errorf("failed to get %s from context", jwt.JWTClaimsContextKey)
	}
	return db, nil
}

func GetUserApiClient(ctx context.Context) pb.UserApiClient {
	return ctx.Value(UserApiClientContextKey).(pb.UserApiClient)
}

func MustGetUserApiClient(ctx context.Context) (pb.UserApiClient, error) {
	client, ok := ctx.Value(UserApiClientContextKey).(pb.UserApiClient)
	if !ok {
		return nil, fmt.Errorf("failed to get %s from context", UserApiClientContextKey)
	}
	return client, nil
}
