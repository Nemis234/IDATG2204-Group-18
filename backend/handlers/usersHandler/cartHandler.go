package usershandler

import (
	cons "backend/constants"
	. "backend/structs"
	utility "backend/utility"
	"encoding/json"
	"log"
	"net/http"
)

/*
CartHandler supports the following methods:
  - GET: Get all cart items for a user
  - POST: Add a new item to the cart
  - DELETE: Delete all car items for a user

Only administrators and the user themselves can access any endpoint.

# GET

This method retrieves all cart items for a user. It requires the user ID to be passed in the URL path.

Example usage:

	Method: GET
	URL: /users/456/cart

	Response:
	HTTP Status: 200 OK
	Body:
	[
		{
			"ProductID": "123",
			"UserID": "456",
			"Quantity": 2
		},
		{
			"ProductID": "789",
			"UserID": "456",
			"Quantity": 1
		}
	]

# POST

This method adds a new item to the cart for a user. It requires the user ID to be passed in the URL path and the product ID and quantity to be passed in the request body.

Example usage:

	Method: POST
	URL: /users/456/cart
	Body:
	{
		"ProductID": "123",
		"Quantity": 2
	}

	Response:
	HTTP Status: 201 Created

# DELETE

This method deletes all cart items for a user. It requires the user ID to be passed in the URL path.

Example usage:

	Method: DELETE
	URL: /users/456/cart
	Response:
	HTTP Status: 204 No Content
*/
func CartHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		userID := r.PathValue("user_id")
		if userID == "" {
			http.Error(w, "User ID is required", http.StatusBadRequest)
			return
		}

		// Check user permission
		if !utility.CheckPrivileges(r, w, &userID) {
			return
		}

		var cartItems []CartItem
		err := cons.DB.Select(&cartItems, cons.QueryCart, userID)
		if err != nil {
			// If the error is a MySQL error, return
			if utility.CheckSQLErr(err, w) {
				return
			}
			log.Println("Error getting cart items: ", err)
			http.Error(w, "Error getting cart items", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(cartItems); err != nil {
			log.Println("Error encoding cart items to JSON: ", err)
			http.Error(w, "Error encoding cart items to JSON", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		userID := r.PathValue("user_id")
		if userID == "" {
			http.Error(w, "User ID is required", http.StatusBadRequest)
			return
		}
		// Check user permission
		if !utility.CheckPrivileges(r, w, &userID) {
			return
		}

		// Decode the request body into a Product struct
		var cartItem CartItem
		err := json.NewDecoder(r.Body).Decode(&cartItem)
		if err != nil {
			log.Println("Error decoding request body: ", err)
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		cartItem.UserID = userID

		_, err = cons.DB.NamedExec(cons.InsertCartItem, cartItem)
		if err != nil {
			// If the error is a MySQL error, return
			if utility.CheckSQLErr(err, w) {
				return
			}
			log.Println("Error inserting product: ", err)
			http.Error(w, "Error inserting product", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

	case http.MethodDelete:
		userID := r.PathValue("user_id")
		if userID == "" {
			http.Error(w, "User ID is required", http.StatusBadRequest)
			return
		}

		// Check user permission
		if !utility.CheckPrivileges(r, w, &userID) {
			return
		}

		result, err := cons.DB.Exec(cons.DeleteCart, userID)
		if err != nil {
			// If the error is a MySQL error, return
			if utility.CheckSQLErr(err, w) {
				return
			}
			log.Println("Error deleting cart: ", err)
			http.Error(w, "Error deleting cart", http.StatusInternalServerError)
			return
		}
		if utility.CheckDeleteResult(result, w) {
			return
		}
		log.Println("Cart for user: " + userID + " deleted successfully")
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

/*
CartItemHandler supports the following methods:
  - GET: Get a specific cart item for a user
  - PUT: Update a specific cart item for a user
  - DELETE: Delete a specific cart item for a user

Only administrators and the user themselves can access any endpoint.

# GET

This method retrieves a specific cart item for a user.
It requires the user ID and product ID to be passed in the URL path.

Example usage:

	Method: GET
	URL: /users/456/cart/123
	Response:
	HTTP Status: 200 OK
	Body:
	{
		"ProductID": "123",
		"UserID": "456",
		"Quantity": 2
	}

# PUT

This method updates a specific cart item for a user.
It requires the user ID and product ID to be passed in the URL path and the new quantity to be passed in the request body.
Example usage:

	Method: PUT
	URL: /users/456/cart/123
	Body:
	{
		"Quantity": 3
	}
	Response:
	HTTP Status: 200 OK

# DELETE

This method deletes a specific cart item for a user.
It requires the user ID and product ID to be passed in the URL path.

Example usage:

	Method: DELETE
	URL: /users/456/cart/123
	Response:
	HTTP Status: 204 No Content
*/
func CartItemHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		userID := r.PathValue("user_id")
		if userID == "" {
			http.Error(w, "user ID is required", http.StatusBadRequest)
			return
		}
		productID := r.PathValue("product_id")
		if productID == "" {
			http.Error(w, "product ID is required", http.StatusBadRequest)
			return
		}

		// Check user permission
		if !utility.CheckPrivileges(r, w, &userID) {
			return
		}

		var cartItem CartItem

		err := cons.DB.Get(&cartItem, cons.QueryCartItem, userID, productID)
		if err != nil {
			if utility.CheckSQLErr(err, w) {
				return
			}
			log.Println("Error getting cart item: ", err)
			http.Error(w, "Error getting cart item", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(cartItem); err != nil {
			log.Println("Error encoding cart items to JSON: ", err)
			http.Error(w, "Error encoding cart items to JSON", http.StatusInternalServerError)
			return
		}
	case http.MethodPut:
		userID := r.PathValue("user_id")
		if userID == "" {
			http.Error(w, "user ID is required", http.StatusBadRequest)
			return
		}
		productID := r.PathValue("product_id")
		if productID == "" {
			http.Error(w, "product ID is required", http.StatusBadRequest)
			return
		}

		// Check user permission
		if !utility.CheckPrivileges(r, w, &userID) {
			return
		}

		var cartItem CartItem
		if err := json.NewDecoder(r.Body).Decode(&cartItem); err != nil {
			log.Println("Error decoding cart item: ", err)
			http.Error(w, "Error decoding cart item", http.StatusBadRequest)
			return
		}
		cartItem.UserID = userID
		cartItem.ProductID = productID

		_, err := cons.DB.NamedExec(cons.UpdateCartItem, cartItem)
		if err != nil {
			if utility.CheckSQLErr(err, w) {
				return
			}
			log.Println("Error updating cart item: ", err)
			http.Error(w, "Error updating cart item", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)

	case http.MethodDelete:
		userID := r.PathValue("user_id")
		if userID == "" {
			http.Error(w, "user ID is required", http.StatusBadRequest)
			return
		}
		productID := r.PathValue("product_id")
		if productID == "" {
			http.Error(w, "product ID is required", http.StatusBadRequest)
			return
		}

		// Check user permission
		if !utility.CheckPrivileges(r, w, &userID) {
			return
		}
		result, err := cons.DB.Exec(cons.DeleteCartItem, userID, productID)
		if err != nil {
			if utility.CheckSQLErr(err, w) {
				return
			}
			log.Println("Error deleting cart item: ", err)
			http.Error(w, "Error deleting cart item", http.StatusInternalServerError)
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
