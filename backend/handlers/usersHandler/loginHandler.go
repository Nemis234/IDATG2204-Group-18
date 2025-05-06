package usershandler

import (
	cons "backend/constants"
	. "backend/structs"
	"encoding/json"
	"log"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("LoginHandler called with method: ", r.Method)
	switch r.Method {
	case http.MethodPost:
		var newUser User
		if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		if newUser.Email == "" || newUser.Password == "" {
			http.Error(w, "Username and password are required", http.StatusBadRequest)
			return
		}
		var checkUser User

		cons.DB.Get(&checkUser, cons.QueryUserLogin, newUser.Email)

		w.Header().Set("Content-Type", "application/json")

		log.Println("User: ", checkUser.Username, " Password: ", checkUser.Password)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
