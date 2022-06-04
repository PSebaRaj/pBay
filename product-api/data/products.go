package data

import (
	"encoding/json"
	"io"
	"time"
)

// CreatedOn, UpdatedOn, and DeletedOn are not send through the API,
// just used w/in backend for tracking

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Seller      string  `json:"seller"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	// *** do any manipulation before conversion to JSON here ***

	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Seller:      "Starbucks",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Seller:      "Atticus",
		Description: "Short and strong coffee without milk",
		Price:       3.00,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
