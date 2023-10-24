package jsonld_helper

import (
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
	expanded, err := ld.NewJsonLdProcessor().Expand(value, options)

	if err != nil {
		return nil, err
	}

	return of(expanded), nil
}
