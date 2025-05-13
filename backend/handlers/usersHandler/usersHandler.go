package usershandler

import (
	cons "backend/constants"
	. "backend/structs"
	utility "backend/utility"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

/*
ProductsHandler support these methods:

  - GET retrieve user info with the given user id.
  - DELETE deletes a use from the database.

# GET

Retrieves the user info from the database.

Example usage:

	Method: GET
	Route: /users/0df01f83-a9a7-4afa-9b62-8d0bb9722849
	Response:
	HTTP code: 200 OK
	[
		{
			"user_id": "0df01f83-a9a7-4afa-9b62-8d0bb9722849",
			"username": "helloWorld",
			"password": "$2a$10$a8ZHeovSX0/NitUQBkHHHeTe8FqVRlJEmet29pUYDjjqkQA6KGaum",
			"email": "john_doe@gmail.com",
			"first_name": "John",
			"last_name": "Doe",
			"address": "Yolostreet 15",
			"role": "admin"
		},
	]

# DELETE

Deletes a user from the database, users can delete their own users, admins are allowed to delete other users.

General function flow:
->

Example usage:

	Method: DELETE
	Route: /users/0df01f83-a9a7-4afa-9b62-8d0bb9722849
	Response:
	HTTP code: ----
*/
func UserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("UserHandler called with method: ", r.Method)
	switch r.Method {
	case http.MethodGet:
		userID := r.PathValue("user_id")
		if userID == "" {
			http.Error(w, "User ID is required", http.StatusBadRequest)
			return
		}

		// Check if the user is logged in with a JWT token
		if !utility.CheckPrivileges(r, w, &userID) {
			return
		}

		var u User
		cons.DB.Get(&u, cons.QueryUser, userID)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(u)

	case http.MethodDelete:
		w.Header().Set("Content-Type", "application/json")

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

/*
ProductsHandler support these methods:

  - GET request for a list of users.
  - POST creates a new user.

# GET

The functions requires admin privileges, this is sent by the frontend through a JWT token in the Header (Authorization) upon request.

General function flow:
-> Validates the Token sent from the client
-> Checks privileges
-> Query the database
-> Write a respond to the client

Example usage:

	Method: GET
	Route: /users
	Response:
	HTTP code: 200 OK
	[
		{
			"user_id": "0df01f83-a9a7-4afa-9b62-8d0bb9722849",
			"username": "helloWorld",
			"password": "$2a$10$a8ZHeovSX0/NitUQBkHHHeTe8FqVRlJEmet29pUYDjjqkQA6KGaum",
			"email": "john_doe@gmail.com",
			"first_name": "John",
			"last_name": "Doe",
			"address": "Yolostreet 15",
			"role": "admin"
		},
		{
			"user_id": "8d29df7f-887d-4a02-aa78-89f25a669a3a",
			"username": "janedoe",
			"password": "$2a$10$qOd0Mt3q9oHZrBwEYGZa2.D8w6xaofMoUSFMG4ZE2VhHOK2j6Pf/.",
			"email": "jane_doe@gmail.com",
			"first_name": "Jane",
			"last_name": "Doe",
			"address": "Main Street 5",
			"role": "user"
		}
	]

# POST

Receives a payload with user informations and creates a new user.

General function flow:
-> Extract the user infor from the payload
-> Check if mandatory fields are missing
-> Query the database for existing username/email
-> Creates the struct to be stored in the database
-> Stores the data(new user) in the database
-> Write a respond to the client

Example usage:

	Method: POST
	Route: /users
	Request body:
	{

		"username":"will00",
		"password":"son00",
		"email":"willson00gmail.com",
		"first_name":"Will",
		"last_name":"Son",
		"address":"Big Street 28"
	}

	Response:
	HTTP code: 201 Created
	{
		"id" : "fiadg09b0-200b-4d08-939a-f90105f6546s"
	}
*/
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("UsersHandler called with method: ", r.Method)
	switch r.Method {
	case http.MethodGet:

		// Only admins can access this endpoint
		if !utility.CheckPrivileges(r, w, nil) {
			return
		}

		var users []User
		err := cons.DB.Select(&users, cons.QueryUsers)
		if err != nil {
			//Check if the error is a sql error
			if utility.CheckSQLErr(err, w) {
				return
			}
			log.Println("Error getting users: ", err)
			http.Error(w, "Error getting users", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(users); err != nil {
			log.Println("Error encoding users to JSON: ", err)
			http.Error(w, "Error encoding users to JSON", http.StatusInternalServerError)
			return
		}

	case http.MethodPost:

		// Decode the request body into a new User struct
		var newUser NewUser
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			log.Println("Error decoding request body: ", err)
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		//Check mandatory fields
		if newUser.Username == "" || newUser.Email == "" || newUser.Password == "" || newUser.FirstName == "" || newUser.LastName == "" {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}

		//Checks if the requested email already exists in the database
		var existingUsers int
		err = cons.DB.Get(&existingUsers, cons.QueryUserCountByEmailOrUsername, newUser.Email, newUser.Username)
		if err != nil {
			log.Println("Database error checking existing user:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		//Checks if the query returned more than 0 users
		if existingUsers > 0 {
			http.Error(w, "Email or username already in use", http.StatusConflict)
			return
		}

		//Creating neccesary fields to insert into struct to be stored.
		roughID, err := uuid.NewRandom()
		if err != nil {
			log.Println("Failed to generate UUID: ", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		newId := roughID.String()
		defaultRole := "user"

		//Encrypting the raw password
		hashedPassword, err := utility.HashPassword(newUser.Password)
		if err != nil {
			log.Println("Failed to hash password.")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		//Preparing struct to store in the database
		userToStore := User{
			UserID:    newId,
			Username:  newUser.Username,
			Password:  hashedPassword,
			Email:     newUser.Email,
			FirstName: newUser.FirstName,
			LastName:  newUser.LastName,
			Address:   newUser.Address,
			Role:      defaultRole,
		}

		_, err = cons.DB.NamedExec(cons.InsertUser, userToStore)
		if err != nil {
			// If the error is a MySQL error, return
			if utility.CheckSQLErr(err, w) {
				return
			}

			log.Println("Error inserting user: ", err)
			http.Error(w, "Error creating user", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		// Return the ID of the newly created user
		response := map[string]string{"id": newId}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Println("Error encoding response to JSON: ", err)
			http.Error(w, "Error encoding response to JSON", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
