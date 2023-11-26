package api

import (
	db "auth/db/sqlc"
	"auth/pb"
	"auth/util"
	"database/sql"
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	_ "github.com/lib/pq"
)

type Server struct {
	pb.UnimplementedAuthServer
	store  *db.Store
	config util.Config
}

func newServer(config util.Config, store *db.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}

	return server, nil
}

func RunServer(config util.Config) {
	// pq driver should be imported
	conn, err := sql.Open(config.DbDriver, config.AuthDbUrl)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to the DB")
	}

	store := db.NewStore(conn)

	server, err := newServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot run server")
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
