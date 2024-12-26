package utils

import (
	"fmt"
	"reflect"
)

// StructToMap converts a struct to a map[string]interface{}.
// This function uses reflection to iterate over the fields of the struct and return a map with key-value pairs.
// StructToMap converte uma struct para um map[string]interface{} usando tags dinâmicas (como neo4j).
func StructToMap(v interface{}, tagName string) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	// Obter o valor da struct
	val := reflect.ValueOf(v)

	// Verificar se o tipo fornecido é uma struct
	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("esperado uma struct, mas obteve %s", val.Kind())
	}

	// Obter o tipo da struct
	typ := reflect.TypeOf(v)

	// Iterar sobre os campos da struct
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		// Obter a tag desejada, por padrão é "json", mas pode ser "neo4j"
		fieldTag := fieldType.Tag.Get(tagName)
		if fieldTag == "" {
			// Se não houver a tag, usa o nome do campo
			fieldTag = fieldType.Name
		}

		// Adicionar o campo no mapa com o nome da chave da tag especificada
		result[fieldTag] = field.Interface()
	}

	return result, nil
}