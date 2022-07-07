package main

import (
	"fmt"
	"log"
	"net/http"
	"pbay/messages/routes"
	"pbay/products/models"

	"github.com/go-openapi/runtime/middleware"
)

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
	listen(9002)

}

func main() {
	Run()
}
