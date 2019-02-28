package service

import (
	"context"
	"errors"
)

// AuthApiService describes the service.
type AuthApiService interface {
	Login(ctx context.Context, payload *LoginInput) (data *LoginOutput, err error)
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginOutput struct {
	Message string `json:"message"`
}

type basicAuthApiService struct{}

func (b *basicAuthApiService) Login(ctx context.Context, payload *LoginInput) (data *LoginOutput, err error) {
	// TODO implement the business logic of Login
	return &LoginOutput{Message: "OK"}, errors.New("Error")
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
