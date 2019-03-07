package grpc

import (
	"context"

	endpoint "github.com/emurmotol/project/user_api/pkg/endpoint"
	pb "github.com/emurmotol/project/user_api/pkg/grpc/pb"
	"github.com/go-kit/kit/auth/jwt"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
	stdgrpc "google.golang.org/grpc"
)

func makeGetByUsernameHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	options = append(options, grpc.ServerBefore(jwt.GRPCToContext()))
	options = append(options, stdgrpc.UnaryInterceptor(grpc.Interceptor))
	return grpc.NewServer(endpoints.GetByUsernameEndpoint, decodeGetByUsernameRequest, encodeGetByUsernameResponse, options...)
}

func decodeGetByUsernameRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GetByUsernameRequest)
	return endpoint.GetByUsernameRequest{Username: req.Username}, nil
}

func encodeGetByUsernameResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.GetByUsernameResponse)
	if res.Err != nil {
		return nil, res.Err
	}
	return &pb.GetByUsernameReply{User: endpoint.ConvertUser(res.User), Error: ""}, nil
}
func (g *grpcServer) GetByUsername(ctx context1.Context, req *pb.GetByUsernameRequest) (*pb.GetByUsernameReply, error) {
	_, rep, err := g.getByUsername.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetByUsernameReply), nil
}

func makeCreateUserHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateUserEndpoint, decodeCreateUserRequest, encodeCreateUserResponse, options...)
}

func decodeCreateUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CreateUserRequest)
	return endpoint.CreateUserRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}, nil
}

func encodeCreateUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.CreateUserResponse)
	if res.Err != nil {
		return nil, res.Err
	}
	return &pb.CreateUserReply{User: endpoint.ConvertUser(res.User), Error: ""}, nil
}
func (g *grpcServer) CreateUser(ctx context1.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	_, rep, err := g.createUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateUserReply), nil
}

func makeGetUserForAuthHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetUserForAuthEndpoint, decodeGetUserForAuthRequest, encodeGetUserForAuthResponse, options...)
}

func decodeGetUserForAuthRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GetUserForAuthRequest)
	return endpoint.GetUserForAuthRequest{Username: req.Username}, nil
}

func encodeGetUserForAuthResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.GetUserForAuthResponse)
	if res.Err != nil {
		return nil, res.Err
	}
	return &pb.GetUserForAuthReply{User: endpoint.ConvertUser(res.User), Error: ""}, nil
}
func (g *grpcServer) GetUserForAuth(ctx context1.Context, req *pb.GetUserForAuthRequest) (*pb.GetUserForAuthReply, error) {
	_, rep, err := g.getUserForAuth.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetUserForAuthReply), nil
}
