package handlers

import (
	"log"
	"net/http"
	"online-store/product-api/data"
	"regexp"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// handle GET requests
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// handle POST requests
	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	// handle PUT requests
	if r.Method == http.MethodPut {
		// id in the URI
		reg := regexp.MustCompile(`/([0-9]+)`)
		group := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(group) != 1 {
			http.Error(rw, "Invalid URI: more than one id", http.StatusBadRequest)
			return
		}

		if len(group[0]) != 2 {
			http.Error(rw, "Invalid URI: more than one capture group", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(group[0][1])

		if err != nil {
			http.Error(rw, "Invalid URI: unable to convert to type INT", http.StatusBadRequest)
			return
		}
		//p.l.Println("got id:", id)

		p.updateProducts(rw, r, id)

	}

	// catch other methods
	rw.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handling GET")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handling POST")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	//p.l.Printf("prod: %#v\n", prod)
	data.AddProduct(prod)

}

func (p *Products) updateProducts(rw http.ResponseWriter, r *http.Request, id int) {
	p.l.Println("Handling PUT")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(prod, id)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Unable to update product", http.StatusInternalServerError)
		return
	}

}
