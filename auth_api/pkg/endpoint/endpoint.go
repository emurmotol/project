package endpoint

import (
	"context"

	service "github.com/emurmotol/project/auth_api/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// LoginRequest collects the request parameters for the Login method.
type LoginRequest struct {
	Payload *LoginInput `json:"payload"`
}

// LoginResponse collects the response parameters for the Login method.
type LoginResponse struct {
	Data *LoginOutput `json:"data"`
	Err  error        `json:"err"`
}

// MakeLoginEndpoint returns an endpoint that invokes Login on the service.
func MakeLoginEndpoint(s service.AuthApiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		data, err := s.Login(ctx, req.Payload)
		return LoginResponse{
			Data: data,
			Err:  err,
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
func (e Endpoints) Login(ctx context.Context, payload *LoginInput) (data *LoginOutput, err error) {
	request := LoginRequest{Payload: payload}
	response, err := e.LoginEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(LoginResponse).Data, response.(LoginResponse).Err
}
