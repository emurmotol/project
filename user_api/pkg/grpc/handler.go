package grpc

import (
	"context"

	endpoint "github.com/emurmotol/project/user_api/pkg/endpoint"
	pb "github.com/emurmotol/project/user_api/pkg/grpc/pb"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

// makeGetByUsernameHandler creates the handler logic
func makeGetByUsernameHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetByUsernameEndpoint, decodeGetByUsernameRequest, encodeGetByUsernameResponse, options...)
}

// decodeGetByUsernameResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
func decodeGetByUsernameRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GetByUsernameRequest)
	return endpoint.GetByUsernameRequest{Username: req.Username}, nil
}

// encodeGetByUsernameResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
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
