package service

import "context"

// UserApiService describes the service.
type UserApiService interface {
	GetByUsername(ctx context.Context, username string) (data *GetByUsernameOutput, err error)
}

type GetByUsernameOutput struct {
	Username string `json:"username"`
}

type basicUserApiService struct{}

func (b *basicUserApiService) GetByUsername(ctx context.Context, username string) (data *GetByUsernameOutput, err error) {
	return &GetByUsernameOutput{Username: username}, nil
}

// NewBasicUserApiService returns a naive, stateless implementation of UserApiService.
func NewBasicUserApiService() UserApiService {
	return &basicUserApiService{}
}

// New returns a UserApiService with all of the expected middleware wired in.
func New(middleware []Middleware) UserApiService {
	var svc UserApiService = NewBasicUserApiService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}