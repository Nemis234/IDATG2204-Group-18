package usershandler

import (
	cons "backend/constants"
	. "backend/structs"
	"encoding/json"
	"log"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("UserHandler called with method: ", r.Method)
	switch r.Method {
	case http.MethodGet:
		userID := r.PathValue("user_id")
		if userID == "" {
			http.Error(w, "User ID is required", http.StatusBadRequest)
			return
		}

		var u User
		cons.DB.Get(&u, cons.QueryUser, userID)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(u)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("UsersHandler called with method: ", r.Method)
	switch r.Method {
	case http.MethodGet:
		http.Error(w, "Method not implemented", http.StatusNotImplemented)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
