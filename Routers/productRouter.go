package Routers

import (
	"Natural/Controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func ProductRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/products", Controllers.GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", Controllers.GetOneProductById).Methods("GET")
	router.HandleFunc("/products", Controllers.CreateProducts).Methods("POST")
	router.HandleFunc("/products/{id}", Controllers.DeletePost).Methods("DELETE")
	router.HandleFunc("/products/{id}", Controllers.UpdatePost).Methods("PUT")

	http.Handle("/", router)
}
