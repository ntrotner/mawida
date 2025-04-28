package validators

import (
	"github.com/go-playground/validator/v10"
)

// SetupValidDynamicAttributes sanitizes the dynamic attributes field by only allowing strings as keys and values
func SetupValidDynamicAttributes(validate *validator.Validate) {
	validate.RegisterValidation("dynamicAttributes", func(fl validator.FieldLevel) bool {
		dynamicAttributes := fl.Field().Interface().(map[string]interface{})
		
		for key, value := range dynamicAttributes {
			valueStr, valueOk := value.(string)

			if !valueOk || key == "" || valueStr == "" {
				return false
			}
		}

		return true
	})
}
