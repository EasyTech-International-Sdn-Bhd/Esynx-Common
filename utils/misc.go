package utils

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

func TrimWhitespace(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	trimmedVal := strings.TrimSpace(val)
	fl.Field().SetString(trimmedVal)
	return true
}
