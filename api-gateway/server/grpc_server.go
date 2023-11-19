package server

import (
	"gateway/pb"
	"gateway/util"
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunGrpcServer(config util.Config) {
	authServer, err := newAuthServer(config)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create auth server")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServer(grpcServer, authServer)
	reflection.Register(grpcServer)

	l, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}
	defer func() {
		if err := l.Close(); err != nil {
			log.Fatal().Err(err).Msgf("failed to close %s at %s", l.Addr().Network(), l.Addr().String())
		}
	}()

	log.Info().Msgf("start gRPC server at %s", l.Addr().String())

	err = grpcServer.Serve(l)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start gRPC server")
	}
}
