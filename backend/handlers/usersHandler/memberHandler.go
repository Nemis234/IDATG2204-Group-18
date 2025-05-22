package usershandler

import (
	cons "backend/constants"
	. "backend/structs"
	utility "backend/utility"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

/*
MemberHandler supports the following methods:
  - GET: Get member data for a user
  - POST: Add a new member
  - PUT: Update member data for a user
  - DELETE: Delete member data for a user

Only administrators and the user themselves can access any endpoint.

# GET

This method retrieves member data for a user. It requires the user ID to be passed in the URL path.
Example usage:

	Method: GET
	URL: /users/456/member
	Response:
	HTTP Status: 200 OK
	Body:
	{
		"UserID": "456",
		"membership_level": "Gold",
		"membership_start": "2023-01-01"
	}

# POST

This method adds a new member.
It requires the user ID to be passed in the URL path and the membership level and start date to be passed in the request body.

Example usage:

	Method: POST
	URL: /users/456/member
	Body:
	{
		"membership_level": "Gold",
		"membership_start": "2023-01-01"
	}
	Response:
	HTTP Status: 201 Created

# PUT

This method updates member data for a user.
It requires the user ID to be passed in the URL path and the membership level and start date to be passed in the request body.

Example usage:

	Method: PUT
	URL: /users/456/member
	Body:
	{
		"MembershipLevel": "Platinum",
		"MembershipStart": "2023-01-01"
	}
	Response:
	HTTP Status: 200 OK

# DELETE

This method deletes member data for a user.
It requires the user ID to be passed in the URL path.

Example usage:

	Method: DELETE
	URL: /users/456/member
	Response:
	HTTP Status: 204 No Content
*/
func MemberHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		userID := r.PathValue("user_id")
		if userID == "" {
			http.Error(w, "user ID is required", http.StatusBadRequest)
			return
		}

		// Check user permission
		if !utility.CheckPrivileges(r, w, &userID) {
			return
		}
		var member Member
		err := cons.DB.Get(&member, cons.QueryMember, userID)
		if err != nil {
			if utility.CheckSQLErr(err, w) {
				return
			}
			log.Println("Error getting member data:", err)
			http.Error(w, "Error getting member data", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(member); err != nil {
			http.Error(w, "Error encoding member data to JSON", http.StatusInternalServerError)
			return
		}

	case http.MethodPost:
		userID := r.PathValue("user_id")
		if userID == "" {
			http.Error(w, "user ID is required", http.StatusBadRequest)
			return
		}
		// Check user permission
		if !utility.CheckPrivileges(r, w, &userID) {
			return
		}
		var member Member
		if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
			http.Error(w, "Error decoding member data", http.StatusBadRequest)
			return
		}
		// Check if the user is already a member
		var existingMember Member
		err := cons.DB.Get(&existingMember, cons.QueryMember, userID)
		if err == nil {
			http.Error(w, "User is already a member", http.StatusConflict)
			return
		} else if err == sql.ErrNoRows {
			// User is not a member, proceed to insert
		} else if !utility.CheckSQLErr(err, w) {
			http.Error(w, "Error checking member status", http.StatusInternalServerError)
			return
		}
		member.UserID = userID
		// Insert new member data
		_, err = cons.DB.NamedExec(cons.InsertMember, member)
		if err != nil {
			http.Error(w, "Error inserting member data", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)

	case http.MethodPut:
		userID := r.PathValue("user_id")
		if userID == "" {
			http.Error(w, "user ID is required", http.StatusBadRequest)
			return
		}
		// Check user permission
		if !utility.CheckPrivileges(r, w, &userID) {
			return
		}
		var member Member
		if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
			http.Error(w, "Error decoding member data", http.StatusBadRequest)
			return
		}
		// Update member data
		_, err := cons.DB.NamedExec(cons.UpdateMember, member)
		if err != nil {
			if utility.CheckSQLErr(err, w) {
				return
			}
			log.Println("Error updating member data:", err)
			http.Error(w, "Error updating member data", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)

	case http.MethodDelete:
		userID := r.PathValue("user_id")
		if userID == "" {
			http.Error(w, "user ID is required", http.StatusBadRequest)
			return
		}
		// Check user permission
		if !utility.CheckPrivileges(r, w, &userID) {
			return
		}
		// Delete member data
		result, err := cons.DB.Exec(cons.DeleteMember, userID)
		if err != nil {
			if utility.CheckSQLErr(err, w) {
				return
			}
			http.Error(w, "Error deleting member data", http.StatusInternalServerError)
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
