package jsonld_helper

import (
	"fmt"
	"reflect"
	"strconv"
)

type Reader struct {
	value  any
	length int
}

func of(value any) JsonLDReader {
	jsonldReader := Reader{
		value:  value,
		length: 1,
	}

	if isArray(reflect.TypeOf(value)) {
		if _, ok := value.([]any); ok {
			jsonldReader.length = len(value.([]any))
		}

		if _, ok := value.([]map[string]any); ok {
			jsonldReader.length = len(value.([]map[string]any))
		}
	}

	return jsonldReader
}

func (r Reader) Value() any                       { return r.value }
func (r Reader) Length() int                      { return r.length }
func (r Reader) ReadKey(key string) JsonLDReader  { return readKey(r.value, key) }
func (r Reader) ReadIndex(index int) JsonLDReader { return readIndex(r.value, index) }

func (r Reader) Get() any                          { return getValue(r.value) }
func (r Reader) GetOrElse(defaultValue any) any    { return getValue(r.value) }
func (r Reader) GetOrThrow(err error) (any, error) { return getValue(r.value), nil }

func (r Reader) StringOrElse(defaultValue string) string {
	value, err := r.StringOrThrow(nil)
	if err != nil {
		return defaultValue
	}

	return value
}

func (r Reader) StringOrThrow(err error) (string, error) {
	value := getValue(r.value)

	stringValue, ok := value.(string)
	if ok {
		return stringValue, nil
	}

	boolValue, boolOK := value.(bool)
	if boolOK {
		return fmt.Sprintf("%t", boolValue), nil
	}

	intValue, intOK := value.(int)
	if intOK {
		return fmt.Sprintf("%d", intValue), nil
	}

	floatValue, floatOK := value.(float64)
	if floatOK {
		return fmt.Sprintf("%f", floatValue), nil
	}

	if err != nil {
		return "", err
	}

	return "", fmt.Errorf("not a string")
}

func (r Reader) BoolOrElse(defaultValue bool) bool {
	value, err := r.BoolOrThrow(nil)
	if err != nil {
		return defaultValue
	}

	return value
}
func (r Reader) BoolOrThrow(err error) (bool, error) {
	value := getValue(r.value)

	if boolValue, ok := value.(bool); ok {
		return boolValue, nil
	}

	if stringValue, ok := value.(string); ok {
		if stringValue == "true" || stringValue == "false" || stringValue == "1" || stringValue == "0" {
			return strconv.ParseBool(stringValue)
		}
	}

	if intValue, ok := value.(int); ok {
		if intValue == 1 || intValue == 0 {
			return intValue == 1, nil
		}
	}

	if err != nil {
		return false, err
	}

	return false, fmt.Errorf("not a bool")
}

func (r Reader) IntOrElse(defaultValue int) int {
	value, err := r.IntOrThrow(nil)
	if err != nil {
		return defaultValue
	}

	return value
}

func (r Reader) IntOrThrow(err error) (int, error) {
	value := getValue(r.value)

	if intValue, ok := value.(int); ok {
		return intValue, nil
	}

	if stringValue, ok := value.(string); ok {
		if intValue, err := strconv.Atoi(stringValue); err == nil {
			return intValue, nil
		}
	}

	if err != nil {
		return 0, err
	}

	return 0, fmt.Errorf("not a int")
}

func (r Reader) FloatOrElse(defaultValue float64) float64 {
	value, err := r.FloatOrThrow(nil)
	if err != nil {
		return defaultValue
	}

	return value
}

func (r Reader) FloatOrThrow(err error) (float64, error) {
	value := getValue(r.value)

	if intValue, ok := value.(float64); ok {
		return intValue, nil
	}

	if stringValue, ok := value.(string); ok {
		if floatValue, err := strconv.ParseFloat(stringValue, 64); err == nil {
			return floatValue, nil
		}
	}

	if err != nil {
		return 0, err
	}

	return 0, fmt.Errorf("not a float")
}
