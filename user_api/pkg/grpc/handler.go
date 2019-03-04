package grpc

import (
	"context"

	endpoint "github.com/emurmotol/project/user_api/pkg/endpoint"
	pb "github.com/emurmotol/project/user_api/pkg/grpc/pb"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

func makeGetByUsernameHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
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
	data := &pb.GetByUsernameData{Username: res.Data.Username}
	return &pb.GetByUsernameReply{Data: data, Error: ""}, nil
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
	data := &pb.CreateUserData{
		User: &pb.User{
			ID:       res.Data.User.ID,
			Username: res.Data.User.Username,
			Email:    res.Data.User.Email,
			Password: res.Data.User.Password,
			Role:     res.Data.User.Role,
		},
	}
	return &pb.CreateUserReply{Data: data, Error: ""}, nil
}
func (g *grpcServer) CreateUser(ctx context1.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	_, rep, err := g.createUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateUserReply), nil
}
