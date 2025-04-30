package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

/*
ProductsHandler handles requests for a list of products.
It retrieves the products from the database and returns them in JSON format.

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

Example usage:

	Method: GET
	Route: /products?page=1&search=example&categoryID=1&brandID=2&priceMin=10&priceMax=100
	Response:
	{
	"products": [
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
	}
*/
func ProductsHandler(w http.ResponseWriter, r *http.Request) {
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
		search := vars.Get("search")
		categoryID := vars.Get("categoryID")
		brandID := vars.Get("brandID")
		priceMin := vars.Get("priceMin")
		priceMax := vars.Get("priceMax")

		pageLimitStr := vars.Get("pageLimit")
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

		query := queryProducts
		// Build the WHERE clause based on the provided query parameters
		var whereStatements []string

		// Using constants for column names

		// !!!!!
		// Sprintf NEEDS TO BE REPLACED WITH PREPARED STATEMENTS
		// !!!!!
		if search != "" {
			//whereStatements = append(whereStatements, fmt.Sprintf("%s LIKE %s OR %s LIKE %s", NAME, search, DESCRIPTION, search))
			whereStatements = append(whereStatements, fmt.Sprintf("%s LIKE %s", NAME, search))
		}
		if categoryID != "" {
			whereStatements = append(whereStatements, fmt.Sprintf("%s = %s", CATEGORYID, categoryID))
		}
		if brandID != "" {
			whereStatements = append(whereStatements, fmt.Sprintf("%s = %s", BRANDID, brandID))
		}
		if priceMin != "" {
			whereStatements = append(whereStatements, fmt.Sprintf("%s >= %s", PRICE, priceMin))
		}
		if priceMax != "" {
			whereStatements = append(whereStatements, fmt.Sprintf("%s <= %s", PRICE, priceMax))
		}
		// Add the WHERE clause to the query if there are any conditions
		if len(whereStatements) > 0 {
			query += " WHERE " + whereStatements[0]
			for _, statement := range whereStatements[1:] {
				query += " AND " + statement
			}
		}

		query += " LIMIT ? OFFSET ?"

		// Execute the query
		log.Println("Executing query: ", query)
		productRows, err := db.Query(query, pageLimit, pageOffset)
		if err != nil {
			log.Println("Error querying products: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer productRows.Close()

		var products ListProducts

		for productRows.Next() {
			// Scan the product values into a Product struct
			p, err := GetProduct(productRows)
			if err != nil {
				log.Println("Error getting product values: ", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			products.Products = append(products.Products, p)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

/*
ProductHandler handles requests for a single product.
It retrieves the product details from the database based on the provided ID in the URL path,
and returns the product details in JSON format.

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
	switch r.Method {
	case http.MethodGet:
		id := r.PathValue("id")
		if id == "" {
			http.Error(w, "ID is required", http.StatusBadRequest)
			return
		}

		productRows, err := db.Query(queryProduct, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("Error querying product: ", err)
			return
		}
		defer productRows.Close()

		if !productRows.Next() {
			log.Println("Product not found: ", productRows.Err())
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}

		p, err := GetProduct(productRows)
		if err != nil {
			log.Println("Error getting product values: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(p)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
