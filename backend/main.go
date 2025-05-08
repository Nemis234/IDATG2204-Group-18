package main

import (
	"backend/constants"
	brandshandler "backend/handlers/brandsHandler"
	categorieshandler "backend/handlers/categoriesHandler"
	orderHandler "backend/handlers/orderHandler"
	producthandler "backend/handlers/productHandler"
	usershandler "backend/handlers/usersHandler"
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
		// This works with XAMPP when running locally
		dsn = "root:@tcp(127.0.0.1:3306)/idatg2204"
	}
	log.Println("Attempting connection with DSN: ", dsn)
	dsn += "?parseTime=true" // Add to parse time into time.Time correctly
	db, err := sqlx.Connect(databaseType, dsn)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	defer db.Close()

	log.Println("Connected to database successfully")

	constants.DB = db // Assign the database connection to the handler package

	router := http.NewServeMux()
	// Note plural and singular in handler names
	router.HandleFunc("/brands", brandshandler.BrandsHandler)
	router.HandleFunc("/brands/{id}", brandshandler.BrandHandler)

	router.HandleFunc("/categories", categorieshandler.CategoriesHandler)
	router.HandleFunc("/categories/{id}", categorieshandler.CategoryHandler)

	router.HandleFunc("/orders", orderHandler.OrdersHandler)
	router.HandleFunc("/orders/{id}", orderHandler.OrderHandler)
	router.HandleFunc("/orders/{id}/payment", orderHandler.PaymentHandler)
	router.HandleFunc("/orders/{id}/items", orderHandler.OrderItemsHandler)
	router.HandleFunc("/orders/{id}/items/{item_id}", orderHandler.OrderItemHandler)
	router.HandleFunc("/orders/status", orderHandler.StatusHandler)

	router.HandleFunc("/products", producthandler.ProductsHandler)
	router.HandleFunc("/products/{id}", producthandler.ProductHandler)
	router.HandleFunc("/products/{id}/reviews", producthandler.ReviewHandler)

	router.HandleFunc("/users", usershandler.UsersHandler)
	router.HandleFunc("/users/{id}", usershandler.UserHandler)
	// Only POST method is allowed for login
	router.HandleFunc("POST /users/login", usershandler.LoginHandler)
	router.HandleFunc("/users/{id}/member", usershandler.MemberHandler)
	router.HandleFunc("/user/{id}/cart", usershandler.CartHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	log.Println("Server running on port ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
