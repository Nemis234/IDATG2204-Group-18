package handler

import (
	"encoding/json"
	"net/http"
)

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		query := "SELECT id, title, description FROM products"

		id := r.URL.Query().Get("id")
		if id == "" {
			query = "SELECT id, title, description FROM products WHERE id = " + id
		}

		rows, err := db.Query(query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var products ListProducts

		for rows.Next() {
			var p Product
			if err := rows.Scan(&p.ProductID,
				&p.Name,
				&p.Description,
				&p.Price,
				&p.StockQuantity,
				&p.Brand.ID,
				&p.Category.ID); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			products.Products = append(products.Products, p)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)

	case http.MethodPost:
		http.Error(w, "POST method not implemented", http.StatusNotImplemented)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
