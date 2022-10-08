package app

import (
	"github.com/gorilla/mux"
	"html/template"
	"internal/database"
	"internal/models"
	"net/http"
)

func mainPage(resp http.ResponseWriter, req *http.Request) {
	Title := "Main page"
	List := database.SelectProducts(resp, req)

	tmpl, _ := template.ParseFiles("src/templates/index.html",
		"src/templates/header.html", "src/templates/footer.html")
	tmpl.ExecuteTemplate(resp, "index", struct {
		Title string
		List  []models.Product
	}{Title: Title, List: List})
}

func addProduct(resp http.ResponseWriter, req *http.Request) {
	Title := "Add product"

	tmpl, _ := template.ParseFiles("src/templates/add_product.html",
		"src/templates/header.html", "src/templates/footer.html")
	tmpl.ExecuteTemplate(resp, "add_product", struct{ Title string }{Title: Title})
}

func productsHandler(resp http.ResponseWriter, req *http.Request) {
	Title := "Product page"
	Row := database.GetProduct(resp, req)

	tmpl, _ := template.ParseFiles("src/templates/product.html",
		"src/templates/header.html", "src/templates/footer.html")

	tmpl.ExecuteTemplate(resp, "product", struct {
		Title string
		Row   models.Product
	}{Title: Title, Row: Row})

}

func Run() {

	//	gorilla/mux routing
	router := mux.NewRouter()
	router.HandleFunc("/", mainPage)
	router.HandleFunc("/add", addProduct)
	router.HandleFunc("/save_product", database.AddToDataBase)
	router.HandleFunc("/product/{id:[0-9]+}", productsHandler)

	http.ListenAndServe(":8181", router)
}
