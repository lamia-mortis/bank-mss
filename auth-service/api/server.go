package api

import (
	"auth/pb"
	"auth/util"
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedAuthServer
	config util.Config
}

func newServer(config util.Config) (*Server, error) {
	server := &Server{
		config: config,
	}

	return server, nil
}

func RunServer(config util.Config) {
	server, err := newServer(config)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.ServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}

	log.Info().Msgf("start Auth gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start Auth gRPC server")
	}
}
