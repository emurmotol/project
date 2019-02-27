package service

import "context"

// AuthApiService describes the service.
type AuthApiService interface {
	Login(ctx context.Context, payload *LoginRequest) (data *LoginResponse, err error)
}

type LoginResponse struct{}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
