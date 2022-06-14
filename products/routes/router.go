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
	r.HandleFunc("/products", controllers.GetProductByProductName).Methods("GET").Queries("product", "{product_name}")

	r.HandleFunc("/products/{product_name}", controllers.RemoveProduct).Methods("DELETE")
	//r.HandleFunc("/products/{product_name}", controllers.ModifyProduct).Methods("POST")

	return r
}
