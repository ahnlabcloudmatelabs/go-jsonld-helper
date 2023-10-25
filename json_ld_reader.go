package jsonld_helper

import (
	"encoding/json"
	"reflect"

	"github.com/piprate/json-gold/ld"
)

type JsonLDReader interface {
	Value() any
	Length() int

	ReadKey(string) JsonLDReader
	ReadIndex(int) JsonLDReader

	Get() any
	GetOrElse(any) any
	GetOrThrow(error) (any, error)

	StringOrElse(string) string
	StringOrThrow(error) (string, error)

	BoolOrElse(bool) bool
	BoolOrThrow(error) (bool, error)

	IntOrElse(int) int
	IntOrThrow(error) (int, error)

	FloatOrElse(float64) float64
	FloatOrThrow(error) (float64, error)
}

func ParseJsonLD(value any, options *ld.JsonLdOptions) (JsonLDReader, error) {
	origin := value
	reflectType := reflect.TypeOf(value).String()

	if reflectType == "[]uint8" {
		if err := json.Unmarshal(value.([]byte), &origin); err != nil {
			return nil, err
		}
	}

	if reflectType == "string" {
		if err := json.Unmarshal([]byte(value.(string)), &origin); err != nil {
			return nil, err
		}
	}

	expanded, err := ld.NewJsonLdProcessor().Expand(origin, options)

	if err != nil {
		return nil, err
	}

	return of(expanded), nil
}
