package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// SetupValidPassword sets up the validator for the password field
func SetupValidPassword(validate *validator.Validate) {
	uppercasePattern := `[A-Z]`
	lowercasePattern := `[a-z]`
	digitPattern := `[0-9]`
	specialCharacterPattern := `[!@#$%^&*()_+{}|:<>?~]`

	uppercaseRegex := regexp.MustCompile(uppercasePattern)
	lowercaseRegex := regexp.MustCompile(lowercasePattern)
	digitRegex := regexp.MustCompile(digitPattern)
	specialCharacterRegex := regexp.MustCompile(specialCharacterPattern)

	validate.RegisterValidation("validpasswd", func(fl validator.FieldLevel) bool {
		input := fl.Field().String()
		return specialCharacterRegex.MatchString(input) && digitRegex.MatchString(input) && uppercaseRegex.MatchString(input) && lowercaseRegex.MatchString(input)
	})
	validate.RegisterAlias("pass", "min=8,max=64,validpasswd")
}
