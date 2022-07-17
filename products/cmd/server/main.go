package main

import (
	"internal/config"
	"log"
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

	log.Printf("HTTP HOST: %s", cfg.HTTPServer.Host)
	log.Printf("HTTP PORT: %s", cfg.HTTPServer.Port)
	log.Printf("DB Dialect: %s", cfg.DBConn.Dialect)
	log.Printf("Cache: %s:%s", cfg.CacheConn.Address, cfg.CacheConn.Port)

}
