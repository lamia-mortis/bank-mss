package server

import (
	"context"
	"gateway/pb"
	"gateway/util"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	pb.UnimplementedAuthServer
	config util.Config
}

func newAuthServer(config util.Config) (pb.AuthServer, error) {
	server := &AuthServer{
		config: config,
	}

	return server, nil
}

func (server *AuthServer) RegisterUser(cxt context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	conn, err := grpc.Dial(server.config.AuthServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot connect the Auth server: %s", err)
	}
	defer conn.Close()

	c := pb.NewAuthClient(conn)

	res, err := c.RegisterUser(cxt, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error response from the Auth server: %s", err)
	}

	return res, nil
}
