package utils

import (
	"context"
	"errors"

	stdcasbin "github.com/casbin/casbin"
	"github.com/go-kit/kit/auth/casbin"
	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/jinzhu/gorm"
)

type contextKey string

const (
	DBContextKey contextKey = "db"
)

func GetDB(ctx context.Context) *gorm.DB {
	return ctx.Value(DBContextKey).(*gorm.DB)
}

func MustGetDB(ctx context.Context) (*gorm.DB, error) {
	db, ok := ctx.Value(DBContextKey).(*gorm.DB)
	if !ok {
		return nil, errors.New("failed to get DBContextKey value from context")
	}
	return db, nil
}

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

func GetRequestMethod(ctx context.Context) string {
	return ctx.Value(grpc.ContextKeyRequestMethod).(string)
}

func MustGetRequestMethod(ctx context.Context) (string, error) {
	rm, ok := ctx.Value(grpc.ContextKeyRequestMethod).(string)
	if !ok {
		return "", errors.New("failed to get ContextKeyRequestMethod value from context")
	}
	return rm, nil
}

func GetCasbinEnforcer(ctx context.Context) *stdcasbin.Enforcer {
	return ctx.Value(casbin.CasbinEnforcerContextKey).(*stdcasbin.Enforcer)
}

func MustGetCasbinEnforcer(ctx context.Context) (*stdcasbin.Enforcer, error) {
	e, ok := ctx.Value(casbin.CasbinEnforcerContextKey).(*stdcasbin.Enforcer)
	if !ok {
		return nil, errors.New("failed to get CasbinEnforcerContextKey value from context")
	}
	return e, nil
}