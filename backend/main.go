package main

import (
	handler "backend/handler"
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var databaseType = "mysql"

func main() {
	dsn := os.Getenv("DSN") // Data Source Name (DSN) for MySQL connection
	if dsn == "" {
		// Replace with MySQL credentials
		dsn = "username:password@tcp(localhost:20)/dbname"
	}

	db, err := sql.Open(databaseType, dsn)
	if err != nil {
		log.Fatal("Failed to open sql database with error: ", err)
	}
	defer db.Close()

	handler.SetDB(db) // Assign the database connection to the handler package

	router := http.NewServeMux()
	router.HandleFunc("/products", handler.ProductsHandler)
	router.HandleFunc("/products/{id}", handler.ProductsHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	log.Println("Server running on port ", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
