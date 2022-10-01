package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"internal/database"
	"internal/models"
	"net/http"
)

func mainPage(resp http.ResponseWriter, req *http.Request) {
	Title := "Главная страница"
	List := database.SelectProducts(resp, req)

	tmpl, _ := template.ParseFiles("src/web/static/index.html",
		"src/web/static/header.html", "src/web/static/footer.html")
	tmpl.ExecuteTemplate(resp, "index", struct {
		Title string
		List  []models.Product
	}{Title: Title, List: List})
}

func addProduct(resp http.ResponseWriter, req *http.Request) {
	Title := "Добавить товар"

	tmpl, _ := template.ParseFiles("src/web/static/add_product.html",
		"src/web/static/header.html", "src/web/static/footer.html")
	tmpl.ExecuteTemplate(resp, "add_product", struct{ Title string }{Title: Title})
}

func productsHandler(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	response := fmt.Sprintf("Product %s", id)
	fmt.Fprint(resp, response)
}

func Run() {
	//	gorilla/mux routing
	router := mux.NewRouter()
	router.HandleFunc("/", mainPage)
	router.HandleFunc("/add", addProduct)
	router.HandleFunc("/save_product", database.AddToDataBase)
	router.HandleFunc("/product/{id:[0-9]+}", productsHandler)
	http.Handle("/", router)

	http.ListenAndServe(":8181", nil)
}
