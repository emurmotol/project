package service

import (
	"context"
	"fmt"

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

func (l loggingMiddleware) GetByUsername(ctx context.Context, username string) (data *GetByUsernameData, err error) {
	defer func() {
		l.logger.Log("method", "GetByUsername", "username", username, "data", fmt.Sprint(data), "err", err)
	}()
	return l.next.GetByUsername(ctx, username)
}

func (l loggingMiddleware) CreateUser(ctx context.Context, username string, email string, password string, role string) (data *CreateUserData, err error) {
	defer func() {
		l.logger.Log("method", "CreateUser", "username", username, "email", email, "password", password, "role", role, "data", fmt.Sprintln(data.User), "err", err)
	}()
	return l.next.CreateUser(ctx, username, email, password, role)
}
