package endpoint

import (
	"context"

	service "github.com/emurmotol/project/auth_api/pkg/service"
	"github.com/emurmotol/project/auth_api/pkg/utils"
	endpoint "github.com/go-kit/kit/endpoint"
)

// LoginRequest collects the request parameters for the Login method.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse collects the response parameters for the Login method.
type LoginResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresAt   int64  `json:"expires_at"`
	Err         error  `json:"error"`
}

// MakeLoginEndpoint returns an endpoint that invokes Login on the service.
func MakeLoginEndpoint(s service.AuthApiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		accessToken, tokenType, expiresAt, err := s.Login(ctx, req.Username, req.Password)
		return LoginResponse{
			AccessToken: accessToken,
			TokenType:   tokenType,
			ExpiresAt:   expiresAt,
			Err:         err,
		}, nil
	}
}

// Failed implements Failer.
func (r LoginResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Login implements Service. Primarily useful in a client.
func (e Endpoints) Login(ctx context.Context, username string, password string) (accessToken string, tokenType string, expiresAt int64, err error) {
	request := LoginRequest{Username: username, Password: password}
	response, err := e.LoginEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(LoginResponse).AccessToken, response.(LoginResponse).TokenType, response.(LoginResponse).ExpiresAt, response.(LoginResponse).Err
}

// RestrictedRequest collects the request parameters for the Restricted method.
type RestrictedRequest struct{}

// RestrictedResponse collects the response parameters for the Restricted method.
type RestrictedResponse struct {
	Claims *utils.JWTClaims `json:"claims"`
	Err    error            `json:"error"`
}

// MakeRestrictedEndpoint returns an endpoint that invokes Restricted on the service.
func MakeRestrictedEndpoint(s service.AuthApiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		claims, err := s.Restricted(ctx)
		return RestrictedResponse{
			Claims: claims,
			Err:    err,
		}, nil
	}
}

// Failed implements Failer.
func (r RestrictedResponse) Failed() error {
	return r.Err
}

// Restricted implements Service. Primarily useful in a client.
func (e Endpoints) Restricted(ctx context.Context) (claims *utils.JWTClaims, err error) {
	request := RestrictedRequest{}
	response, err := e.RestrictedEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(RestrictedResponse).Claims, response.(RestrictedResponse).Err
}

// HealthCheckRequest collects the request parameters for the HealthCheck method.
type HealthCheckRequest struct{}

// HealthCheckResponse collects the response parameters for the HealthCheck method.
type HealthCheckResponse struct {
	Status string `json:"status"`
	Err    error  `json:"error"`
}

// MakeHealthCheckEndpoint returns an endpoint that invokes HealthCheck on the service.
func MakeHealthCheckEndpoint(s service.AuthApiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		status, err := s.HealthCheck(ctx)
		return HealthCheckResponse{
			Status: status,
			Err:    err,
		}, nil
	}
}

// Failed implements Failer.
func (r HealthCheckResponse) Failed() error {
	return r.Err
}

// HealthCheck implements Service. Primarily useful in a client.
func (e Endpoints) HealthCheck(ctx context.Context) (status string, err error) {
	request := HealthCheckRequest{}
	response, err := e.HealthCheckEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(HealthCheckResponse).Status, response.(HealthCheckResponse).Err
}
