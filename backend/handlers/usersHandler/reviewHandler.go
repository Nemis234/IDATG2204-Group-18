package usershandler

import (
	cons "backend/constants"
	. "backend/structs"
	utility "backend/utility"
	"encoding/json"
	"log"
	"net/http"
)

func UserReviewsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		id := r.PathValue("user_id")
		if id == "" {
			http.Error(w, "User ID is required", http.StatusBadRequest)
			return
		}

		var reviews []Review
		err := cons.DB.Select(&reviews, cons.QueryReviewsByUser, id)
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
		http.Error(w, "Method not implemented", http.StatusNotImplemented)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func UserReviewHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		http.Error(w, "Method not implemented", http.StatusNotImplemented)
	case http.MethodPut:
		http.Error(w, "Method not implemented", http.StatusNotImplemented)
	case http.MethodDelete:
		http.Error(w, "Method not implemented", http.StatusNotImplemented)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
