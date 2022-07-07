package routes

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/ads", controllers.PostAd).Methods("POST")

	r.HandleFunc("/ads/{ad_name}", controllers.GetSpecificAd).Methods("GET")
	r.HandleFunc("ads/random", controllers.GetRandomAd).Methods("GET")

	return r
}
