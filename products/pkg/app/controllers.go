package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/psebaraj/pbay/products/internal/cache"
	"github.com/psebaraj/pbay/products/pkg/database"
	"github.com/psebaraj/pbay/products/pkg/models"
)

// swagger:route GET /product/{id} Product getProduct
//
//
// Produces:
// - application/json
//
// responses:
//   200: Product
//   404: nil
//
// controller: get singular (regular) product
// res: one product as JSON
func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var product models.Product

	// If the element is found in the redis cache, directly return it
	res := cache.GetFromCache(cache.REDIS, params["id"])
	if res != nil {
		fmt.Println("Cache hit")
		io.WriteString(w, res.(string))
		return
	}
	fmt.Println("Cache miss")
	database.DB.First(&product, params["id"])

	if product.Title == "" { // i.e. product not found
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&product) // still want to send as json
		return
	}

	// Set element in the redis cache before returning the result
	// "id" is what I query with
	cache.SetInCache(cache.REDIS, params["id"], product)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&product)
}

// swagger:route GET /products Product getProducts
//
//
// Produces:
// - application/json
//
// responses:
//   200: []Product
//   204: nil
//
// controller: get all (regular) products
// res: list of products as JSON
func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product

	database.DB.Find(&products)

	if len(products) == 0 {
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(&products) // still want to send as json
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&products)
}

// swagger:route POST /product Product createProduct
//
//
// Produces:
// - application/json
//
// Consumes:
// - application/json
//
// responses:
//   201: Product
//   507: nil
//
// controller: create singular (regular) product
// res: created (regular) product as JSON
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)

	createdProduct := database.DB.Create(&product)
	err := createdProduct.Error
	if err != nil {
		fmt.Printf("Error creating product %s, error: %s", product.Title, err)
		w.WriteHeader(http.StatusInsufficientStorage)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&createdProduct)
}

// swagger:route DELETE /product/{id} Product deleteProduct
//
//
// Produces:
// - application/json
//
// responses:
//   200: Product
//   404: nil
//   500: nil
//
// controller: delete singular (regular) product
// res: deleted product as JSON
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var product models.Product

	database.DB.First(&product, params["id"])
	if product.Title == "" {
		fmt.Printf("Error finding product %s before deletion", product.Title)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	deleted := database.DB.Delete(&product)

	err := deleted.Error
	if err != nil {
		fmt.Printf("Error deleting product %s, error: %s", product.Title, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// also delete from cache
	cache.DeleteFromCache(cache.REDIS, params["id"])

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&product)
}
