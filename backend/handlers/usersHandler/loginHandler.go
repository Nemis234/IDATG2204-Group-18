package usershandler

import (
	cons "backend/constants"
	. "backend/structs"
	utility "backend/utility"
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

/*
LoginHandler support these methods:

  - POST Request for comparing login credentials and returns a JWT token used for authentications.

# POST

It retrieves the user information from the database and compares the password.
Then generates a JWT and returns this to the client as a Header. This must be stored in the frontend
and will be used for later on protected requests that requires admin privileges.
It can be made to return a JSON body containing the JWT token if needed, but it does not support that for now.

General function flow :
-> Extract the Payload
-> Check if the payload is not empty
-> Query the database for the user login information
-> Compare the password
-> Generates a JWT token, this is sent back to the client. Used to distinguish user roles on protected requests
-> Writes a respond to the client

Possible Responds:
200: OK, user is authenticated
400: Bad Request, Payload is not complete/missing fields required
401: Unauthorized, wrong password or username/email
404: Not found, query returned 0 rows, no user found in the database with the given username/email
500: Internal Server Error

Example usage:

	Method: Post
	Route: /users/login
	Request body:
	{
		"Email": "john_doe@gmail.com",
		"Password": "helloWorld"
	}

	Respond body:
	HTTP code: 200 No Content
	Respond header:
	Authorization: Bearer "jwt-token"
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
		var err error
		if newUser.Email != "" {
			err = cons.DB.Get(&checkUser, cons.QueryUserLoginByEmail, newUser.Email)
		} else {
			err = cons.DB.Get(&checkUser, cons.QueryUserLoginByUsername, newUser.Username)
		}

		//Check if user was found in the database or for any other sql errors
		if utility.CheckSQLErr(err, w) {
			return
		}

		//Using bcrypt to compared the raw password in the payload with the stored hashed password
		err = bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(newUser.Password))
		if err != nil {
			//Password does not match
			log.Println("Invalid credentials")
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		//Creating a JWT token to attach to the user, used to check user privileges.
		token, err := utility.GenerateJWT(checkUser.UserID, checkUser.RoleName, cons.JwtKey)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Authorization", "Bearer "+token)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		id := checkUser.UserID
		response := map[string]string{"user_id": id, "auth_token": "Bearer " + token}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Println("Error encoding response: ", err)
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}

		log.Println("JWT-token: ", token)
		log.Println("User: ", checkUser.Username, " Password: ", checkUser.Password)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
