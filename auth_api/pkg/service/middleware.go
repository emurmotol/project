package service

import (
	"context"
	"fmt"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(AuthApiService) AuthApiService

type loggingMiddleware struct {
	logger log.Logger
	next   AuthApiService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a AuthApiService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next AuthApiService) AuthApiService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Login(ctx context.Context, payload *LoginInput) (data *LoginOutput, err error) {
	defer func() {
		l.logger.Log("method", "Login", "payload", fmt.Sprint(payload), "data", fmt.Sprint(data), "err", err)
	}()
	return l.next.Login(ctx, payload)
}

func (l loggingMiddleware) Restricted(ctx context.Context) (data *RestrictedOutput, err error) {
	defer func() {
		l.logger.Log("method", "Restricted", "data", fmt.Sprint(data), "err", err)
	}()
	return l.next.Restricted(ctx)
}

func (l loggingMiddleware) HealthCheck(ctx context.Context) (status string, err error) {
	defer func() {
		l.logger.Log("method", "HealthCheck", "status", status, "err", err)
	}()
	return l.next.HealthCheck(ctx)
}
