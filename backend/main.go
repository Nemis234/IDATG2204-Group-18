package main

import (
	"backend/constants"
	brandshandler "backend/handlers/brandsHandler"
	categorieshandler "backend/handlers/categoriesHandler"
	orderHandler "backend/handlers/orderHandler"
	producthandler "backend/handlers/productHandler"
	usershandler "backend/handlers/usersHandler"
	utility "backend/utility"
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

	//Creating the JWT Key
	constants.JwtKey, err = utility.GetJwtKey()
	if err != nil {
		log.Fatal("Failed to initialize JWT key:", err)
	}

	constants.DB = db // Assign the database connection to the handler package

	router := http.NewServeMux()
	// Note plural and singular in handler names
	router.HandleFunc("/brands", brandshandler.BrandsHandler)
	router.HandleFunc("/brands/{id}", brandshandler.BrandHandler)

	router.HandleFunc("/categories", categorieshandler.CategoriesHandler)
	router.HandleFunc("/categories/{id}", categorieshandler.CategoryHandler)

	router.HandleFunc("/orders", orderHandler.OrdersHandler)
	router.HandleFunc("/orders/{order_id}", orderHandler.OrderHandler)
	router.HandleFunc("/orders/{order_id}/items", orderHandler.OrderItemsHandler)
	router.HandleFunc("/orders/{order_id}/items/{item_id}", orderHandler.OrderItemHandler)
	router.HandleFunc("/orders/{order_id}/payment", orderHandler.PaymentsHandler)
	router.HandleFunc("/orders/{order_id}/payment/{payment_id}", orderHandler.PaymentHandler)

	router.HandleFunc("/orderstatus", orderHandler.StatusHandler)
	router.HandleFunc("/orderstatus/{statusName}", orderHandler.StatusHandler)

	router.HandleFunc("/products", producthandler.ProductsHandler)
	router.HandleFunc("/products/{product_id}", producthandler.ProductHandler)
	router.HandleFunc("/products/{product_id}/reviews", producthandler.ProductReviewsHandler)
	router.HandleFunc("/products/{product_id}/reviews/{user_id}", producthandler.ReviewHandler)

	router.HandleFunc("/users", usershandler.UsersHandler)
	router.HandleFunc("/users/{user_id}", usershandler.UserHandler)
	router.HandleFunc("/users/{user_id}/member", usershandler.MemberHandler)
	router.HandleFunc("/users/{user_id}/cart", usershandler.CartHandler)
	router.HandleFunc("/users/{user_id}/cart/{product_id}", usershandler.CartItemHandler)
	router.HandleFunc("/users/{user_id}/reviews", usershandler.UserReviewsHandler)
	router.HandleFunc("/users/{user_id}/reviews/{product_id}", usershandler.ReviewHandler)
	// Only POST method is allowed for login
	router.HandleFunc("POST /users/login", usershandler.LoginHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	log.Println("Server running on port ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
