package utils

import (
	"fmt"
	"reflect"
)

// StructToMap converts a struct to a map[string]interface{}.
// This function uses reflection to iterate over the fields of the struct and return a map with key-value pairs.
func StructToMap(v interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	// Get the value of the struct
	val := reflect.ValueOf(v)

	// Check if the provided value is a struct
	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a struct but got %s", val.Kind())
	}
	// Get the type of the struct
	typ := reflect.TypeOf(v)

	// Iterate over the struct fields
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name
		result[fieldName] = field.Interface()
	}

	return result, nil
}
