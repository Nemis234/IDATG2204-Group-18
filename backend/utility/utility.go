package producthandler

import (
	cons "backend/constants"
	structs "backend/structs"
	"database/sql"
	"log"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"
	"os"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/golang-jwt/jwt/v5"
)

/*
Thanks to https://www.reddit.com/r/golang/comments/a85ex4/comment/ec89v3b/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button
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
		http.Error(w, "Order item not found", http.StatusNotFound)
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
		http.Error(w, "Item not found.", http.StatusNotFound)
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
func GenerateJWT(userID, role string, jwtKey []byte) (string, error) {

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
        return nil, fmt.Errorf("No authorization header")
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
        return nil, fmt.Errorf("Invalid token")
    }

    return jwtToken, nil
}