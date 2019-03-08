package service

import (
	"context"
	"fmt"

	"github.com/emurmotol/project/auth/pkg/utils"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(AuthService) AuthService

type loggingMiddleware struct {
	logger log.Logger
	next   AuthService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a AuthService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next AuthService) AuthService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Login(ctx context.Context, username string, password string) (accessToken string, tokenType string, expiresAt int64, err error) {
	defer func() {
		l.logger.Log("method", "Login", "username", username, "password", password, "accessToken", accessToken, "tokenType", tokenType, "expiresAt", expiresAt, "err", err)
	}()
	return l.next.Login(ctx, username, password)
}

func (l loggingMiddleware) Restricted(ctx context.Context) (claims *utils.JWTClaims, err error) {
	defer func() {
		l.logger.Log("method", "Restricted", "claims", fmt.Sprint(claims), "err", err)
	}()
	return l.next.Restricted(ctx)
}

func (l loggingMiddleware) HealthCheck(ctx context.Context) (status string, err error) {
	defer func() {
		l.logger.Log("method", "HealthCheck", "status", status, "err", err)
	}()
	return l.next.HealthCheck(ctx)
}
