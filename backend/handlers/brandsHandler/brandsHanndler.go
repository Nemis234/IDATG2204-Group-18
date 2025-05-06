package brandshandler

import (
	cons "backend/constants"
	. "backend/structs"
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
			"id": 1,
			"name": "Brand 1",
			"description": "Brand Description 1"
		},
		{
			"id": 2,
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

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

/*
BrandHandler handles requests for a single brand.
It retrieves the brand details from the database based on the provided ID in the URL path,
and returns the brand details in JSON format.

Example usage:

	Method: GET
	Route: /brands/1
	Response:
	{
		"id": 1,
		"name": "Brand 1",
		"description": "Brand Description 1"
	}
*/
func BrandHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("BrandHandler called with method: ", r.Method)
	switch r.Method {
	case http.MethodGet:
		brandID := r.PathValue("id")
		if brandID == "" {
			http.Error(w, "Brand ID is required", http.StatusBadRequest)
			return
		}
		var b Brand
		cons.DB.Get(&b, cons.QueryBrand, brandID)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(b)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
