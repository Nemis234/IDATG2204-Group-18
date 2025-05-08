package categorieshandler

import (
	cons "backend/constants"
	. "backend/structs"
	utility "backend/utility"
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
			"name": "Category 1",
			"description": "Category Description 1"
		},
		{
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

	case http.MethodPost:
		var c Category
		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		if c.Name == "" {
			http.Error(w, "Category name is required", http.StatusBadRequest)
			return
		}
		_, err := cons.DB.NamedExec(cons.InsertCategory, c)
		if err != nil {
			if utility.CheckSQLErr(err, w) {
				return
			}
			http.Error(w, "Failed to insert category", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

/*
CategoryHandler handles requests for a single category.
It retrieves the category details from the database based on the provided name in the URL path,
and returns the category details in JSON format.

Example usage:

	Method: GET
	Route: /categories/category 1
	Response:
	{
		"name": "Category 1",
		"description": "Category Description 1"
	}
*/
func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("CategoryHandler called with method: ", r.Method)
	switch r.Method {
	case http.MethodGet:
		categoryName := r.PathValue("id")
		if categoryName == "" {
			http.Error(w, "Category Name is required", http.StatusBadRequest)
			return
		}

		var cat Category

		cons.DB.Get(&cat, cons.QueryCategory, categoryName)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cat)

	case http.MethodPut:
		categoryName := r.PathValue("id")
		if categoryName == "" {
			http.Error(w, "Category name is required", http.StatusBadRequest)
			return
		}
		var c Category
		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		if c.Name != "" && c.Name != categoryName {
			http.Error(w, "Category name cannot be changed", http.StatusBadRequest)
			return
		}
		_, err := cons.DB.NamedExec(cons.UpdateCategory, c)
		if err != nil {
			if utility.CheckSQLErr(err, w) {
				return
			}
			http.Error(w, "Failed to update category", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	case http.MethodDelete:
		categoryName := r.PathValue("id")
		if categoryName == "" {
			http.Error(w, "Category name is required", http.StatusBadRequest)
			return
		}
		_, err := cons.DB.Exec(cons.DeleteCategory, categoryName)
		if err != nil {
			if utility.CheckSQLErr(err, w) {
				return
			}
			http.Error(w, "Failed to delete category", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
