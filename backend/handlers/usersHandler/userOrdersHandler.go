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
UserOrdersHandler suppoerts the following methods:
  - GET: Returns all orders for a specific user. The user ID is passed as a URL parameter.

# GET

Example usage:

	  	Method: GET
	  	Path: /users/1234/orders
		Response:
		HTTP Status: 200 OK
		[
			{
				"order_id": 1,
				"user_id": 1234,
				"order_date": "2023-10-01",
				"total_amount": 100.00,
				"status": "Shipped",
				"items": none
			},
			{
				"order_id": 2,
				"user_id": 1234,
				"order_date": "2023-10-02",
				"total_amount": 50.00,
				"status": "Pending",
				"items": none
			}
		]
*/
func UserOrdersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		userID := r.PathValue("user_id")
		if userID == "" {
			http.Error(w, "User ID is required", http.StatusBadRequest)
			return
		}
		// Check admin privileges
		if !utility.CheckPrivileges(r, w, &userID) {
			return
		}

		var orders []Order
		err := cons.DB.Select(&orders, cons.QueryOrdersByUser, userID)
		if err != nil {
			if utility.CheckSQLErr(err, w) {
				return
			}
			log.Println("Error fetching orders: ", err)
			http.Error(w, "Error fetching orders", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(orders); err != nil {
			log.Println("Error encoding orders to JSON: ", err)
			http.Error(w, "Error encoding orders to JSON", http.StatusInternalServerError)
			return
		}
	}
}
