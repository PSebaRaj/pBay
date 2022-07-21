package app

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// used to initialize a new Gorilla Mux Router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	productRoutes(router)

	return router
}

// product routes
func productRoutes(router *mux.Router) {
	router.HandleFunc("/products", GetProducts).Methods("GET")
	router.HandleFunc("/product/{id}", GetProduct).Methods("GET")
	router.HandleFunc("/product", CreateProduct).Methods("POST")
	router.HandleFunc("/product/{id}", DeleteProduct).Methods("DELETE")
}

// CORS for Gorilla Mux
func LoadCors(r http.Handler) http.Handler {
	headers := handlers.AllowedHeaders([]string{"X-Request", "Content-Type", "AuthorizationX-Request", "Content-Type", "Authorization"}) // X-Request", "Content-Type", "AuthorizationX-Request", "Content-Type", "Authorization
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"}) // should change in the future when client is ready
	return handlers.CORS(headers, methods, origins)(r)

}

// Loading swagger into /docs
func LoadSwagger(r *mux.Router) {
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	fmt.Println("Swagger can be found /docs")

	r.Handle("/docs", sh)
	r.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
	// end swagger
}
