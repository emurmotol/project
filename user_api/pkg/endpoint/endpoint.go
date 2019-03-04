package endpoint

import (
	"context"

	service "github.com/emurmotol/project/user_api/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// GetByUsernameRequest collects the request parameters for the GetByUsername method.
type GetByUsernameRequest struct {
	Username string `json:"username"`
}

// GetByUsernameResponse collects the response parameters for the GetByUsername method.
type GetByUsernameResponse struct {
	Data *service.GetByUsernameData `json:"data"`
	Err  error                      `json:"error"`
}

// MakeGetByUsernameEndpoint returns an endpoint that invokes GetByUsername on the service.
func MakeGetByUsernameEndpoint(s service.UserApiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByUsernameRequest)
		data, err := s.GetByUsername(ctx, req.Username)
		return GetByUsernameResponse{
			Data: data,
			Err:  err,
		}, nil
	}
}

// Failed implements Failer.
func (r GetByUsernameResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// GetByUsername implements Service. Primarily useful in a client.
func (e Endpoints) GetByUsername(ctx context.Context, username string) (data *service.GetByUsernameData, err error) {
	request := GetByUsernameRequest{Username: username}
	response, err := e.GetByUsernameEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetByUsernameResponse).Data, response.(GetByUsernameResponse).Err
}

// CreateUserRequest collects the request parameters for the CreateUser method.
type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// CreateUserResponse collects the response parameters for the CreateUser method.
type CreateUserResponse struct {
	Data *service.CreateUserData `json:"data"`
	Err  error                   `json:"error"`
}

// MakeCreateUserEndpoint returns an endpoint that invokes CreateUser on the service.
func MakeCreateUserEndpoint(s service.UserApiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		data, err := s.CreateUser(ctx, req.Username, req.Email, req.Password, req.Role)
		return CreateUserResponse{
			Data: data,
			Err:  err,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateUserResponse) Failed() error {
	return r.Err
}

// CreateUser implements Service. Primarily useful in a client.
func (e Endpoints) CreateUser(ctx context.Context, username string, email string, password string, role string) (data *service.CreateUserData, err error) {
	request := CreateUserRequest{
		Email:    email,
		Password: password,
		Role:     role,
		Username: username,
	}
	response, err := e.CreateUserEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateUserResponse).Data, response.(CreateUserResponse).Err
}
