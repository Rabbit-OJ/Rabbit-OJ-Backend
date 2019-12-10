package validators

import (
	"Rabbit-OJ-Backend/utils"
	"gopkg.in/go-playground/validator.v8"
	"reflect"
)

func Language(
	_ *validator.Validate, _ reflect.Value, _ reflect.Value,
	field reflect.Value,
	_ reflect.Type, _ reflect.Kind, _ string,
) bool {
	if language, ok := field.Interface().(string); ok {
		supportLanguage := false

		for _, lang := range utils.SupportLanguage {
			if language == lang.Value {
				supportLanguage = true
				break
			}
		}
		return supportLanguage
	}
	return false
}
