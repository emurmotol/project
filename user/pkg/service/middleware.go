package service

import (
	"context"
	"fmt"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(UserService) UserService

type loggingMiddleware struct {
	logger log.Logger
	next   UserService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a UserService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next UserService) UserService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) GetByUsername(ctx context.Context, username string) (user User, err error) {
	defer func() {
		l.logger.Log("method", "GetByUsername", "username", username, "user", fmt.Sprint(user), "err", err)
	}()
	return l.next.GetByUsername(ctx, username)
}

func (l loggingMiddleware) CreateUser(ctx context.Context, username string, email string, password string, role string) (user User, err error) {
	defer func() {
		l.logger.Log("method", "CreateUser", "username", username, "email", email, "password", password, "role", role, "user", fmt.Sprint(user), "err", err)
	}()
	return l.next.CreateUser(ctx, username, email, password, role)
}

func (l loggingMiddleware) GetUserForAuth(ctx context.Context, username string) (user User, err error) {
	defer func() {
		l.logger.Log("method", "GetUserForAuth", "username", username, "user", fmt.Sprint(user), "err", err)
	}()
	return l.next.GetUserForAuth(ctx, username)
}
