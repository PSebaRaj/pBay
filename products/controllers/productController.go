package controllers

import (
	"encoding/json"
	"net/http"
	"pbay/products/models"
	"pbay/products/utils"
)

// swagger:response productsResponse
// A list of all products
type productsResponse struct {
	// All users in the system
	// in: body
	Body []models.Product
}

// swagger:response productResponse
// A singular product
type productResponse struct {
	// Singular user
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

// Function to create a user
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
// Returns all users in system (name, email, password, created_at, updated_at)
//
// Produces:
// - application/json
//
// responses:
//   200: productsResponse
//   204: error

// Function to create a user
func GetProducts(w http.ResponseWriter, r *http.Request) {
	products := models.GetProducts()

	if products == nil {
		utils.ToJson(w, nil, http.StatusNoContent)
		return
	}

	utils.ToJson(w, products, http.StatusFound)
}
