package database

import (
	"fmt"
	"github.com/psebaraj/pbay/products/pkg/models"
)

// Make migrations to the database if they haven't been made already
func AutoMigrateAll(dialect string) {

	err = DB.AutoMigrate(&models.Product{})
	// err = DB.AutoMigrate(&models.Product{}).Error

	if err != nil {
		fmt.Printf("Unable to AutoMigrate model(s) %s to %s DB", "Product", dialect)
		panic(err)
	}

}
