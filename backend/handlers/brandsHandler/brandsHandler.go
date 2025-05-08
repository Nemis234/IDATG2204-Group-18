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
BrandsHandler handles requests for a list of brands.
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
*/
func BrandsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("BrandsHandler called with method: ", r.Method)
	switch r.Method {
	case http.MethodGet:
		var brands []Brand

		cons.DB.Select(&brands, cons.QueryBrands)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(brands)
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
BrandHandler handles requests for a single brand.
It retrieves the brand details from the database based on the provided name in the URL path,
and returns the brand details in JSON format.

Example usage:

	Method: GET
	Route: /brands/1
	Response:
	{
		"name": "Brand 1",
		"description": "Brand Description 1"
	}
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
		json.NewEncoder(w).Encode(b)
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
