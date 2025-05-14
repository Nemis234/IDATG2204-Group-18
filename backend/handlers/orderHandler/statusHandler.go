package orderHandler

import (
	cons "backend/constants"
	. "backend/structs"
	utility "backend/utility"
	"encoding/json"
	"log"
	"net/http"
)

/*
StatusHandler supports the following methods:
- GET: Fetches all order statuses. Requires admin privileges.
- POST: Creates a new order status. Requires admin privileges.

# GET
The function retrieves all order statuses from the database and returns them as a JSON response.

Example usage:

	Method: GET
	URL: /orderstatus
	Response:
	HTTP code: 200 OK
	[
		{
			"status_name": "Pending",
			"status_desc": "Order is pending"
		},
		{
			"status_name": "Shipped",
			"status_desc": "Order has been shipped"
		}
	]

# POST
The function creates a new order status in the database. The request body should contain the order status details in JSON format.

Example usage:

	Method: POST
	URL: /orderstatus
	Request body:
	{
		"status_name": "Delivered",
		"status_desc": "Order has been delivered"
	}
	Response:
	HTTP code: 201 Created
*/
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("OrderStatusHandler called with method: ", r.Method)
	switch r.Method {
	case http.MethodGet:
		// Check admin privileges
		if !utility.CheckPrivileges(r, w, nil) {
			return
		}

		var orderStatus []OrderStatus
		err := cons.DB.Select(&orderStatus, cons.QueryOrderStatus)
		if err != nil {
			if utility.CheckSQLErr(err, w) {
				return
			}
			log.Println("Error fetching order status: ", err)
			http.Error(w, "Error fetching order status", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(&orderStatus); err != nil {
			log.Println("Error writing JSON response: ", err)
			http.Error(w, "Error writing JSON response", http.StatusInternalServerError)
			return
		}

	case http.MethodPost:
		// Check admin privileges
		if !utility.CheckPrivileges(r, w, nil) {
			return
		}

		var orderStatus OrderStatus
		if err := json.NewDecoder(r.Body).Decode(&orderStatus); err != nil {
			log.Println("Error decoding JSON: ", err)
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}
		if orderStatus.StatusName == "" {
			log.Println("Name is required")
			http.Error(w, "Name is required", http.StatusBadRequest)
			return
		}

		_, err := cons.DB.NamedExec(cons.InsertOrderStatus, orderStatus)
		if err != nil {
			if utility.CheckSQLErr(err, w) {
				return
			}
			log.Println("Error inserting order status: ", err)
			http.Error(w, "Error inserting order status", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

/*
UpdateStatusHandler supports the following methods:
- PUT: Updates an existing order status. Requires admin privileges.
- DELETE: Deletes an existing order status. Requires admin privileges.

# PUT
The function updates an existing order status in the database. The request body should contain the updated order status details in JSON format.

Example usage:

	Method: PUT
	URL: /orderstatus/{statusName}
	Request body:
	{
		"status_desc": "Order has been shipped"
	}
	Response:
	HTTP code: 200 OK

# DELETE
The function deletes an existing order status from the database.

Example usage:

	Method: DELETE
	URL: /orderstatus/{statusName}
	Response:
	HTTP code: 204 No Content
*/
func UpdateStatusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("UpdateStatusHandler called with method: ", r.Method)
	switch r.Method {
	case http.MethodPut:
		statusName := r.PathValue("statusName")
		if statusName == "" {
			log.Println("Name is required")
			http.Error(w, "Name is required", http.StatusBadRequest)
			return
		}

		// Check admin privileges
		if !utility.CheckPrivileges(r, w, nil) {
			return
		}

		var orderStatus OrderStatus
		if err := json.NewDecoder(r.Body).Decode(&orderStatus); err != nil {
			log.Println("Error decoding JSON: ", err)
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}

		if orderStatus.StatusName != "" && orderStatus.StatusName != statusName {
			log.Println("Name cannot be changed")
			http.Error(w, "Name cannot be changed", http.StatusBadRequest)
			return
		}
		orderStatus.StatusName = statusName

		_, err := cons.DB.NamedExec(cons.UpdateOrderStatus, orderStatus)
		if err != nil {
			if utility.CheckSQLErr(err, w) {
				return
			}
			log.Println("Error updating order status: ", err)
			http.Error(w, "Error updating order status", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)

	case http.MethodDelete:
		statusName := r.PathValue("statusName")
		if statusName == "" {
			log.Println("Name is required")
			http.Error(w, "Name is required", http.StatusBadRequest)
			return
		}

		// Check admin privileges
		if !utility.CheckPrivileges(r, w, nil) {
			return
		}

		result, err := cons.DB.Exec(cons.DeleteOrderStatus, statusName)
		if err != nil {
			if utility.CheckSQLErr(err, w) {
				return
			}
			log.Println("Error deleting order status: ", err)
			http.Error(w, "Error deleting order status", http.StatusInternalServerError)
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
