package routes

import (
	"pbay/products/controllers"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	r.HandleFunc("/products", controllers.PostProduct).Methods("POST")

	r.HandleFunc("/products/seller/{seller_name}", controllers.GetProductsBySeller).Methods("GET")
	r.HandleFunc("/products/search", controllers.GetProductByProductName).Methods("GET") // takes in product_name as JSON

	return r
}
