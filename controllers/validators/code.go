package validators

import (
	"gopkg.in/go-playground/validator.v8"
	"reflect"
)

func Code(
	_ *validator.Validate, _ reflect.Value, _ reflect.Value,
	field reflect.Value,
	_ reflect.Type, _ reflect.Kind, _ string,
) bool {
	if code, ok := field.Interface().(string); ok {
		return len(code) <= 100 * 1024
	}
	return false
}
