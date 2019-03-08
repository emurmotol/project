package utils

import (
	"context"
	"errors"

	"github.com/emurmotol/project/auth/pkg/grpc/pb"
	"github.com/go-kit/kit/auth/jwt"
)

type contextKey string

const (
	UserClientContextKey contextKey = "user"
)

func GetClaims(ctx context.Context) *JWTClaims {
	return ctx.Value(jwt.JWTClaimsContextKey).(*JWTClaims)
}

func MustGetClaims(ctx context.Context) (*JWTClaims, error) {
	db, ok := ctx.Value(jwt.JWTClaimsContextKey).(*JWTClaims)
	if !ok {
		return nil, errors.New("failed to get JWTClaimsContextKey value from context")
	}
	return db, nil
}

func GetUserClient(ctx context.Context) pb.UserClient {
	return ctx.Value(UserClientContextKey).(pb.UserClient)
}

func MustGetUserClient(ctx context.Context) (pb.UserClient, error) {
	client, ok := ctx.Value(UserClientContextKey).(pb.UserClient)
	if !ok {
		return nil, errors.New("failed to get UserClientContextKey value from context")
	}
	return client, nil
}
