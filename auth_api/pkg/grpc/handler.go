package grpc

import (
	"context"

	endpoint "github.com/emurmotol/project/auth_api/pkg/endpoint"
	pb "github.com/emurmotol/project/auth_api/pkg/grpc/pb"
	"github.com/emurmotol/project/auth_api/pkg/service"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

// makeLoginHandler creates the handler logic
func makeLoginHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.LoginEndpoint, decodeLoginRequest, encodeLoginResponse, options...)
}

// decodeLoginResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
func decodeLoginRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.LoginRequest)
	loginInput := &service.LoginInput{
		Username: req.Username,
		Password: req.Password,
	}
	return endpoint.LoginRequest{loginInput}, nil
}

// encodeLoginResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeLoginResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.LoginResponse)
	if res.Err != nil {
		return nil, res.Err
	}
	d := res.Data
	data := &pb.LoginOutput{
		AccessToken: d.AccessToken,
		TokenType:   d.TokenType,
		ExpiresAt:   d.ExpiresAt,
	}
	return &pb.LoginReply{Data: data, Error: ""}, nil
}
func (g *grpcServer) Login(ctx context1.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	_, rep, err := g.login.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.LoginReply), nil
}

// makeRestrictedHandler creates the handler logic
func makeRestrictedHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.RestrictedEndpoint, decodeRestrictedRequest, encodeRestrictedResponse, options...)
}

// decodeRestrictedResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
func decodeRestrictedRequest(_ context.Context, r interface{}) (interface{}, error) {
	_ = r.(*pb.RestrictedRequest)
	return endpoint.RestrictedRequest{}, nil
}

// encodeRestrictedResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeRestrictedResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.RestrictedResponse)
	if res.Err != nil {
		return nil, res.Err
	}
	c := res.Data.Claims
	data := &pb.RestrictedOutput{
		Claims: &pb.Claims{
			Audience:  c.Audience,
			ExpiresAt: c.ExpiresAt,
			Id:        c.Id,
			IssuedAt:  c.IssuedAt,
			Issuer:    c.Issuer,
			NotBefore: c.NotBefore,
			Subject:   c.Subject,
			UserID:    c.UserID,
		},
	}
	return &pb.RestrictedReply{Data: data, Error: ""}, nil
}
func (g *grpcServer) Restricted(ctx context1.Context, req *pb.RestrictedRequest) (*pb.RestrictedReply, error) {
	_, rep, err := g.restricted.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.RestrictedReply), nil
}

// makeHealthCheckHandler creates the handler logic
func makeHealthCheckHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.HealthCheckEndpoint, decodeHealthCheckRequest, encodeHealthCheckResponse, options...)
}

// decodeHealthCheckResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
func decodeHealthCheckRequest(_ context.Context, r interface{}) (interface{}, error) {
	_ = r.(*pb.HealthCheckRequest)
	return endpoint.HealthCheckRequest{}, nil
}

// encodeHealthCheckResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeHealthCheckResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.HealthCheckResponse)
	if res.Err != nil {
		return nil, res.Err
	}
	return &pb.HealthCheckReply{Status: res.Status, Error: ""}, nil
}
func (g *grpcServer) HealthCheck(ctx context1.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckReply, error) {
	_, rep, err := g.healthCheck.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.HealthCheckReply), nil
}
