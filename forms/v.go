package forms

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/utils"
	"github.com/go-playground/validator/v10"
)

var FormValidator *validator.Validate

func init() {
	FormValidator = validator.New()
	FormValidator.RegisterValidation("trim", utils.TrimWhitespace)
}

func ValidateForm(i interface{}) error {
	err := FormValidator.Struct(&i)
	if err != nil {
		return err
	}
	return nil
}
