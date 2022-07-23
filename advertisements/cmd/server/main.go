package main

import (
	"internal/config"
	"log"

	"github.com/psebaraj/pbay/advertisements/pkg/app"
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

	app.StartAdvertisementServer(*cfg)
}

//
// package main
//
// import (
//     "net/http"
//     "github.com/gorilla/mux"
// )
//
//
//
// func main() {
//     r := mux.NewRouter()
//     s := http.StripPrefix("/static/", http.FileServer(http.Dir("../../static/")))
//     r.PathPrefix("/static/").Handler(s)
//     http.Handle("/", r)
//     err := http.ListenAndServe(":8081", nil)
//
// 	if err != nil {
// 		panic(err)
// 	}
// }
