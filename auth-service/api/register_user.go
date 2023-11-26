package api

import (
	"auth/pb"
	"context"
	"database/sql"
	"fmt"

	db "auth/db/sqlc"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) RegisterUser(cxt context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	user, err := server.store.GetUser(cxt, req.GetEmail())
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, status.Errorf(codes.Internal, "failed to register user")
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
		if err != nil {
			return nil, fmt.Errorf("failed to hash password: %w", err)
		}

		args := db.CreateUserParams{
			Name:     req.GetName(),
			Surname:  req.GetSurname(),
			Username: req.GetUsername(),
			Password: string(hashedPassword),
			Email:    req.GetEmail(),
		}

		_, err = server.store.CreateUser(cxt, args)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
		}

		return &pb.RegisterUserResponse{
			RegistrationStatus: "created",
		}, nil
	}

	rStatus := "unverified"

	if user.IsEmailVerified {
		rStatus = "verified"
	}

	return &pb.RegisterUserResponse{
		RegistrationStatus: rStatus,
	}, nil
}
