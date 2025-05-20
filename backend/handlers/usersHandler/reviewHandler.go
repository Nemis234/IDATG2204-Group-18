package usershandler

import (
	cons "backend/constants"
	producthandler "backend/handlers/productHandler"
	. "backend/structs"
	utility "backend/utility"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

/*
UserReviewsHandler supports the following methods:
- GET: Fetches all reviews for a specific user.
- POST: Creates a new review for a specific user.

# GET
The function retrieves all reviews for a specific user from the database and returns them as a JSON response.

Example usage:

	Method: GET
	URL: /users/{user_id}/reviews
	Response:

	HTTP code: 200 OK
	[
		{
		"user_id": "0df01f83-a9a7-4afa-9b62-8d0bb9722849",
		"product_id": "12345678-1234-5678-1234-567812345678",
		"comment": "Great product!",
		"rating": 5,
		"post_date": "2023-10-01T12:00:00Z"
		}
	]

# POST

The function creates a new review for a specific user in the database. The request body should contain the review details in JSON format.
Only the user who made the review or an admin can access this endpoint.

Example usage:

	Method: POST
	URL: /users/{user_id}/reviews
	Request body:

	{
		"product_id": "12345678-1234-5678-1234-567812345678",
		"comment": "Great product!",
		"rating": 5
	}

	Response:
	HTTP code: 201 Created
*/
func UserReviewsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		userID := r.PathValue("user_id")
		if userID == "" {
			http.Error(w, "User ID is required", http.StatusBadRequest)
			return
		}

		var reviews []Review
		err := cons.DB.Select(&reviews, cons.QueryReviewsByUser, userID)
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
		userID := r.PathValue("user_id")
		if userID == "" {
			http.Error(w, "User ID is required", http.StatusBadRequest)
			return
		}
		var review Review
		if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
			http.Error(w, "Error decoding request body", http.StatusBadRequest)
			return
		}

		// Check user privileges
		if !utility.CheckPrivileges(r, w, &userID) {
			return
		}

		review.UserID = userID
		review.PostDate = time.Now()

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

// See productHandler.ReviewHandler()
var ReviewHandler = producthandler.ReviewHandler
