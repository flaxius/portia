package grpc

import (
	"context"
	"github.com/flaxius/portia/proto-gen/authentication/rest_credentials"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

// RunServer runs gRPC service to publish ToDo service
func RunServer(ctx context.Context, v1API rest_credentials.RestAuthCredentialsServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	rest_credentials.RegisterRestAuthCredentialsServer(server, v1API)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Println("starting gRPC server...")
	return server.Serve(listen)
}
