package main

import (
	"backend/constants"
	handlers "backend/handlers"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	sqlx "github.com/jmoiron/sqlx"
)

var databaseType = "mysql"

func main() {
	dsn := os.Getenv("DSN") // Data Source Name (DSN) for MySQL connection
	if dsn == "" {
		// Replace with MySQL credentials
		dsn = "username:password@tcp(localhost:20)/dbname"
	}

	db, err := sqlx.Connect(databaseType, dsn)
	if err != nil {
		log.Println("Error connecting to database: ", err)
	}
	//defer db.Close()

	log.Println("Connected to database successfully")

	constants.DB = db // Assign the database connection to the handler package

	router := http.NewServeMux()
	// Note plural and singular
	router.HandleFunc("/products", handlers.ProductsHandler)
	router.HandleFunc("/products/{id}", handlers.ProductHandler)

	router.HandleFunc("/categories", handlers.CategoriesHandler)
	router.HandleFunc("/categories/{id}", handlers.CategoryHandler)

	router.HandleFunc("/brands", handlers.BrandsHandler)
	router.HandleFunc("/brands/{id}", handlers.BrandHandler)

	router.HandleFunc("/users", handlers.UsersHandler)
	router.HandleFunc("/users/{id}", handlers.UserHandler)
	router.HandleFunc("POST /users/login", handlers.LoginHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	log.Println("Server running on port ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
