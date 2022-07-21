package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Title string
	Description string
	StartingBidPrice float32
	CurrentBidPrice float32
	BuyNowPrice float32
}
