package server

import (
	"context"
	"gateway/pb"
	"gateway/util"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthServer struct {
	pb.UnimplementedAuthServer
	config util.Config
}

func newAuthServer(config util.Config) (*AuthServer, error) {
	server := &AuthServer{
		config: config,
	}

	return server, nil
}

func (server *AuthServer) RegisterUser(cxt context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	user := pb.User{
		Username:          req.GetUsername(),
		FullName:          req.GetFullName(),
		Email:             req.GetEmail(),
		PasswordChangedAt: timestamppb.New(time.Now()),
		CreatedAt:         timestamppb.New(time.Date(2019, 10, 2, 12, 22, 22, 121, time.UTC)),
	}

	return &pb.RegisterUserResponse{
		User: &user,
	}, nil
}
