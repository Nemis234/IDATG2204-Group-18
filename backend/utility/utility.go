package producthandler

import (
	cons "backend/constants"
	structs "backend/structs"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

/*
Thanks to https://www.reddit.com/r/golang/comments/a85ex4/comment/ec89v3b/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button
How to use:

	transactError := utility.Transact(func(tx *sqlx.Tx) error {
		// Your database operations here
		// Use tx instead of cons.DB for all operations
		// Example:
		_, err := tx.Exec("INSERT INTO users (name) VALUES (?)", "John Doe")
		if err != nil {
			return err
		}
		return nil
	}
*/
func Transact(fn func(*sqlx.Tx) error) error {
	tx, err := cons.DB.Beginx()
	if err != nil {
		return err
	}

	err = fn(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func CheckDeleteResult(res sql.Result, w http.ResponseWriter) bool {
	affected, err := res.RowsAffected()
	if err != nil {
		log.Println("Error getting affected rows: ", err)
		http.Error(w, "Error getting affected rows", http.StatusInternalServerError)
		return true
	}
	if affected == 0 {
		log.Println("No rows affected by delete")
		http.Error(w, "Item not found", http.StatusNotFound)
		return true
	}
	return false
}

/*
Returns true if the error is a MySQL error
*/
func CheckSQLErr(err error, w http.ResponseWriter) bool {
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		if mysqlErr.Number == 1451 {
			log.Println("Foreign key constraint error: ", err)
			http.Error(w, "Cannot update field with existing references, probably due to a foreign key constraint. Error: "+mysqlErr.Message, http.StatusConflict)
			return true
		}
		if mysqlErr.Number == 1048 {
			log.Println("Column cannot be null error: ", err)
			http.Error(w, "A mandatory column cannot be null. Error: "+mysqlErr.Message, http.StatusBadRequest)
			return true
		}
		if mysqlErr.Number == 1062 {
			log.Println("Duplicate entry error: ", err)
			http.Error(w, "Duplicate entry. Another value of this type already exists. Error: "+mysqlErr.Message, http.StatusConflict)
			return true
		}
		log.Println("MySQL error: ", err)
		http.Error(w, "MySQL error: "+mysqlErr.Message, http.StatusInternalServerError)
		return true
	}
	if err == sql.ErrNoRows {
		log.Println("No rows found error: ", err)
		http.Error(w, "Item not found", http.StatusNotFound)
		return true
	}
	return false
}

/*
Made by ChatGPT

# BuildUpdateQuery generates the SET clause and corresponding args for a PATCH query

It takes a struct as input with the custom NullField[] struct in its fields and returns a string and a slice of the struct's field values.
The string is a comma-separated list of column assignments for an SQL UPDATE statement.
The slice contains the values of the fields that are not nil.
*/
func BuildUpdateQuery(input any) (string, []any) {
	v := reflect.ValueOf(input)
	if v.Kind() != reflect.Struct {
		log.Println("Input must be a struct")
		return "", nil
	}

	var setParts []string
	var args []any
	t := v.Type()

	for i := range v.NumField() {
		field := v.Field(i)
		fieldType := t.Field(i)

		// Only look at fields of type NullField[T]
		if field.Kind() != reflect.Struct || !field.FieldByName("Set").Bool() {
			continue
		}

		// Get database tag as column name (or fallback to field name)
		tag := fieldType.Tag.Get("db")
		if tag == "" || tag == "-" {
			continue
		}
		tag = strings.Split(tag, ",")[0]

		valField := field.FieldByName("Value")
		if valField.IsNil() {
			setParts = append(setParts, tag+" = NULL")
		} else {
			setParts = append(setParts, tag+" = ?")
			args = append(args, valField.Elem().Interface())
		}
	}

	return strings.Join(setParts, ", "), args
}

/*
Used to get the ENV "JWT_TOKEN_KEY", if it does not exit it defautls to the test file in the rootfolder :"jwt-key-for-testing.txt"

This function is mainly used in the two functions below: GenerateJWT and ValidateJWT
*/
func GetJwtKey() ([]byte, error) {
	keyStr := os.Getenv("JWT_TOKEN_KEY")
	if keyStr != "" {
		return []byte(keyStr), nil
	}

	keyBytes, err := os.ReadFile("jwt-key-for-testing.txt")
	if err != nil {
		log.Println("Failed to read jwt-key-for-testing.txt:", err)
		return nil, err
	}
	return keyBytes, nil
}

/*
Made with help from ChatGPT

It generates a JWT-token, this stores the userID togethers with the corresponding role in a struct.
This token will be used by the frontend/backend to check privileges.
*/
func GenerateJWT(userID string, role *string, jwtKey []byte) (string, error) {

	expirationTime := time.Now().Add(24 * time.Hour)

	jwtToken := &structs.JWTToken{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtToken)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		log.Println("Error signing the token:", err)
		return "", err
	}

	return signedToken, nil
}

/*
Made with help from ChatGPT

This functions takes in request and validates the JWT in the http.Request.
First it checks if any JWT token is recieved from the client,
then proceeds to validate it, then extracts the data before returning.
*/

func ValidateJWT(r *http.Request, jwtKey []byte) (*structs.JWTToken, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("no authorization header")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, fmt.Errorf("invalid authorization header format")
	}

	tokenStr := parts[1]

	jwtToken := &structs.JWTToken{}
	token, err := jwt.ParseWithClaims(tokenStr, jwtToken, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return jwtToken, nil
}

/*
CheckPrivileges checks if the user has the right privileges to access the resource.
It checks if the user is logged in with a JWT token.

Admins will always have access to the resource, regardless of the userID.

To check only for admin privileges, pass nil as the userID parameter.

To check if the right user is logged in with a JWT token, pass the userID parameter as a pointer &string.

Example:

	if !CheckPrivileges(r, w, nil) {
		return
	}
	if !CheckPrivileges(r, w, &userID) {
		return
	}

This function will return true if the user has access, false otherwise.
*/
func CheckPrivileges(r *http.Request, w http.ResponseWriter, userID *string) bool {
	//Check if user is logged in with a JWT token
	jwtTokenData, err := ValidateJWT(r, cons.JwtKey)
	if err != nil {
		log.Println("Error validating JWT token: ", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return false
	}

	//Check if the request has admin privileges
	if jwtTokenData.Role != nil {
		return true
	}

	if userID != nil {
		userIDValue := *userID
		// Check if the user ID in the token matches the user ID in the request
		if jwtTokenData.UserID == userIDValue {
			return true
		}
		log.Println("User ID does not match, expected: ", userID, " got: ", jwtTokenData.UserID)
	} else {
		log.Println("User does not have admin privileges")
	}

	http.Error(w, "Forbidden", http.StatusForbidden)
	return false
}

/*
Helper function to check if the user has privileges to access a table.
It returns true if the user has access, false otherwise.

To use, prepare a query that gets the userID from the database, and pass it to this function.
Then pass it the required arguments to the query.

Example:

	query := "SELECT userID FROM orders WHERE orderID = ?"
	args := []any{orderID}
	if !CheckUser(r, w, query, args...) {
		return
	}
	// or
	if !CheckUser(r, w, query, orderID) {
		return
	}
*/
func CheckUser(r *http.Request, w http.ResponseWriter, query string, args ...any) bool {
	// Get user ID from database
	var userID string
	err := cons.DB.Get(&userID, query, args...)
	if err != nil {
		if CheckSQLErr(err, w) {
			return false
		}
		log.Println("Error fetching user ID: ", err)
		http.Error(w, "Error fetching user ID", http.StatusInternalServerError)
		return false
	}
	// Check user privileges
	if !CheckPrivileges(r, w, &userID) {
		return false
	}
	return true
}

func GetUserID(r *http.Request, w http.ResponseWriter) *string {
	jwtTokenData, err := ValidateJWT(r, cons.JwtKey)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return nil
	}
	return &jwtTokenData.UserID
}

/*
*   Hash the password before storing it. Using the bycrypt which is in GO's libary.
*   It includes salt.
*
*   password - This is the password string that is to be hashed
 */
func HashPassword(password string) (string, error) {
	// Generate a salt + hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
