package usershandler

import (
	cons "backend/constants"
	. "backend/structs"
	"encoding/json"
	"log"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

/*
LoginHandler support these methods:

  - POST Request for comparing login credentials.

# GET

It retrieves the user information from the database and compares the password.
General function flow :
-> Extract the Payload
-> Check if the payload is not empty
-> Query the database for the user login information
-> Compare the password
-> Writes a respond to the client

Possible Responds:
200: OK, user is authenticated
401: Wrong credentials, wrong password or username/email
400: Missing fields, Payload is not complete/missing fields required
500: Internal Server Error

Example usage:

	Method: Post
	Route: /users/login
	Request body:
	{
		"Username": "helloWorld",
		"Email": "john_doe@gmail.com",
		"Password": "helloWorld"
	}
	Respond body:
	HTTP code: 200 No Content

*/
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("LoginHandler called with method: ", r.Method)
	switch r.Method {
	case http.MethodPost:
		var newUser User
		if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		if (newUser.Email == "" && newUser.Username == "") || newUser.Password == "" {
			http.Error(w, "Username/Email and password are required", http.StatusBadRequest)
			return
		}
		var checkUser User

		//Use username or email to query the database, depending on which was given
		if (newUser.Email != ""){
			cons.DB.Get(&checkUser, cons.QueryUserLogin, newUser.Email)
		}else{
			cons.DB.Get(&checkUser, cons.QueryUserLoginUsername, newUser.Username)
		}

		//Using bcrypt to compared the raw password in the payload with the stored hashed password
		err := bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(newUser.Password))
		if err != nil {
			//Password does not match
			log.Println("Invalid credentials")
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		log.Println("User: ", checkUser.Username, " Password: ", checkUser.Password)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}