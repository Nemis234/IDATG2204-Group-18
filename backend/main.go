package main

import (
	handlers "backend/handlers"
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

	handlers.SetDB(db) // Assign the database connection to the handler package

	router := http.NewServeMux()
	// Note plural and singular
	router.HandleFunc("/products/{page}", handlers.ProductsHandler)
	router.HandleFunc("/products/{id}", handlers.ProductHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	log.Println("Server running on port ", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
