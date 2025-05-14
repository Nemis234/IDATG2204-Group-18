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

/*
ProductReviewsHandler support these methods:
  - GET retrieve all reviews for a product.
  - POST create a new review for a product.

# GET

Retrieves all reviews for a product from the database.
Example usage:

	Method: GET
	Route: /products/0df01f83-a9a7-4afa-9b62-8d0bb9722849/reviews
	Response:
	HTTP code: 200 OK
	[
		{
			"product_id": "0df01f83-a9a7-4afa-9b62-8d0bb9722849",
			"user_id": "0df01f83-a9a7-4afa-9b62-8d0bb9722849",
			"rating": 5,
			"comment": "Great product!",
			"created_at": "2023-10-01T12:00:00Z"
		},
	]

# POST

Creates a new review for a product in the database.
Only a user that provides its own user_id can create a review.
An admin can create a review for any user.

Example usage:

	Method: POST
	Route: /products/0df01f83-a9a7-4afa-9b62-8d0bb9722849/reviews
	Request body:
	{
		"user_id": "0df01f83-a9a7-4afa-9b62-8d0bb9722849",
		"rating": 5,
		"comment": "Great product!"
	}
	Response:
	HTTP code: 201 Created
*/
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

/*
ReviewHandler support these methods:
  - GET retrieve a review for a product by user and product ID.
  - PUT update a review for a product by user and product ID.
  - PATCH partially update a review for a product by user and product ID.
  - DELETE delete a review for a product by user and product ID.

# GET

Retrieves a review for a product by user and product ID from the database.

Example usage:

	Method: GET

	Route: /products/0df01f83-a9a7-4afa-9b62-8d0bb9722849/reviews/0df01f83-a9a7-4afa-9b62-8d0bb9722849
	Response:
	HTTP code: 200 OK
	{
		"product_id": "0df01f83-a9a7-4afa-9b62-8d0bb9722849",
		"user_id": "0df01f83-a9a7-4afa-9b62-8d0bb9722849",
		"rating": 5,
		"comment": "Great product!",
		"created_at": "2023-10-01T12:00:00Z"
	}

# PUT

Updates a review for a product by user and product ID in the database.
Only the user that created the review  or an admin can update it.

Example usage:

	Method: PUT
	Route: /products/0df01f83-a9a7-4afa-9b62-8d0bb9722849/reviews/0df01f83-a9a7-4afa-9b62-8d0bb9722849
	Request body:
	{
		"rating": 4,
		"comment": "Good product!"
	}
	Response:
	HTTP code: 200 OK

# PATCH
Partially updates a review for a product by user and product ID in the database.
Only the user that created the review  or an admin can update it.

All fields are optional, and all fields used in the PUT method are available.

# DELETE
Deletes a review for a product by user and product ID from the database.
Only the user that created the review  or an admin can delete it.

Example usage:

	Method: DELETE
	Route: /products/0df01f83-a9a7-4afa-9b62-8d0bb9722849/reviews/0df01f83-a9a7-4afa-9b62-8d0bb9722849
	Response:
	HTTP code: 204 No Content
*/
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
