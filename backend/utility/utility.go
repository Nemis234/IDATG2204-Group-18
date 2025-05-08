package producthandler

import (
	"database/sql"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-sql-driver/mysql"
)

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
		log.Println("MySQL error: ", err)
		http.Error(w, "MySQL error: "+mysqlErr.Message, http.StatusInternalServerError)
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
