package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx"
	"os"
)

const (
	username = "postgres"
	password = ""
	host     = "localhost"
	port     = "5432"
	database = "go_site"
)

var Conn *pgx.Conn

func CloseConnection(link *pgx.Conn) {
	err := link.Close(context.Background())
	if err != nil {
		fmt.Println("Connection closed.")
	}
}

func Connection() {
	connectionUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", username, password, host, port, database)
	Conn, err := pgx.Connect(context.Background(), connectionUrl)
	if err != nil {
		fmt.Printf("Unable to connect to database: %v\n", err)
		os.Exit(1) // error output
	} else {
		fmt.Printf("Connected to database!\nConnection: %v", Conn)
	}
}
