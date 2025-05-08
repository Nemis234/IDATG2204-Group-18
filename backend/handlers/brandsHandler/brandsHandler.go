package brandshandler

import (
	cons "backend/constants"
	. "backend/structs"
	utility "backend/utility"
	"encoding/json"
	"log"
	"net/http"
)

/*
BrandsHandler supports these HTTP methods:
- GET: Retrieves a list of all brands from the database.
- POST: Inserts a new brand into the database.

# GET

GET handles requests for a list of brands.
It retrieves all brands from the database and returns them in JSON format.

Example usage:

	Method: GET
	Route: /brands
	Response:
	[
		{
			"name": "Brand 1",
			"description": "Brand Description 1"
		},
		{
			"name": "Brand 2",
			"description": "Brand Description 2"
		}
	]

# POST

POST handles requests to insert a new brand into the database.
It expects a JSON payload with the brand details.
Example usage:

	Method: POST
	Route: /brands
	Request Body:
	{
		"name": "Brand 3",
		"description": "Brand Description 3"
	}
	Response:
	HTTP Status: 201 Created
*/
func BrandsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("BrandsHandler called with method: ", r.Method)
	switch r.Method {
	case http.MethodGet:
		var brands []Brand

		cons.DB.Select(&brands, cons.QueryBrands)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(brands); err != nil {
			log.Println("Error encoding brands to JSON: ", err)
			http.Error(w, "Error encoding brands to JSON", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		var b Brand
		if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		if b.Name == "" {
			http.Error(w, "Brand name is required", http.StatusBadRequest)
			return
		}
		_, err := cons.DB.NamedExec(cons.InsertBrand, b)
		if err != nil {
			if utility.CheckSQLErr(err, w) {
				return
			}
			http.Error(w, "Failed to insert brand", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

/*
BrandHandler supports these HTTP methods:
- GET: Retrieves details of a specific brand based on the name in the URL path.
- PUT: Updates the details of a specific brand.
- DELETE: Deletes a specific brand from the database.

# GET

GET handles requests for a single brand.
It retrieves the brand details from the database based on the provided name in the URL path,
and returns the brand details in JSON format.

Example usage:

	Method: GET
	Route: /brands/Brand 1
	Response:
	{
		"name": "Brand 1",
		"description": "Brand Description 1"
	}

# PUT

PUT handles requests to update a specific brand in the database.
It expects a JSON payload with the brand details.
Example usage:

	Method: PUT
	Route: /brands/Brand 1
	Request Body:
	{
		"description": "Updated Brand Description"
	}
	Response:
	HTTP Status: 200 OK

# DELETE

DELETE handles requests to delete a specific brand from the database.
Example usage:

	Method: DELETE
	Route: /brands/Brand 1
	Response:
	HTTP Status: 204 No Content
*/
func BrandHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("BrandHandler called with method: ", r.Method)
	switch r.Method {
	case http.MethodGet:
		brandName := r.PathValue("id")
		if brandName == "" {
			http.Error(w, "Brand name is required", http.StatusBadRequest)
			return
		}
		var b Brand
		cons.DB.Get(&b, cons.QueryBrand, brandName)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(b); err != nil {
			log.Println("Error encoding brand to JSON: ", err)
			http.Error(w, "Error encoding brand to JSON", http.StatusInternalServerError)
			return
		}
	case http.MethodPut:
		brandName := r.PathValue("id")
		if brandName == "" {
			http.Error(w, "Brand name is required", http.StatusBadRequest)
			return
		}
		var b Brand
		if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		if b.Name != "" && b.Name != brandName {
			http.Error(w, "Brand name cannot be changed", http.StatusBadRequest)
			return
		}
		b.Name = brandName

		_, err := cons.DB.NamedExec(cons.UpdateBrand, b)
		if err != nil {
			if utility.CheckSQLErr(err, w) {
				return
			}
			http.Error(w, "Failed to update brand", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)

	case http.MethodDelete:
		brandName := r.PathValue("id")
		if brandName == "" {
			http.Error(w, "Brand name is required", http.StatusBadRequest)
			return
		}
		_, err := cons.DB.Exec(cons.DeleteBrand, brandName)
		if err != nil {
			if utility.CheckSQLErr(err, w) {
				return
			}
			http.Error(w, "Failed to delete brand", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
