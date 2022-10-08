package services

import (
	"html/template"
	"internal/database"
	"internal/models"
	"log"
	"net/http"
)

func MainPage(resp http.ResponseWriter, req *http.Request) {
	Title := "Main page"
	List := database.SelectProducts(resp, req)

	tmpl, _ := template.ParseFiles("src/templates/index.html",
		"src/templates/header.html", "src/templates/footer.html")
	err := tmpl.ExecuteTemplate(resp, "index", struct {
		Title string
		List  []models.Product
	}{Title: Title, List: List})
	if err != nil {
		log.Println("Error: can't execute template.")
	}
}

func AddProduct(resp http.ResponseWriter, req *http.Request) {
	Title := "Add product"

	tmpl, _ := template.ParseFiles("src/templates/add_product.html",
		"src/templates/header.html", "src/templates/footer.html")
	err := tmpl.ExecuteTemplate(resp, "add_product", struct{ Title string }{Title: Title})
	if err != nil {
		log.Println("Error: can't execute template.")
	}
}

func ProductsHandler(resp http.ResponseWriter, req *http.Request) {
	Title := "Product page"
	Row := database.GetProductById(resp, req)

	tmpl, _ := template.ParseFiles("src/templates/product.html",
		"src/templates/header.html", "src/templates/footer.html")

	err := tmpl.ExecuteTemplate(resp, "product", struct {
		Title string
		Row   models.Product
	}{Title: Title, Row: Row})
	if err != nil {
		log.Println("Error: can't execute template.")
	}
}
