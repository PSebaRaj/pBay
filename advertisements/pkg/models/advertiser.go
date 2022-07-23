package models

import "gorm.io/gorm"

type Advertiser struct {
	gorm.Model

	Name string
	InGoodStanding bool
	AdsDisplayed int
}
