package validators

import (
	fields "template_backend/infrastructure/validators/fields"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var Validate = validator.New(validator.WithRequiredStructEnabled())

// SetupValidator initializes the validator with the necessary fields
func SetupValidator() {
	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(Validate, trans)

	fields.SetupValidPassword(Validate)
}
