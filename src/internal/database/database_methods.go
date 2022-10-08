package database

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
	"internal/models"
	"log"
	"net/http"
	"os"
)

const (
	username = "postgres"
	password = ""
	host     = "localhost"
	port     = "5432"
	database = "go_site"
)

func closeConnection(link *pgx.Conn) {
	err := link.Close(context.Background())
	if err != nil {
		fmt.Println("Connection closed.")
	}
}

func connection() *pgx.Conn {
	connectionUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", username, password, host, port, database)
	conn, err := pgx.Connect(context.Background(), connectionUrl)
	if err != nil {
		fmt.Printf("Unable to connect to database: %v\n", err)
		return nil
	} else {
		fmt.Printf("Connected to database!\n")
		return conn
	}
}

func AddToDataBase(resp http.ResponseWriter, req *http.Request) {
	Name := req.FormValue("Name")
	Price := req.FormValue("Price")

	conn := connection()
	defer closeConnection(conn)

	err := conn.QueryRow(context.Background(), "INSERT INTO products (name, price) VALUES ($1, $2);", Name, Price).Scan(&Name, &Price)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	http.Redirect(resp, req, "/", 301)
}

func SelectProducts(resp http.ResponseWriter, req *http.Request) []models.Product {
	conn := connection()
	defer closeConnection(conn)

	rows, err := conn.Query(context.Background(), "select * from products;")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	var rowSlice []models.Product
	for rows.Next() {
		var r models.Product
		err := rows.Scan(&r.Id, &r.Name, &r.Price)
		if err != nil {
			log.Fatal(err)
		}
		rowSlice = append(rowSlice, r)
	}

	return rowSlice
}

func GetProduct(resp http.ResponseWriter, req *http.Request) models.Product {
	vars := mux.Vars(req)
	id := vars["id"]

	var r models.Product

	conn := connection()
	defer closeConnection(conn)

	conn.QueryRow(context.Background(), "select * from products where product_id=$1;", id).Scan(&r.Id, &r.Name, &r.Price)

	return r
}
