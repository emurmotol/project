package grpc

import (
	"context"
	"errors"

	endpoint "github.com/emurmotol/project/auth_api/pkg/endpoint"
	pb "github.com/emurmotol/project/auth_api/pkg/grpc/pb"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

// makeLoginHandler creates the handler logic
func makeLoginHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.LoginEndpoint, decodeLoginRequest, encodeLoginResponse, options...)
}

// decodeLoginResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodeLoginRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'AuthApi' Decoder is not impelemented")
}

// encodeLoginResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeLoginResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'AuthApi' Encoder is not impelemented")
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
// TODO implement the decoder
func decodeRestrictedRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'AuthApi' Decoder is not impelemented")
}

// encodeRestrictedResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeRestrictedResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'AuthApi' Encoder is not impelemented")
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
// TODO implement the decoder
func decodeHealthCheckRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'AuthApi' Decoder is not impelemented")
}

// encodeHealthCheckResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeHealthCheckResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'AuthApi' Encoder is not impelemented")
}
func (g *grpcServer) HealthCheck(ctx context1.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckReply, error) {
	_, rep, err := g.healthCheck.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.HealthCheckReply), nil
}
