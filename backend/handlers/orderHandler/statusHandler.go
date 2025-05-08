package orderHandler

import (
	cons "backend/constants"
	. "backend/structs"
	utility "backend/utility"
	"encoding/json"
	"log"
	"net/http"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("OrderStatusHandler called with method: ", r.Method)
	switch r.Method {
	case http.MethodGet:
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
