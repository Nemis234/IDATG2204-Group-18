package producthandler

import (
	cons "backend/constants"
	. "backend/structs"
	utility "backend/utility"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func ProductReviewsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productID := r.PathValue("product_id")
		if productID == "" {
			http.Error(w, "Product ID is required", http.StatusBadRequest)
			return
		}

		var reviews []Review
		err := cons.DB.Select(&reviews, cons.QueryReviewsByProduct, productID)
		if err != nil {
			if utility.CheckSQLErr(err, w) {
				return
			}
			http.Error(w, "Error fetching reviews", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(reviews); err != nil {
			log.Println("Error encoding response: ", err)
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}

	case http.MethodPost:
		productID := r.PathValue("product_id")
		if productID == "" {
			http.Error(w, "Product ID is required", http.StatusBadRequest)
			return
		}
		var review Review
		if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
			http.Error(w, "Error decoding request body", http.StatusBadRequest)
			return
		}

		// Check user privileges
		if !utility.CheckPrivileges(r, w, &review.UserID) {
			return
		}

		review.ProductID = productID
		_, err := cons.DB.NamedExec(cons.InsertReview, review)
		if err != nil {
			if utility.CheckSQLErr(err, w) {
				return
			}
			http.Error(w, "Error inserting review", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func ReviewHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productID := r.PathValue("product_id")
		if productID == "" {
			http.Error(w, "Product ID is required", http.StatusBadRequest)
			return
		}
		userID := r.PathValue("user_id")
		if userID == "" {
			http.Error(w, "User ID is required", http.StatusBadRequest)
			return
		}
		var review Review
		err := cons.DB.Get(&review, cons.QueryReview, productID, userID)
		if err != nil {
			if utility.CheckSQLErr(err, w) {
				return
			}
			http.Error(w, "Error fetching review", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(review); err != nil {
			log.Println("Error encoding response: ", err)
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}

	case http.MethodPut:
		productID := r.PathValue("product_id")
		if productID == "" {
			http.Error(w, "Product ID is required", http.StatusBadRequest)
			return
		}
		userID := r.PathValue("user_id")
		if userID == "" {
			http.Error(w, "User ID is required", http.StatusBadRequest)
			return
		}

		// Check user privileges
		if !utility.CheckPrivileges(r, w, &userID) {
			return
		}

		var review Review
		if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
			http.Error(w, "Error decoding request body", http.StatusBadRequest)
			return
		}
		review.ProductID = productID
		review.UserID = userID
		_, err := cons.DB.NamedExec(cons.UpdateReview, review)
		if err != nil {
			if utility.CheckSQLErr(err, w) {
				return
			}
			http.Error(w, "Error updating review", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)

	case http.MethodPatch:
		productID := r.PathValue("product_id")
		if productID == "" {
			http.Error(w, "Product ID is required", http.StatusBadRequest)
			return
		}
		userID := r.PathValue("user_id")
		if userID == "" {
			http.Error(w, "User ID is required", http.StatusBadRequest)
			return
		}

		// Check user privileges
		if !utility.CheckPrivileges(r, w, &userID) {
			return
		}

		var review ReviewPatch
		if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
			http.Error(w, "Error decoding request body", http.StatusBadRequest)
			return
		}
		if review.ProductID != "" && review.ProductID != productID {
			http.Error(w, "Product ID cannot be changed", http.StatusBadRequest)
			return
		}
		if review.UserID != "" && review.UserID != userID {
			http.Error(w, "User ID cannot be changed", http.StatusBadRequest)
			return
		}

		review.ProductID = productID
		review.UserID = userID

		query := "UPDATE " + cons.REVIEWS_TABLE + " SET "

		setClause, args := utility.BuildUpdateQuery(review)
		if setClause == "" {
			http.Error(w, "No fields to update", http.StatusBadRequest)
			return
		}
		query += setClause + " WHERE " + cons.PRODUCT_ID + " = ? AND " + cons.USER_ID + " = ?"
		// Add the product_id and user_id to the args map
		args = append(args, productID, userID)

		// Log the query for debugging
		log.Println("Executing query: ", fmt.Sprintf(strings.ReplaceAll(query, "?", "%s"), args...))

		_, err := cons.DB.Exec(query, args...)
		if err != nil {
			if utility.CheckSQLErr(err, w) {
				return
			}
			http.Error(w, "Error updating review", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)

	case http.MethodDelete:
		productID := r.PathValue("product_id")
		if productID == "" {
			http.Error(w, "Product ID is required", http.StatusBadRequest)
			return
		}
		userID := r.PathValue("user_id")
		if userID == "" {
			http.Error(w, "User ID is required", http.StatusBadRequest)
			return
		}

		// Check user privileges
		if !utility.CheckPrivileges(r, w, &userID) {
			return
		}

		_, err := cons.DB.Exec(cons.DeleteReview, productID, userID)
		if err != nil {
			if utility.CheckSQLErr(err, w) {
				return
			}
			http.Error(w, "Error deleting review", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
