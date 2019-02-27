package service

import "context"

// AuthApiService describes the service.
type AuthApiService interface {
	Login(ctx context.Context, payload *LoginInput) (data *LoginOutput, err error)
}

type LoginInput struct{}

type LoginOutput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type basicAuthApiService struct{}

func (b *basicAuthApiService) Login(ctx context.Context, payload *LoginInput) (data *LoginOutput, err error) {
	// TODO implement the business logic of Login
	return data, err
}

// NewBasicAuthApiService returns a naive, stateless implementation of AuthApiService.
func NewBasicAuthApiService() AuthApiService {
	return &basicAuthApiService{}
}

// New returns a AuthApiService with all of the expected middleware wired in.
func New(middleware []Middleware) AuthApiService {
	var svc AuthApiService = NewBasicAuthApiService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
