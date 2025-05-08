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
CategoriesHandler supports these HTTP methods:
- GET: Retrieves a list of all categories from the database.
- POST: Inserts a new category into the database.

# GET

GET handles requests for a list of categories.
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

# POST

POST handles requests to insert a new category into the database.
It expects a JSON payload with the category details.
Example usage:

	Method: POST
	Route: /categories
	Request Body:
	{
		"name": "Category 3",
		"description": "Category Description 3"
	}
	Response:
	HTTP Status: 201 Created
*/
func CategoriesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("CategoriesHandler called with method: ", r.Method)
	switch r.Method {
	case http.MethodGet:
		var categories []Category

		cons.DB.Select(&categories, cons.QueryCategories)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(categories); err != nil {
			log.Println("Error encoding categories to JSON: ", err)
			http.Error(w, "Error encoding categories to JSON", http.StatusInternalServerError)
			return
		}

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
CategoryHandler supports these HTTP methods:
- GET: Retrieves details of a specific category based on the name in the URL path.
- PUT: Updates the details of a specific category.
- DELETE: Deletes a specific category based on the name in the URL path.

# GET

GET handles requests for a single category.
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

# PUT

PUT handles requests to update a specific category in the database.
It expects a JSON payload with the category details.
Example usage:

	Method: PUT
	Route: /categories/category 1
	Request Body:
	{
		"description": "Updated Category Description"
	}
	Response:
	HTTP Status: 200 OK

# DELETE

DELETE handles requests to delete a specific category from the database.
It retrieves the category details from the database based on the provided name in the URL path,
and deletes the category from the database.
Example usage:

	Method: DELETE
	Route: /categories/category 1
	Response:
	HTTP Status: 204 No Content
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
		if err := json.NewEncoder(w).Encode(cat); err != nil {
			log.Println("Error encoding category to JSON: ", err)
			http.Error(w, "Error encoding category to JSON", http.StatusInternalServerError)
			return
		}

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
		result, err := cons.DB.Exec(cons.DeleteCategory, categoryName)
		if err != nil {
			if utility.CheckSQLErr(err, w) {
				return
			}
			http.Error(w, "Failed to delete category", http.StatusInternalServerError)
			return
		}
		if utility.CheckDeleteResult(result, w) {
			return
		}
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
