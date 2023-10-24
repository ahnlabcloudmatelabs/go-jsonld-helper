package jsonld_helper

import (
	"fmt"
	"reflect"
	"strings"
)

func readIndex(jsonld any, index int) JsonLDReader {
	if !isArray(reflect.TypeOf(jsonld)) {
		return Nothing{
			Error: fmt.Errorf("not an array"),
		}
	}

	return readAsArray(jsonld.([]any), index)
}

func readKey(jsonld any, key string) JsonLDReader {
	if isArray(reflect.TypeOf(jsonld)) {
		if scope, ok := jsonld.([]any); ok {
			if len(scope) == 1 {
				return readKey(scope[0], key)
			}
		}

		if scope, ok := jsonld.([]map[string]any); ok {
			if len(scope) == 1 {
				return readKey(scope[0], key)
			}
		}

		return Nothing{
			Error: fmt.Errorf("not an map"),
		}
	}

	if isMap(reflect.TypeOf(jsonld)) {
		return readAsMap(jsonld.(map[string]any), key)
	}

	return Nothing{
		Error: fmt.Errorf("not an map"),
	}
}

func readAsArray(scope []any, index int) JsonLDReader {
	if len(scope) <= index {
		return Nothing{
			Error: fmt.Errorf("not found index: %d", index),
		}
	}

	value := scope[index]

	if value != nil {
		return of(value)
	}

	return Nothing{
		Error: fmt.Errorf("not found index: %d", index),
	}
}

func readAsMap(scope map[string]any, key string) JsonLDReader {
	if scope["@value"] != nil && isArray(reflect.TypeOf(scope["@value"])) {
		return readKey(scope["@value"], key)
	}

	value, isPreDefined := readAsPreDefinedKey(scope, key)

	if isPreDefined {
		return of(value)
	}

	fullKey := getFullKey(scope, key)

	if fullKey == nil {
		return Nothing{
			Error: fmt.Errorf("not found key: %s", key),
		}
	}

	if key == "type" {
		return of(extractType(scope[*fullKey]))
	}

	return of(scope[*fullKey])
}

func getFullKey(scope map[string]any, key string) *string {
	for k := range scope {
		if k == key {
			return &k
		}

		sharpSplit := strings.Split(k, "#")

		if len(sharpSplit) == 2 && sharpSplit[1] == key {
			return &k
		}

		slashSplit := strings.Split(k, "/")

		if len(slashSplit) > 1 && slashSplit[len(slashSplit)-1] == key {
			return &k
		}
	}

	return nil
}

func readAsPreDefinedKey(scope map[string]any, key string) (any, bool) {
	if key == "type" || key == "@type" {
		_type, founded := readType(scope)

		if founded {
			return _type, true
		}
	}

	if key == "id" || key == "@id" {
		id, founded := readID(scope)

		if founded {
			return id, true
		}
	}

	if key == "value" || key == "@value" {
		value, founded := readValue(scope)

		if founded {
			return value, true
		}
	}

	return nil, false
}

func readType(scope map[string]any) (any, bool) {
	value := scope["@type"]

	if value != nil {
		return extractType(value), true
	}

	for key := range scope {
		sharpSplit := strings.Split(key, "#")

		if len(sharpSplit) == 2 && sharpSplit[1] == "type" {
			return scope[key], true
		}
	}

	return nil, false
}

func readID(scope map[string]any) (any, bool) {
	value := scope["@id"]

	if value != nil {
		return value, true
	}

	for key := range scope {
		sharpSplit := strings.Split(key, "#")

		if len(sharpSplit) == 2 && sharpSplit[1] == "id" {
			return scope[key], true
		}
	}

	return nil, false
}

func readValue(scope map[string]any) (any, bool) {
	value := scope["@value"]

	if value != nil {
		return value, true
	}

	for key := range scope {
		sharpSplit := strings.Split(key, "#")

		if len(sharpSplit) == 2 && sharpSplit[1] == "value" {
			return scope[key], true
		}
	}

	return nil, false
}

func extractType(value any) string {
	valueType := reflect.TypeOf(value)

	if isArray(valueType) {
		return extractType(value.([]any)[0])
	}

	if isMap(valueType) {
		return value.(map[string]any)["@value"].(string)
	}

	hashSplit := strings.Split(value.(string), "#")

	if len(hashSplit) == 2 {
		return hashSplit[1]
	}

	slashSplit := strings.Split(value.(string), "/")

	return slashSplit[len(slashSplit)-1]
}
