package models

import (
	"time"
)

type Product struct {
	Product_ID    uint32     `gorm:"primary_key;auto_increment" json:"id"`
	Product_Name  string     `gorm:"size:35;not null;unique_index" json:"product_name,omitempty"`
	Seller        string     `gorm:"size:35;not null" json:"seller_name,omitempty"`
	Current_Price uint32     `json:"current_price" required:"true"`
	Buy_Now_Price uint32     `json:"buy_now_price"`
	Num_Of_Bids   uint32     `json:"num_of_bids"`
	Image         string     `json:"image"`
	CreatedAt     *time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	ExpiresAt     *time.Time `gorm:"default:current_timestamp()" json:"expires_at"`
}

// need to change the Image to an actual image later

func CreateProduct(product Product) (interface{}, error) {
	db := Connect()
	defer db.Close()

	rs := db.Create(&product)
	return rs.Value, rs.Error
}

// return all currentlly listed products
func GetProducts() []Product {
	db := Connect()
	defer db.Close()

	var products []Product
	db.Order("id asc").Find(&products)
	if len(products) == 0 { // to throw empty error code later
		return nil
	}
	return products
}

// returns a single product
func GetProductByProductName(p_name string) Product {
	db := Connect()
	defer db.Close()

	var product Product
	db.Where("product_name = ?", p_name).Find(&product)
	return product
}

// returns array of products, as seller might have multiple listings
func GetProductsBySeller(s_name string) []Product {
	db := Connect()
	defer db.Close()

	var products []Product
	db.Where("seller_name = ?", s_name).Order("id asc").Find(&products)
	if len(products) == 0 { // throw empty http error code later
		return nil
	}
	return products
}
