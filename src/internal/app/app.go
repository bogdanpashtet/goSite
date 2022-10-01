package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"internal/database"
	"net/http"
)

type httpHandler struct {
	fileName string
}

func (h httpHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	http.ServeFile(resp, req, h.fileName)
}

func productsHandler(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	response := fmt.Sprintf("Product %s", id)
	fmt.Fprint(resp, response)
}

func addProduct(resp http.ResponseWriter, req *http.Request) {
	name := req.FormValue("Name")
	price := req.FormValue("Price")

	// add name and price to database

	fmt.Fprintf(resp, "Имя: %s Возраст: %s", name, price)
}

func Run() {
	// connect to postgresql database
	database.Connection()

	//	gorilla/mux routing
	h1 := httpHandler{fileName: "src/web/static/index.html"}
	h2 := httpHandler{fileName: "src/web/static/add_product.html"}
	router := mux.NewRouter()
	router.HandleFunc("/", h1.ServeHTTP)
	router.HandleFunc("/add", h2.ServeHTTP)
	router.HandleFunc("/product/{id:[0-9]+}", productsHandler)
	http.Handle("/", router)

	http.ListenAndServe(":8181", nil)

	// close connection
	defer database.CloseConnection(database.Conn)
}
