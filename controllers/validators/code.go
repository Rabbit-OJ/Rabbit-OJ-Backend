package validators

import (
	"github.com/go-playground/validator/v10"
)

func Code(fieldLevel validator.FieldLevel) bool {
	field := fieldLevel.Field()

	if code, ok := field.Interface().(string); ok {
		return len(code) <= 100*1024
	}
	return false
}
