package categorieshandler

import (
	cons "backend/constants"
	. "backend/structs"
	"encoding/json"
	"log"
	"net/http"
)

/*
CategoriesHandler handles requests for a list of categories.
It retrieves all categories from the database and returns them in JSON format.

Example usage:

	Method: GET
	Route: /categories
	Response:
	[
		{
			"id": 1,
			"name": "Category 1",
			"description": "Category Description 1"
		},
		{
			"id": 2,
			"name": "Category 2",
			"description": "Category Description 2"
		}
	]
*/
func CategoriesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("CategoriesHandler called with method: ", r.Method)
	switch r.Method {
	case http.MethodGet:
		var categories []Category

		cons.DB.Select(&categories, cons.QueryCategories)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(categories)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

/*
CategoryHandler handles requests for a single category.
It retrieves the category details from the database based on the provided ID in the URL path,
and returns the category details in JSON format.

Example usage:

	Method: GET
	Route: /categories/1
	Response:
	{
		"id": 1,
		"name": "Category 1",
		"description": "Category Description 1"
	}
*/
func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("CategoryHandler called with method: ", r.Method)
	switch r.Method {
	case http.MethodGet:
		categoryID := r.PathValue("id")
		if categoryID == "" {
			http.Error(w, "Category ID is required", http.StatusBadRequest)
			return
		}

		var cat Category

		cons.DB.Get(&cat, cons.QueryCategory, categoryID)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cat)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
