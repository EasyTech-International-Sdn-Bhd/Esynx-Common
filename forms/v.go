package forms

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/utils"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("trim", utils.TrimWhitespace)
}
