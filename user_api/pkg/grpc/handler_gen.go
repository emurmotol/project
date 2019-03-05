// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package grpc

import (
	endpoint "github.com/emurmotol/project/user_api/pkg/endpoint"
	pb "github.com/emurmotol/project/user_api/pkg/grpc/pb"
	grpc "github.com/go-kit/kit/transport/grpc"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {
	getByUsername  grpc.Handler
	createUser     grpc.Handler
	getUserForAuth grpc.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]grpc.ServerOption) pb.UserApiServer {
	return &grpcServer{
		createUser:     makeCreateUserHandler(endpoints, options["CreateUser"]),
		getByUsername:  makeGetByUsernameHandler(endpoints, options["GetByUsername"]),
		getUserForAuth: makeGetUserForAuthHandler(endpoints, options["GetUserForAuth"]),
	}
}
