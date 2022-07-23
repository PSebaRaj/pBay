package app

import (
	// lazy, change to github package
	"fmt"
	"internal/config"

	currencyProtos "github.com/psebaraj/pbay/products/proto/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var configCONF config.Config

func connectGRPCClient(cfg config.Config) currencyProtos.CurrencyClient {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", cfg.CurrencyGRPCConn.Address, cfg.CurrencyGRPCConn.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("ERROR: Could not connect to gRPC server @ %s:%s", cfg.CurrencyGRPCConn.Address, cfg.CurrencyGRPCConn.Port)
		panic(err)
	}

	currencyClient := currencyProtos.NewCurrencyClient(conn)

	return currencyClient

}

func FetchConfig(cfg config.Config) {
	configCONF = cfg
}

func ConnGRPCServers() currencyProtos.CurrencyClient {
	cc := connectGRPCClient(configCONF)
	return cc
}

