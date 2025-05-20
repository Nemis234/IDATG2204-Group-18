package orderHandler

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
PaymentsHandler supports the following methods:
  - GET: Get all payments for an order. The order ID is required in the URL path.
  - POST: Create a new payment for an order. The order ID is required in the URL path.

All methods require the user to have privileges to access the order.

# GET

Get all payments for an order. The order ID is required in the URL path.

Example usage:

	Method: GET
	Path: /orders/1/payments
	Response:
	[
		{
			"payment_id": "1",
			"order_id": "1",
			"payment_date": "2023-10-01",
			"payment_amount": 100.00,
			"payment_method": "credit_card",
			"payment_status": "completed"
		}
	]

# POST

Create a new payment for an order. The order ID is required in the URL path.

The request body should contain the following fields:

	{
		  "payment_date": "string",
		  "payment_amount": 0,
		  "payment_method": "card"/"vipps"/"bank_transfer",
		  "payment_status": "pending"/"successful"/"failed",
	}

Example usage:

	Method: POST
	Path: /orders/1/payments
	{
		"payment_date": "2023-10-01",
		"payment_amount": 100.00,
		"payment_method": "card",
		"payment_status": "pending"
	}
	Response:
	HTTP 201 Created
	{
		"id": "1"
	}
*/
func PaymentsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("PaymentHandler called with method: ", r.Method)
	switch r.Method {
	case http.MethodGet:
		orderID := r.PathValue("order_id")
		if orderID == "" {
			http.Error(w, "Order ID is required", http.StatusBadRequest)
			return
		}

		// Check user permission with a sql query
		if !utility.CheckUser(r, w, cons.GetUserIDByOrderID, orderID) {
			return
		}

		var payments []Payment
		err := cons.DB.Select(&payments, cons.QueryPayment, orderID)
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
		if err := json.NewEncoder(w).Encode(payments); err != nil {
			log.Println("Error encoding cart items to JSON: ", err)
			http.Error(w, "Error encoding cart items to JSON", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		orderID := r.PathValue("order_id")
		if orderID == "" {
			http.Error(w, "Order ID is required", http.StatusBadRequest)
			return
		}
		var payment Payment
		if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
			log.Println("Error decoding json: ", err)
			http.Error(w, "Error decoding json", http.StatusBadRequest)
			return
		}

		// Check user privileges
		if !utility.CheckUser(r, w, cons.GetUserIDByOrderID, payment.OrderID) {
			return
		}

		payment.OrderID = orderID

		// Generate a new payment ID
		err := cons.DB.Get(&payment, "SELECT UUID() AS "+cons.PAYMENT_ID+";")
		if err != nil {
			if utility.CheckSQLErr(err, w) {
				return
			}

			log.Println("Error generating payment ID: ", err)
			http.Error(w, "Error generating payment ID", http.StatusInternalServerError)
			return
		}
		log.Println("Payment ID: ", payment.PaymentID)
		paymentID := payment.PaymentID

		_, err = cons.DB.NamedExec(cons.InsertPayment, payment)
		if err != nil {
			if utility.CheckSQLErr(err, w) {
				return
			}
			log.Println("Error inserting payment: ", err)
			http.Error(w, "Error inserting payment", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		// Return the ID of the newly created payment
		response := map[string]string{"id": paymentID}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Println("Error encoding payment to JSON: ", err)
			http.Error(w, "Error encoding payment to JSON", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

/*
PaymentHandler supports the following methods:
  - PUT: Update an existing payment. The payment ID and order ID are required in the URL path.
  - PATCH: Partially update an existing payment. The payment ID and order ID are required in the URL path.
  - DELETE: Delete an existing payment. The payment ID and order ID are required in the URL path.

All methods require the user to have privileges to access the payment.

# PUT

Update an existing payment. The payment ID and order ID are required in the URL path.
The request body should contain the following fields:

	{
		  "payment_date": "string",
		  "payment_amount": 0,
		  "payment_method": "card"/"vipps"/"bank_transfer",
		  "payment_status": "pending"/"successful"/"failed",
	}

Example usage:

	Method: PUT
	Path: /orders/1/payments/1
	{
		"payment_date": "2023-10-01",
		"payment_amount": 100.00,
		"payment_method": "card",
		"payment_status": "successful"
	}
	Response:
	HTTP 200 OK

# PATCH

Partially update an existing payment. The payment ID and order ID are required in the URL path.
The request body can contain any of the fields used in the PUT method.

# DELETE

Delete an existing payment. The payment ID and order ID are required in the URL path.

Example usage:

	Method: DELETE
	Path: /orders/1/payments/1
	Response:
	HTTP 204 No Content
*/
func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		orderID := r.PathValue("order_id")
		if orderID == "" {
			http.Error(w, "Order ID is required", http.StatusBadRequest)
			return
		}
		paymentID := r.PathValue("payment_id")
		if paymentID == "" {
			http.Error(w, "Payment ID is required", http.StatusBadRequest)
			return
		}

		// Check user privileges
		if !utility.CheckUser(r, w, cons.GetUserIDByOrderID, orderID) {
			return
		}

		var payment Payment
		if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
			log.Println("Error decoding payment: ", err)
			http.Error(w, "Error decoding payment", http.StatusBadRequest)
			return
		}
		if payment.OrderID != "" && payment.OrderID != orderID {
			http.Error(w, "Order ID cannot be changed", http.StatusBadRequest)
			return
		}
		if payment.PaymentID != "" && payment.PaymentID != paymentID {
			http.Error(w, "Payment ID cannot be changed", http.StatusBadRequest)
			return
		}
		payment.OrderID = orderID
		payment.PaymentID = paymentID

		// Update the payment in the database
		_, err := cons.DB.NamedExec(cons.UpdatePayment, payment)
		if err != nil {
			if utility.CheckSQLErr(err, w) {
				return
			}
			log.Println("Error updating payment: ", err)
			http.Error(w, "Error updating payment", http.StatusInternalServerError)
			return
		}
		log.Println("Payment updated successfully with ID: ", payment.PaymentID)
		w.WriteHeader(http.StatusOK)
	case http.MethodPatch:
		orderID := r.PathValue("order_id")
		if orderID == "" {
			http.Error(w, "Order ID is required", http.StatusBadRequest)
			return
		}
		paymentID := r.PathValue("payment_id")
		if paymentID == "" {
			http.Error(w, "Payment ID is required", http.StatusBadRequest)
			return
		}
		// Check user privileges
		if !utility.CheckUser(r, w, cons.GetUserIDByOrderID, orderID) {
			return
		}

		var paymentPatch PaymentPatch
		if err := json.NewDecoder(r.Body).Decode(&paymentPatch); err != nil {
			log.Println("Error decoding payment: ", err)
			http.Error(w, "Error decoding payment", http.StatusBadRequest)
			return
		}

		if paymentPatch.OrderID != "" && paymentPatch.OrderID != orderID {
			http.Error(w, "Order ID cannot be changed", http.StatusBadRequest)
			return
		}
		if paymentPatch.PaymentID != "" && paymentPatch.PaymentID != paymentID {
			http.Error(w, "Payment ID cannot be changed", http.StatusBadRequest)
			return
		}
		paymentPatch.OrderID = orderID
		paymentPatch.PaymentID = paymentID

		query := "UPDATE " + cons.PAYMENT_TABLE + " SET "

		setClause, args := utility.BuildUpdateQuery(paymentPatch)
		if setClause == "" {
			http.Error(w, "No fields to update", http.StatusBadRequest)
			return
		}
		query += setClause + " WHERE " + cons.PAYMENT_ID + " = ?"
		args = append(args, paymentID)
		log.Println("Executing query: ", fmt.Sprintf(strings.ReplaceAll(query, "?", "%s"), args...))

		_, err := cons.DB.Exec(query, args...)
		if err != nil {
			if utility.CheckSQLErr(err, w) {
				return
			}
			log.Println("Error updating payment: ", err)
			http.Error(w, "Error updating payment", http.StatusInternalServerError)
			return
		}
		log.Println("Payment updated successfully with ID: ", paymentID)
		w.WriteHeader(http.StatusOK)
	case http.MethodDelete:
		orderID := r.PathValue("order_id")
		if orderID == "" {
			http.Error(w, "Order ID is required", http.StatusBadRequest)
			return
		}
		paymentID := r.PathValue("payment_id")
		if paymentID == "" {
			http.Error(w, "Payment ID is required", http.StatusBadRequest)
			return
		}

		// Check user privileges
		if !utility.CheckUser(r, w, cons.GetUserIDByOrderID, orderID) {
			return
		}

		// Delete the payment from the database
		result, err := cons.DB.Exec(cons.DeletePayment, paymentID)
		if err != nil {
			if utility.CheckSQLErr(err, w) {
				return
			}
			log.Println("Error deleting payment: ", err)
			http.Error(w, "Error deleting payment", http.StatusInternalServerError)
			return
		}
		if utility.CheckDeleteResult(result, w) {
			return
		}
		log.Println("Payment deleted successfully with ID: ", paymentID)
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
