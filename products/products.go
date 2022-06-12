package main

import (
	"fmt"
	"log"
	"net/http"
	"pbay/products/models"
	"pbay/products/routes"

	"github.com/go-openapi/runtime/middleware"
)

// for Seller, use the NAME / any non-hashed/encrypted identifying detail

func listen(p int) {
	port := fmt.Sprintf(":%d", p)
	fmt.Printf("Listening Port %s...\n", port)
	r := routes.NewRouter()

	// swagger
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	r.Handle("/docs", sh)
	r.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
	// end swagger

	log.Fatal(http.ListenAndServe(port, routes.LoadCors(r)))
}

func Run() {
	db := models.Connect()

	// check if MySQL table exists, if not, create import
	if !db.HasTable(&models.Product{}) {
		db.Debug().CreateTable(&models.Product{})

	}
	db.Close()
	listen(9001)

}

func main() {
	Run()
}
