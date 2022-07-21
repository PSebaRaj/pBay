package api

import (
	"fmt"


	//"github.com/psebaraj/pbay/products/internal/cache"
	"github.com/psebaraj/pbay/products/pkg/database"
	"github.com/psebaraj/pbay/products/pkg/models"
)

//
// Functions for other services to call
// Note: endpoint controllers are in the ./app
//

func GetProduct(productID uint) (product models.Product, err error) {
	database.DB.First(&product, productID)

	if product.Title == "" { // i.e. product not found
		return models.Product{}, fmt.Errorf("Error finding valid product %d", productID)
	}

	// Set element in the redis cache before returning the result
	// "id" is what I query with
	//cache.SetInCache(cache.REDIS, params["id"], product)

	return product, nil

}


func GetProducts() (products []models.Product, err error) {
	database.DB.Find(&products)

	if len(products) == 0 {
		return []models.Product{}, fmt.Errorf("Error: could not find any product")
	}

	return products, nil

}


func CreateProduct(product models.Product) (models.Product, error) {
	createdProduct := database.DB.Create(&product)

	err := createdProduct.Error
	if err != nil {
		return models.Product{}, fmt.Errorf("Error creating product %s, error: %s", product.Title, err)
	}

	return product, nil
}


func DeleteProduct(productID uint) (product models.Product, err error) {
	database.DB.First(&product, productID)
	if product.Title == "" {
		return models.Product{}, fmt.Errorf("Error finding valid product %d before deletion", productID)
	}

	deleted := database.DB.Delete(&product)

	err = deleted.Error
	if err != nil {
		return product, fmt.Errorf("Error deleting product %s, error: %s", product.Title, err)
	}

	// also delete from cache
	// cache.DeleteFromCache(cache.REDIS, params["id"])

	return product, nil
}
