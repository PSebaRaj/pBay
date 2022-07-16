package main

import (
	"fmt"
	"log"
	"net"

	"github.com/psebaraj/pbay/currency/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	protos "github.com/psebaraj/pbay/currency/protos/currency"
)

func main() {

	var opts []grpc.ServerOption

	gs := grpc.NewServer(opts...)

	protos.RegisterCurrencyServer(gs, &server.Currency{})

	// TODO: disable, just for testing with grpcurl
	reflection.Register(gs)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 9092))
	if err != nil {
		log.Fatal("Unable to listen", "error", err)
	}

	gs.Serve(l)

}
