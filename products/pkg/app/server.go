// Package classification Products (PBay)
//
// Documentation of the Products service endpoints (backend of PBay)
//
// Schemes: http
// BasePath: /products
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta

package app

import (
	"fmt"
	"internal/config"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	//"github.com/psebaraj/pbay/products/internal/cache"
	"github.com/psebaraj/pbay/products/internal/grpcserver"
	"github.com/psebaraj/pbay/products/pkg/database"

	protos "github.com/psebaraj/pbay/products/proto-gen"
)

func StartProductServer(cfg config.Config) {

	log.Printf("Server Running @ %s:%s", cfg.HTTPServer.Host, cfg.HTTPServer.Port)
	log.Printf("Cache: %s:%s", cfg.CacheConn.Address, cfg.CacheConn.Port)

	DB_Interface := database.Connect(cfg.DBConn.Dialect, cfg.DBConn.Address, cfg.DBConn.Port, cfg.DBConn.Username, cfg.DBConn.Database, cfg.DBConn.Password)
	defer DB_Interface.Close()

	database.AutoMigrateAll(cfg.DBConn.Dialect)


	//cache.ConnectRedisCache(cfg.CacheConn.Address, cfg.CacheConn.Port, cfg.CacheConn.Password)

	router := NewRouter()
	LoadSwagger(router) // swagger found @ localhost:8080/docs

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.HTTPServer.Port), LoadCors(router)))


}

func StartProductGRPCServer(cfg config.Config) {

	var opts []grpc.ServerOption

	gs := grpc.NewServer(opts...)

	protos.RegisterProductServer(gs, &grpcserver.Product{})

	// TODO: disable, just for testing with grpcurl
	reflection.Register(gs)

	//l, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.ProductGRPCServer.Port))
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.ProductGRPCServer.Address, cfg.ProductGRPCServer.Port))

	if err != nil {
		log.Fatal("Unable to listen", "error", err)
	}

	log.Printf("gRPC Server running @ %s:%s", cfg.ProductGRPCServer.Address, cfg.ProductGRPCServer.Port)
	gs.Serve(l)

}
