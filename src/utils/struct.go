package utils

import (
	"fmt"
	"reflect"
)

// StructToMap converts a struct to a map[string]interface{} using dynamic tags (like neo4j).
// This function uses reflection to iterate over the fields of the struct and return a map with key-value pairs.
func StructToMap(v interface{}, tagName string) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("esperado uma struct, mas obteve %s", val.Kind())
	}
	typ := reflect.TypeOf(v)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		fieldTag := fieldType.Tag.Get(tagName)
		if fieldTag == "" {
			fieldTag = fieldType.Name
		}
		result[fieldTag] = field.Interface()
	}

	return result, nil
}