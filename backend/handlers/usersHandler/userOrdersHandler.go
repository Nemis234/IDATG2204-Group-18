package usershandler

import (
	cons "backend/constants"
	. "backend/structs"
	utility "backend/utility"
	"encoding/json"
	"log"
	"net/http"
)

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
