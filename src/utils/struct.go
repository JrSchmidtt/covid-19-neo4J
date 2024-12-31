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

// MapToStruct converts a map[string]interface{} to a struct using dynamic tags (like neo4j).
// This function uses reflection to iterate over the map and populate the struct fields.
func MapToStruct(data map[string]interface{}, result interface{}) error {
	value := reflect.ValueOf(result).Elem()
	typ := reflect.TypeOf(result).Elem()
	for key, val := range data {
		field, ok := typ.FieldByNameFunc(func(fieldName string) bool {
			field, found := typ.FieldByName(fieldName)
			if !found {
				return false
			}
			jsonTag := field.Tag.Get("json")
			return jsonTag == key
		})
		if !ok {
			continue
		}
		fieldValue := value.FieldByName(field.Name)
		if !fieldValue.IsValid() || !fieldValue.CanSet() {
			continue
		}
		fieldValueReflect := reflect.ValueOf(val)
		if fieldValueReflect.Type().ConvertibleTo(fieldValue.Type()) {
			fieldValue.Set(fieldValueReflect.Convert(fieldValue.Type()))
		} else {
			return fmt.Errorf("type mismatch for field %s", key)
		}
	}
	return nil
}