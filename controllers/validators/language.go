package validators

import (
	"Rabbit-OJ-Backend/services/config"
	"github.com/go-playground/validator/v10"
)

func Language(fieldLevel validator.FieldLevel) bool {
	field := fieldLevel.Field()

	if language, ok := field.Interface().(string); ok {
		supportLanguage := false

		for _, lang := range config.SupportLanguage {
			if language == lang.Value {
				supportLanguage = true
				break
			}
		}
		return supportLanguage
	}
	return false
}
