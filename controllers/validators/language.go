package validators

import (
	config2 "Rabbit-OJ-Backend/services/judger/config"
	"github.com/go-playground/validator/v10"
)

func Language(fieldLevel validator.FieldLevel) bool {
	field := fieldLevel.Field()

	if language, ok := field.Interface().(string); ok {
		supportLanguage := false

		for _, lang := range config2.SupportLanguage {
			if language == lang.Value {
				supportLanguage = true
				break
			}
		}
		return supportLanguage
	}
	return false
}
