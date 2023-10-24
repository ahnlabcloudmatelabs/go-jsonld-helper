package jsonld_helper

import (
	"reflect"
	"strings"
)

func getValue(value any) any {
	valueType := reflect.TypeOf(value)

	if isString(valueType) || isNumber(valueType) || isBool(valueType) {
		return value
	}

	if isMap(valueType) {
		return getValueInMap(value)
	}

	if isArray(valueType) {
		return getValueInArray(value)
	}

	return value
}

func getValueInMap(value any) any {
	valueMap := value.(map[string]any)

	if valueMap["@value"] != nil {
		return valueMap["@value"]
	}

	if valueMap["@id"] != nil {
		return valueMap["@id"]
	}

	return value
}

func getValueInArray(value any) any {
	if _, ok := value.([]any); ok {
		valueArray := value.([]any)

		if len(valueArray) == 1 {
			scope := valueArray[0]

			if isMap(reflect.TypeOf(scope)) {
				return getValueInMap(scope)
			}

			return scope
		}
	}

	if _, ok := value.([]map[string]any); ok {
		valueArray := value.([]map[string]any)

		if len(valueArray) == 1 {
			scope := valueArray[0]

			if isMap(reflect.TypeOf(scope)) {
				return getValueInMap(scope)
			}

			return scope
		}
	}

	return value
}

func isString(reflectType reflect.Type) bool {
	return strings.HasPrefix(reflectType.String(), "string")
}

func isNumber(reflectType reflect.Type) bool {
	for _, typeName := range []string{"int", "uint", "float"} {
		if strings.HasPrefix(reflectType.String(), typeName) {
			return true
		}
	}

	return false
}

func isBool(reflectType reflect.Type) bool {
	return strings.HasPrefix(reflectType.String(), "bool")
}

func isArray(reflectType reflect.Type) bool {
	return strings.HasPrefix(reflectType.String(), "[]")
}

func isMap(reflectType reflect.Type) bool {
	return strings.HasPrefix(reflectType.String(), "map")
}
