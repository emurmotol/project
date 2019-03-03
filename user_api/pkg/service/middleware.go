package service

import (
	"context"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(UserApiService) UserApiService

type loggingMiddleware struct {
	logger log.Logger
	next   UserApiService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a UserApiService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next UserApiService) UserApiService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) GetByUsername(ctx context.Context, username string) (data *GetByUsernameOutput, err error) {
	defer func() {
		l.logger.Log("method", "GetByUsername", "username", username, "data", data, "err", err)
	}()
	return l.next.GetByUsername(ctx, username)
}
