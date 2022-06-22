package routes

import (
	"net/http"
	"pbay/messages/pkg/websocket"
	"pbay/products/controllers"

	"github.com/gorilla/mux"
)

func serveWS(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {

}

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/products", controllers.GetProducts).Methods("GET")

	return r
}
