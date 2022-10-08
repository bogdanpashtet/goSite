package app

import (
	"github.com/gorilla/mux"
	"internal/database"
	"log"
	"net/http"
	"services"
)

func Run() {
	//	gorilla/mux routing
	router := mux.NewRouter()
	router.HandleFunc("/", services.MainPage)
	router.HandleFunc("/add", services.AddProduct)
	router.HandleFunc("/save_product", database.AddToDataBase)
	router.HandleFunc("/product/{id:[0-9]+}", services.ProductsHandler)

	err := http.ListenAndServe(":8181", router)
	if err != nil {
		log.Printf("Server launched")
	}
}
