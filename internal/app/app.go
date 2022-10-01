package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

//type httpHandler struct {
//	fileName string
//}
//
//func (h httpHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
//	http.ServeFile(resp, req, h.fileName)
//}

func productsHandler(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	response := fmt.Sprintf("Product %s", id)
	fmt.Fprint(resp, response)
}

func Run() {

	var router = mux.NewRouter()
	router.HandleFunc("/products/{id:[0-9]+}", productsHandler)
	http.Handle("/", router)

	//h1 := httpHandler{fileName: "static/index.html"}
	//
	//http.Handle("/", h1)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
