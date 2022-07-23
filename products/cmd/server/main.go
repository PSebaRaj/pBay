package main

import (
	"internal/config"
	"log"

	"github.com/psebaraj/pbay/products/pkg/app"
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

	app.FetchConfig(*cfg)

	go app.StartProductGRPCServer(*cfg)
	app.StartProductServer(*cfg)
}


