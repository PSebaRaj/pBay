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
	"net/http"

	"github.com/psebaraj/pbay/advertisements/pkg/database"

)

func StartAdvertisementServer(cfg config.Config) {

	log.Printf("Advertisement Server Running @ %s:%s", cfg.HTTPServer.Host, cfg.HTTPServer.Port)

	DB_Interface := database.Connect(cfg.DBConn.Dialect, cfg.DBConn.Address, cfg.DBConn.Port, cfg.DBConn.Username, cfg.DBConn.Database, cfg.DBConn.Password)
	defer DB_Interface.Close()

	database.AutoMigrateAll(cfg.DBConn.Dialect)

	// Note: no server-side caching here! All caching for ads (jpeg/png/gif) should be done client-side

	router := NewRouter()
	LoadSwagger(router)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.HTTPServer.Port), LoadCors(router)))


}
