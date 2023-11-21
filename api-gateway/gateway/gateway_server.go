package gateway

import (
	"context"
	"fmt"
	"gateway/pb"
	"gateway/util"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

func newGateway(ctx context.Context, conn *grpc.ClientConn, opts runtime.ServeMuxOption) (http.Handler, error) {
	mux := runtime.NewServeMux(opts)

	for _, f := range []func(context.Context, *runtime.ServeMux, *grpc.ClientConn) error {
		pb.RegisterAuthHandler,
	} {
		if err := f(ctx, mux, conn); err != nil {
			return nil, err
		}
	}
	return mux, nil
}

func RunGatewayServer(config util.Config) {
	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: false,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := dial(ctx, "tcp", config.GRPCServerAddress)

	if err != nil {
		log.Fatal().Err(err)
	}

	go func() {
		<-ctx.Done()
		if err := conn.Close(); err != nil {
			log.Fatal().Err(err).Msg("failed to close a client connection to the gRPC server: ")
		}
	}()

	grpcMux, err := newGateway(ctx, conn, jsonOption)

	if err != nil {
		log.Fatal().Err(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", config.HTTPServerAddress)

	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener:")
	}

	log.Info().Msgf("start HTTP Gateway server at %s", listener.Addr().String())

	err = http.Serve(listener, mux)

	if err != nil {
		log.Fatal().Err(err).Msg("cannot start HTTP gateway server:")
	}
}

func dial(ctx context.Context, network, addr string) (*grpc.ClientConn, error) {
	switch network {
	case "tcp":
		return dialTCP(ctx, addr)
	case "unix":
		return dialUnix(ctx, addr)
	default:
		return nil, fmt.Errorf("unsupported network type %q", network)
	}
}

// dialTCP creates a client connection via TCP.
// "addr" must be a valid TCP address with a port number.
func dialTCP(ctx context.Context, addr string) (*grpc.ClientConn, error) {
	return grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

// dialUnix creates a client connection via a unix domain socket.
// "addr" must be a valid path to the socket.
func dialUnix(ctx context.Context, addr string) (*grpc.ClientConn, error) {
	d := func(ctx context.Context, addr string) (net.Conn, error) {
		return (&net.Dialer{}).DialContext(ctx, "unix", addr)
	}
	return grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithContextDialer(d))
}
