package api

import (
	"auth/pb"
	"context"
	"math/rand"
)

func (server *Server) RegisterUser(cxt context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	statuses := map[int]string{
		0: "created",
		1: "unverified",
		2: "verified",
	}

	return &pb.RegisterUserResponse{
		RegistrationStatus: statuses[rand.Intn(3)],
	}, nil
}
