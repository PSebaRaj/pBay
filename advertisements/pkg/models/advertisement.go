package models

import "gorm.io/gorm"

type Advertisement struct {
	gorm.Model

	Name string
	Type AdType
	AdvertiserID int
}

type AdType uint

const (
	SIDEBAR AdType = iota
	BASE
	TOPBAR
	ANIMATED
	end
)


