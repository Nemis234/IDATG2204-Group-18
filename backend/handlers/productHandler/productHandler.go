package producthandler

import (
	cons "backend/constants"
	. "backend/structs"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

/*
ProductsHandler handles requests for a list of products.
It retrieves the products from the database and returns them in JSON format.

It is also used for creating new products, by means of a POST request.

It supports pagination and filtering based on various query parameters.
Pagination is done using the page number as a query parameter, and is mandatory.
The page number is used to determine the offset for the SQL query.

Supported query parameters:
  - page: The page number for pagination (mandatory)
  - search: Filter products by name or description
  - categoryID: Filter products by category ID
  - brandID: Filter products by brand ID
  - priceMin: Filter products by minimum price
  - priceMax: Filter products by maximum price
  - pageLimit: The number of products to return per page (default is 10)

Example usage:

	Method: GET
	Route: /products?page=1&search=example&categoryID=1&brandID=2&priceMin=10&priceMax=100
	Response:
	[
		{
			"product_id": 1,
			"name": "Product 1, example text",
			"description": "Description 1",
			"price": 100,
			"stock_quantity": 50,
			"category": {
				"id": 1,
				"name": "Category 1",
				"description": "Category Description 1"
			},
			"brand": {
				"id": 1,
				"name": "Brand 2",
				"description": "Brand Description 2"
			}
		},
	]

	Method: POST
	Route: /products
	Request body:
	{
		"product_id": 1,
		"name": "Product 1, example text",
		"description": "Description 1",
		"price": 100,
		"stock_quantity": 50,
		"category": {
			"id": 1,
			"name": "Category 1",
			"description": "Category Description 1"
		},
		"brand": {
			"id": 1,
			"name": "Brand 2",
			"description": "Brand Description 2"
		}
	}

	Response:
	{
		"id" : 3
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
		http.Error(w, "Method not implemented", http.StatusNotImplemented)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

/*
ProductHandler handles requests for a single product.
It retrieves the product details from the database based on the provided ID in the URL path,
and returns the product details in JSON format.

It can also handle PUT requests to update the product details.

Example usage:

	Method: GET
	Route: /products/12345
	Response:
	{
		"product_id": 12345,
		"name": "Product Name",
		"description": "Product Description",
		"price": 100,
		"stock_quantity": 50,
		"category": {
			"id": 1,
			"name": "Category Name",
			"description": "Category Description"
			},
		"brand": {
			"id": 1,
			"name": "Brand Name",
			"description": "Brand Description"
			}
	}
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
		http.Error(w, "Method not implemented", http.StatusNotImplemented)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
