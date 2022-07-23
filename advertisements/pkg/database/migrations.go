package database

import (
	"fmt"
	"github.com/psebaraj/pbay/advertisements/pkg/models"
)

// Make migrations to the database if they haven't been made already
func AutoMigrateAll(dialect string) {

	err = DB.AutoMigrate(&models.Advertiser{}, &models.Advertisement{})

	if err != nil {
		fmt.Printf("Unable to AutoMigrate model(s) %s, %s to %s DB", "Advertiser", "Advertisement", dialect)
		panic(err)
	}

}
