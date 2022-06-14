package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pbay/products/models"
	"pbay/products/utils"

	"github.com/gorilla/mux"
)

// swagger:response productsResponse
// A list of all products
type productsResponse struct {
	// All products in the system
	// in: body
	Body []models.Product
}

// swagger:response productResponse
// A singular product
type productResponse struct {
	// Singular product
	// in: body
	Body models.Product
}

// swagger:route POST /products ProductAdd addProduct
// Adds a product to the database
//
// Produces:
// - application/json
//
// Consumes:
// - application/json
//
// responses:
//   201: productResponse
//   422: error
//   500: error

// Function to create a product
func PostProduct(w http.ResponseWriter, r *http.Request) {
	body := utils.BodyParser(r)
	var user models.Product
	err := json.Unmarshal(body, &user)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	rs, err := models.CreateProduct(user)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.ToJson(w, rs, http.StatusCreated)
}

// swagger:route GET /products ProductData getProducts
// Returns all products in system (id, product_name, seller_name, current_price, buy_now_price, num_of_bids, created_at, expires_at)
//
// Produces:
// - application/json
//
// responses:
//   200: productsResponse
//   204: error

// Function to create a product
func GetProducts(w http.ResponseWriter, r *http.Request) {
	products := models.GetProducts()

	if products == nil {
		utils.ToJson(w, nil, http.StatusNoContent)
		return
	}

	utils.ToJson(w, products, http.StatusFound)
}

// swagger:route GET /products/{seller_name} ProductData getProducts
// Returns all products in system listed under the queried seller
//
// Produces:
// - application/json
//
// responses:
//   200: productsResponse
//   204: error

// to see all of seller's listings
func GetProductsBySeller(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	products := models.GetProductsBySeller(vars["seller_name"])

	if products == nil {
		utils.ToJson(w, nil, http.StatusNoContent)
		return
	}

	utils.ToJson(w, products, http.StatusFound)
}

// NOT IN SWAGGER DOC BECAUSE STILL FIXING FUNCTIONALITY OF SEARCH BY PRODUCT NAME

// for search functionality
// takes in json data
func GetProductByProductName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Fprintf(w, "input: %s\n", vars["product_name"])
	product := models.GetProductByProductName(vars["product_name"])

	//	if product == nil {
	//		utils.ToJson(w, nil, http.StatusNoContent)
	//		return
	//	}

	utils.ToJson(w, product, http.StatusFound)

}

// swagger:route DELETE /products/{product_name} ProductData removeProduct
// Removes product given name of the product
//
// responses:
//   200: nil
//   400: error

// to be used to delete, not modify, a product
func RemoveProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	err := models.RemoveProduct(vars["product_name"])

	if err != nil {
		utils.ToJson(w, nil, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
