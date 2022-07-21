package main

import (
	"internal/config"
	"log"
	"net/http"

	"github.com/psebaraj/pbay/products/internal/cache"
	"github.com/psebaraj/pbay/products/pkg/app"
	"github.com/psebaraj/pbay/products/pkg/database"
)

func main() {
	// Generate our config based on the config supplied
	// by the user in the flags
	cfgPath, err := config.ParseFlags()
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := config.NewConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server Running @ %s:%s", cfg.HTTPServer.Host, cfg.HTTPServer.Port)
	log.Printf("Cache: %s:%s", cfg.CacheConn.Address, cfg.CacheConn.Port)

	DB_Interface := database.Connect(cfg.DBConn.Dialect, cfg.DBConn.Address, cfg.DBConn.Port, cfg.DBConn.Username, cfg.DBConn.Database, cfg.DBConn.Password)
	defer DB_Interface.Close()

	database.AutoMigrateAll(cfg.DBConn.Dialect)


	cache.ConnectRedisCache(cfg.CacheConn.Address, cfg.CacheConn.Port, cfg.CacheConn.Password)

	router := app.NewRouter()
	app.LoadSwagger(router) // swagger found @ localhost:8080/docs

	log.Fatal(http.ListenAndServe(":9001", app.LoadCors(router)))


}


