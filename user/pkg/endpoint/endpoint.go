package endpoint

import (
	"context"

	"github.com/emurmotol/project/user/pkg/grpc/pb"
	service "github.com/emurmotol/project/user/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// GetByUsernameRequest collects the request parameters for the GetByUsername method.
type GetByUsernameRequest struct {
	Username string `json:"username"`
}

// GetByUsernameResponse collects the response parameters for the GetByUsername method.
type GetByUsernameResponse struct {
	User service.User `json:"user"`
	Err  error        `json:"error"`
}

func ConvertUser(user service.User) *pb.User {
	deletedAt := ""
	if user.DeletedAt != nil {
		deletedAt = user.DeletedAt.String()
	}
	return &pb.User{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
		DeletedAt: deletedAt,
	}
}

// MakeGetByUsernameEndpoint returns an endpoint that invokes GetByUsername on the service.
func MakeGetByUsernameEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByUsernameRequest)
		user, err := s.GetByUsername(ctx, req.Username)
		return GetByUsernameResponse{
			User: user,
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
func (e Endpoints) GetByUsername(ctx context.Context, username string) (user service.User, err error) {
	request := GetByUsernameRequest{Username: username}
	response, err := e.GetByUsernameEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetByUsernameResponse).User, response.(GetByUsernameResponse).Err
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
	User service.User `json:"user"`
	Err  error        `json:"error"`
}

// MakeCreateUserEndpoint returns an endpoint that invokes CreateUser on the service.
func MakeCreateUserEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		user, err := s.CreateUser(ctx, req.Username, req.Email, req.Password, req.Role)
		return CreateUserResponse{
			User: user,
			Err:  err,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateUserResponse) Failed() error {
	return r.Err
}

// CreateUser implements Service. Primarily useful in a client.
func (e Endpoints) CreateUser(ctx context.Context, username string, email string, password string, role string) (user service.User, err error) {
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
	return response.(CreateUserResponse).User, response.(CreateUserResponse).Err
}

// GetUserForAuthRequest collects the request parameters for the GetUserForAuth method.
type GetUserForAuthRequest struct {
	Username string `json:"username"`
}

// GetUserForAuthResponse collects the response parameters for the GetUserForAuth method.
type GetUserForAuthResponse struct {
	User service.User `json:"user"`
	Err  error        `json:"error"`
}

// MakeGetUserForAuthEndpoint returns an endpoint that invokes GetUserForAuth on the service.
func MakeGetUserForAuthEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserForAuthRequest)
		user, err := s.GetUserForAuth(ctx, req.Username)
		return GetUserForAuthResponse{
			Err:  err,
			User: user,
		}, nil
	}
}

// Failed implements Failer.
func (r GetUserForAuthResponse) Failed() error {
	return r.Err
}

// GetUserForAuth implements Service. Primarily useful in a client.
func (e Endpoints) GetUserForAuth(ctx context.Context, username string) (user service.User, err error) {
	request := GetUserForAuthRequest{Username: username}
	response, err := e.GetUserForAuthEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetUserForAuthResponse).User, response.(GetUserForAuthResponse).Err
}
