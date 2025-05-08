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
OrderItemsHandler supports the following methods:
- GET: Fetches all order items for a specific order by ID.
- POST: Adds a new order item to an existing order.

# GET

GET handles the retrieval of all order items for a specific order by ID from the database.
It queries the database for all order items associated with the specified order ID and returns them as a JSON response.

Example usage:

	Method: GET
	Route: /orders/12345/items
	Response:
	Http Status: 200 OK
	[
		{
			"order_id": "12345",
			"product_id": "67890",
			"quantity": 2
		},
		{
			"order_id": "12345",
			"product_id": "54321",
			"quantity": 1
		}
	]

# POST

POST handles the addition of a new order item to an existing order.
It expects a JSON payload with the order item details, in JSON format.

The order ID is specified in the URL, and the order item details are provided in the request body.
The order ID cannot be changed.
Mandatory fields cannot be null, while optional can be null.

The request body should contain the following fields:

	{
	- product_id	(string)| mandatory	: The ID of the product being ordered.
	- quantity	(int)| mandatory	: The quantity of the product in the order.
	}

Example usage:

	Method: POST
	Route: /orders/12345/items
	Request Body:
	{
		"product_id": "67890",
		"quantity": 2
	}
	Response:
	Http Status: 201 Created
*/
func OrderItemsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("OrderItemHandler called with method : ", r.Method)
	switch r.Method {
	case http.MethodGet:
		orderID := r.PathValue("order_id")

		var orderItems []OrderItem
		err := cons.DB.Select(&orderItems, cons.QueryOrderItemsByID, orderID)
		if err != nil {
			if utility.CheckSQLErr(err, w) {
				return
			}
			log.Println("Error fetching order items: ", err)
			http.Error(w, "Error fetching order items", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(orderItems); err != nil {
			log.Println("Error encoding order items to JSON: ", err)
			http.Error(w, "Error encoding order items to JSON", http.StatusInternalServerError)
			return
		}

	case http.MethodPost:
		orderID := r.PathValue("order_id")
		if orderID == "" {
			log.Println("Order ID is required for POST request")
			http.Error(w, "Order ID is required", http.StatusBadRequest)
			return
		}
		var orderItem OrderItem
		if err := json.NewDecoder(r.Body).Decode(&orderItem); err != nil {
			log.Println("Error decoding order item: ", err)
			http.Error(w, "Error decoding order item", http.StatusBadRequest)
			return
		}

		log.Println("Order :", orderItem)

		// check for mandatory fields
		if orderItem.ProductID == "" || orderItem.Quantity < 0 {
			log.Println("Mandatory fields are missing")
			http.Error(w, "Mandatory fields are missing", http.StatusBadRequest)
			return
		}
		orderItem.OrderID = orderID

		_, err := cons.DB.NamedExec(cons.InsertOrderItem, orderItem)
		if err != nil {
			if utility.CheckSQLErr(err, w) {
				return
			}
			log.Println("Error inserting order item: ", err)
			http.Error(w, "Error inserting order item", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

/*
OrderItemHandler supports the following methods:
- GET: Fetches a specific order item by order ID and item ID.
- PUT: Updates an existing order item.
- DELETE: Deletes an existing order item.

# GET

GET handles the retrieval of a specific order item by order ID and item ID from the database.
It queries the database for the order item with the specified order ID and item ID and returns it as a JSON response.

Example usage:

	Method: GET
	Route: /orders/12345/items/67890
	Response:
	Http Status: 200 OK
	{
		"order_id": "12345",
		"product_id": "67890",
		"quantity": 2
	}

# PUT

PUT handles the update of an existing order item in the database.
It expects a JSON payload with the order item details, in JSON format.

The order ID and item ID are specified in the URL, and the order item details are provided in the request body.
The order ID and product ID cannot be changed.

Mandatory fields cannot be null, while optional can be null.

The request body should contain the following fields:

	{
	- quantity	(int)| mandatory	: The quantity of the product in the order.
	}

Example usage:

	Method: PUT
	Route: /orders/12345/items/67890
	Request Body:
	{
		"quantity": 2
	}
	Response:
	Http Status: 200 OK

# DELETE

DELETE handles the deletion of an existing order item in the database.
It expects the order ID and item ID to be specified in the URL.
The order item is deleted from the database, and a 204 No Content response is returned.

Example usage:

	Method: DELETE
	Route: /orders/12345/items/67890
	Response:
	Http Status: 204 No Content
*/
func OrderItemHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("OrderItemHandler called with method : ", r.Method)
	switch r.Method {
	case http.MethodGet:
		orderID := r.PathValue("order_id")
		itemID := r.PathValue("item_id")

		if orderID == "" {
			log.Println("Order ID is required for GET request")
			http.Error(w, "Order ID is required", http.StatusBadRequest)
			return
		}
		if itemID == "" {
			log.Println("Item ID is required for GET request")
			http.Error(w, "Item ID is required", http.StatusBadRequest)
			return
		}

		var orderItem OrderItem
		err := cons.DB.Get(&orderItem, cons.QueryOrderItemByID, orderID, itemID)
		if err != nil {
			if utility.CheckSQLErr(err, w) {
				return
			}
			log.Println("Error fetching order item: ", err)
			http.Error(w, "Error fetching order item", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(orderItem); err != nil {
			log.Println("Error encoding order item to JSON: ", err)
			http.Error(w, "Error encoding order item to JSON", http.StatusInternalServerError)
			return
		}

	case http.MethodPut:
		orderID := r.PathValue("order_id")
		itemID := r.PathValue("item_id")
		if orderID == "" {
			log.Println("Order ID is required for PUT request")
			http.Error(w, "Order ID is required", http.StatusBadRequest)
			return
		}

		if itemID == "" {
			log.Println("Product ID is required for PUT request")
			http.Error(w, "Product ID is required", http.StatusBadRequest)
			return
		}

		var orderItem OrderItem
		if err := json.NewDecoder(r.Body).Decode(&orderItem); err != nil {
			log.Println("Error decoding order item: ", err)
			http.Error(w, "Error decoding order item", http.StatusBadRequest)
			return
		}
		if orderItem.OrderID != "" && orderItem.OrderID != orderID {
			http.Error(w, "Order ID cannot be changed", http.StatusBadRequest)
			return
		}
		if orderItem.ProductID != "" && orderItem.ProductID != itemID {
			http.Error(w, "Product ID cannot be changed", http.StatusBadRequest)
			return
		}

		orderItem.OrderID = orderID
		orderItem.ProductID = itemID

		// check for mandatory fields
		if orderItem.Quantity < 0 {
			log.Println("Mandatory fields are missing")
			http.Error(w, "Mandatory fields are missing", http.StatusBadRequest)
			return
		}
		// Update the order item in the database
		_, err := cons.DB.NamedExec(cons.UpdateOrderItem, orderItem)
		if err != nil {
			if utility.CheckSQLErr(err, w) {
				return
			}
			log.Println("Error updating order item: ", err)
			http.Error(w, "Error updating order item", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)

	case http.MethodDelete:
		orderID := r.PathValue("order_id")
		itemID := r.PathValue("item_id")
		if orderID == "" {
			log.Println("Order ID is required for DELETE request")
			http.Error(w, "Order ID is required", http.StatusBadRequest)
			return
		}
		if itemID == "" {
			log.Println("Item ID is required for DELETE request")
			http.Error(w, "Item ID is required", http.StatusBadRequest)
			return
		}
		// Delete the order item from the database
		result, err := cons.DB.Exec(cons.DeleteOrderItem, orderID, itemID)
		if err != nil {
			if utility.CheckSQLErr(err, w) {
				return
			}
			log.Println("Error deleting order item: ", err)
			http.Error(w, "Error deleting order item", http.StatusInternalServerError)
			return
		}
		if utility.CheckDeleteResult(result, w) {
			return
		}

		log.Println("Order item deleted successfully with ID: ", itemID)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
