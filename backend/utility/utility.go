package producthandler

import (
	"reflect"
	"strings"
)

// BuildUpdateQuery generates the SET clause and corresponding args for a PATCH query
func BuildUpdateQuery(input interface{}) (string, []interface{}) {
	v := reflect.ValueOf(input)
	if v.Kind() != reflect.Struct {
		panic("input must be a struct")
	}

	var setParts []string
	var args []interface{}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		// Get json tag as column name (or fallback to field name)
		tag := fieldType.Tag.Get("json")
		if tag == "" || tag == "-" {
			continue
		}
		tag = strings.Split(tag, ",")[0]

		if field.Kind() == reflect.Ptr && !field.IsNil() {
			setParts = append(setParts, tag+" = ?")
			args = append(args, field.Elem().Interface())
		}
	}

	return strings.Join(setParts, ", "), args
}
