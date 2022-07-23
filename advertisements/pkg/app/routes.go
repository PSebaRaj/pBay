
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
	router.HandleFunc("/sidebar", GetRandomSidebarAd).Methods("GET")
	router.HandleFunc("/base", GetRandomBaseAd).Methods("GET")
	router.HandleFunc("/topbar", GetRandomTopbarAd).Methods("GET")
	router.HandleFunc("/animated", GetRandomAnimatedAd).Methods("GET")
	router.HandleFunc("/siderbar/{id}", GetSpecificSidebarAd).Methods("GET")
	router.HandleFunc("/base/{id}", GetSpecificBaseAd).Methods("GET")
	router.HandleFunc("/topbar/{id}", GetSpecificTopbarAd).Methods("GET")
	router.HandleFunc("/animated/{id}", GetSpecificAnimatedAd).Methods("GET")
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
	r.Handle("/swagger.yaml", http.FileServer(http.Dir("../../swagger-gen")))
	// end swagger
}
