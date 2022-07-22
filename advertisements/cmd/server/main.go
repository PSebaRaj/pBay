package main

import (
    "net/http"
    "github.com/gorilla/mux"
)



func main() {
    r := mux.NewRouter()
    s := http.StripPrefix("/static/", http.FileServer(http.Dir("../../static/")))
    r.PathPrefix("/static/").Handler(s)
    http.Handle("/", r)
    err := http.ListenAndServe(":8081", nil)

	if err != nil {
		panic(err)
	}
}
