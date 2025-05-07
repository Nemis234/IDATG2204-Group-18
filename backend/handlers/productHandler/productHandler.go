package producthandler

import (
	cons "backend/constants"
	. "backend/structs"
	utility "backend/utility"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-sql-driver/mysql"
)

/*
ProductsHandler handles requests for a list of products.
It retrieves the products from the database and returns them in JSON format.

It is also used for creating new products, by means of a POST request.

It supports pagination and filtering based on various query parameters.
Pagination is done using the page number as a query parameter, and is mandatory.
The page number is used to determine the offset for the SQL query.

Supported query parameters:

	{
	- page 		(string)| mandatory	: The page number for pagination
	- search 	(string)| optional	: Filter products by name or description
	- category 	(string)| optional	: Filter products by category name
	- brand 	(string)| optional	: Filter products by brand name
	- priceMin 	(float64)| optional	: Filter products by minimum price
	- priceMax 	(float64)| optional	: Filter products by maximum price
	- pageLimit	(int64)| optional	: The number of products to return per page (default is 10)
	}

Example usage:

	Method: GET
	Route: /products?page=1&search=example&category=example1&brand=example2&priceMin=9.9&priceMax=100
	Response:
	HTTP code: 200 OK
	[
		{
			"product_id": "12345",
			"name": "example product",
			"description": "Product Description",
			"img_url": "https://example.com/image9.jpg",
			"price": 100.00,
			"stock_quantity": 50,
			"category_name": "example1",
			"brand_name": "example2",
		},
		{
			"product_id": "10",
			"name": "TechBrave Laptop Pro",
			"description": null,
			"img_url": null,
			"price": 15999.99,
			"stock_quantity": 40,
			"category_name": null,
			"brand_name": null,
	  	},
	]

When using the POST method, the request body should contain the product details in JSON format.
Mandatory fields cannot be null, and optional fields can be null.
The request body should include the following fields:

	{
	- product_id     (string)| optional	: The ID of the product
	- name           (string)| mandatory	: The name of the product
	- description    (string)| optional	: The description of the product
	- img_url        (string)| optional	: The URL of the product image
	- price          (float64)| mandatory	: The price of the product
	- stock_quantity (int64)| mandatory	: The quantity of the product in stock
	- category_name  (string)| optional	: The name of the category
	- brand_name     (string)| optional	: The name of the brand
	}

Example usage:

	Method: POST
	Route: /products
	Request body:
	{
		"name": "Product 1, example text",
		"description": "Description 1",
		"img_url": "https://example.com/image1.jpg",
		"price": 99.99,
		"stock_quantity": 50,
		"category_name": "Audio & Headphones",
		"brand_name": "TechBrave"
	}

	Response:
	HTTP code: 201 Created
	{
		"id" : "7c063863-2b6e-11f0-ad3b-58cdc90b0639"
	}
*/
func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("ProductsHandler called with method: ", r.Method)
	switch r.Method {
	case http.MethodGet:
		// Get query parameters from the URL
		// Example: /products?page=1&search=example&categoryID=1&brandID=2&priceMin=10&priceMax=100
		vars := r.URL.Query()

		pageStr := vars.Get("page")
		if pageStr == "" {
			http.Error(w, "Page number is required", http.StatusBadRequest)
			return
		}
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			log.Println("Error converting page number to int: ", err)
			http.Error(w, "Invalid page number", http.StatusBadRequest)
			return
		}
		if page < 1 {
			http.Error(w, "Page number must be greater than 0", http.StatusBadRequest)
			return
		}

		pageLimitStr := vars.Get("pageLimit")

		// Page limit and offset for pagination, default to 10 if not specified
		pageLimit := 10 // Default page limit if not specified
		if pageLimitStr != "" {
			var err error
			pageLimit, err = strconv.Atoi(pageLimitStr)
			if err != nil {
				log.Println("Error converting pageLimit to int: ", err)
				http.Error(w, "Invalid page limit", http.StatusBadRequest)
				return
			}
		}
		pageOffset := (page - 1) * pageLimit // Calculate the offset for pagination

		search := vars.Get("search")
		categoryID := vars.Get("categoryID")
		brandID := vars.Get("brandID")
		priceMin := vars.Get("priceMin")
		priceMax := vars.Get("priceMax")

		rawConditions := []struct {
			value string
			query string
		}{ // Using constants for column names
			{search, cons.PRODUCT_NAME + " LIKE ?"},
			{categoryID, cons.CATEGORYID + " = ?"},
			{brandID, cons.BRANDID + " = ?"},
			{priceMin, cons.PRODUCT_PRICE + " >= ?"},
			{priceMax, cons.PRODUCT_PRICE + " <= ?"},
		}
		var variables []any
		var conditions []string

		// Remove empty conditions
		for _, cond := range rawConditions {
			if cond.value != "" {
				variables = append(variables, cond.value)
				conditions = append(conditions, cond.query)
			}
		}
		// Build the SQL query with the conditions
		// Gets the default query for products
		query := cons.QueryProducts
		// Add the WHERE clause to the query if there are any conditions
		if len(conditions) > 0 {
			query += " WHERE " + conditions[0]
			// Add the rest of the conditions with AND, if any
			for _, statement := range conditions[1:] {
				query += " AND " + statement
			}
		}

		// Add the pagination to the query
		query += " LIMIT ? OFFSET ?"
		variables = append(variables, strconv.Itoa(pageLimit), strconv.Itoa(pageOffset))

		// Log the query for debugging purposes
		// Replace the placeholders with the actual values for logging
		log.Println("Executing query: ", fmt.Sprintf(strings.ReplaceAll(query, "?", "%s"), variables...))
		// Execute the query
		var products []Product
		cons.DB.Select(&products, query, variables...)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	case http.MethodPost:
		// Decode the request body into a Product struct
		var product Product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			log.Println("Error decoding request body: ", err)
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Validate mandatory fields
		if product.Name == "" || product.Price <= 0 || product.StockQuantity <= 0 {
			http.Error(w, "Missing mandatory fields", http.StatusBadRequest)
			return
		}

		if product.ProductID == "" {
			// Generate a new product ID
			err := cons.DB.Get(&product, "SELECT UUID() AS ProductID;")
			if err != nil {
				log.Println("Error generating product ID: ", err)
				http.Error(w, "Error generating product ID", http.StatusInternalServerError)
				return
			}
			log.Println("Product ID: ", product.ProductID)

		}

		id := product.ProductID

		_, err = cons.DB.NamedExec(cons.InsertProduct, product)
		if err != nil {
			log.Println("Error inserting product: ", err)
			http.Error(w, "Error inserting product", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		// Return the ID of the newly created product
		response := map[string]string{"id": id}
		json.NewEncoder(w).Encode(response)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

/*
ProductHandler handles GET requests for a single product.
It retrieves the product details from the database based on the provided ID in the URL path,
and returns the product details in JSON format.

It also supports PUT requests to update the product details.

It also supports PATCH requests to update the product details.

It also supports DELETE requests to delete a product.

Example usage:

	Method: GET
	Route: /products/12345
	Response:
	HTTP code: 200 OK
	{
		"product_id": "12345",
		"name": "example product",
		"description": "Product Description",
		"img_url": "https://example.com/image9.jpg",
		"price": 100.00,
		"stock_quantity": 50,
		"category_name": "example1",
		"brand_name": "example2",
	}

When using the PUT method, the request body should contain the product details in JSON format.
Mandatory fields cannot be null, while optional can be null.
The request body should include the following fields:

	{
	- name           (string)| mandatory	: The name of the product
	- description    (string)| optional	: The description of the product
	- img_url        (string)| optional	: The URL of the product image
	- price          (float64)| mandatory	: The price of the product
	- stock_quantity (int64)| mandatory	: The quantity of the product in stock
	- category_name  (string)| optional	: The name of the category
	- brand_name     (string)| optional	: The name of the brand
	}

Example usage:

	Method: PUT
	Route: /products/12345
	Request body:
	{
		"name": "Updated Product",
		"description": "Updated Description",
		"img_url": "https://example.com/updated_image.jpg",
		"price": 150.00,
		"stock_quantity": 30,
		"category_name": "Updated Category",
		"brand_name": "Updated Brand"
	}
	Response:
	HTTP code: 201 No Content

When using the PATCH method, the request body should contain the product details in JSON format.
Any amount of fields can be updated, but mandatory fields cannot be null, while optional can be null.
The request body can use any field(s) available in the PUT method.

When using the DELETE method, the product will be deleted from the database.
If the product is in a foreign key constraint, the delete will fail with a 409 Conflict error.
Example usage:

	Method: DELETE
	Route: /products/12345
	Response:
	HTTP code: 204 No Content
*/
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("ProductHandler called with method: ", r.Method)
	switch r.Method {
	case http.MethodGet:
		id := r.PathValue("id")
		if id == "" {
			http.Error(w, "ID is required", http.StatusBadRequest)
			return
		}

		var product Product

		cons.DB.Get(&product, cons.QueryProduct, id)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(product)
	case http.MethodPut:
		id := r.PathValue("id")
		if id == "" {
			http.Error(w, "ID is required", http.StatusBadRequest)
			return
		}
		// Decode the request body into a Product struct
		var product Product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			log.Println("Error decoding request body: ", err)
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if product.ProductID != "" && product.ProductID != id {
			http.Error(w, "Product ID cannot be changed", http.StatusBadRequest)
			return
		}

		// Validate mandatory fields
		if product.Name == "" || product.Price <= 0 || product.StockQuantity <= 0 {
			http.Error(w, "Missing mandatory fields", http.StatusBadRequest)
			return
		}

		product.ProductID = id
		// Update the product in the database
		_, err = cons.DB.NamedExec(cons.InsertProduct, product)
		if err != nil {
			log.Println("Error updating product: ", err)
			http.Error(w, "Error updating product", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)

	case http.MethodPatch:
		id := r.PathValue("id")
		if id == "" {
			http.Error(w, "ID is required", http.StatusBadRequest)
			return
		}

		// Decode the request body into a Product struct
		var product Product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			log.Println("Error decoding request body: ", err)
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if product.ProductID != "" && product.ProductID != id {
			http.Error(w, "Product ID cannot be changed", http.StatusBadRequest)
			return
		}

		query := "UPDATE " + cons.PRODUCTS_TABLE + " SET "

		setClause, args := utility.BuildUpdateQuery(product)
		if setClause == "" {
			http.Error(w, "No fields to update", http.StatusBadRequest)
			return
		}
		query += setClause + " WHERE " + cons.PRODUCTID + " = ?"
		args = append(args, id)

		// Log the query for debugging purposes
		// Replace the placeholders with the actual values for logging
		log.Println("Executing query: ", fmt.Sprintf(strings.ReplaceAll(query, "?", "%s"), args...))
		// Execute the query
		_, err = cons.DB.Exec(query, args...)
		if err != nil {
			log.Println("Error updating product: ", err)
			http.Error(w, "Error updating product", http.StatusInternalServerError)
			return
		}

	case http.MethodDelete:
		id := r.PathValue("id")
		if id == "" {
			http.Error(w, "ID is required", http.StatusBadRequest)
			return
		}

		// Needs to check if an admin is preforming the delete

		// Delete the product from the database
		result, err := cons.DB.Exec(cons.DeleteProduct, id)
		if err != nil {
			if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1451 {
				// Foreign key constraint error
				log.Println("Foreign key constraint error: ", err)
				http.Error(w, "Cannot delete product with existing references, probably due to a foreign key constraint", http.StatusConflict)
				return
			}
			log.Println("Error deleting product: ", err)
			http.Error(w, "Error deleting product", http.StatusInternalServerError)
			return
		}
		affected, err := result.RowsAffected()
		if err != nil {
			log.Println("Error getting affected rows: ", err)
			http.Error(w, "Error getting affected rows", http.StatusInternalServerError)
			return
		}
		if affected == 0 {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
		// Return a 204 No Content response
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
