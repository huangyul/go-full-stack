package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func ValidateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	// ^(((13[0-9]{1})|(15[0-9]{1})|(18[0-9]{1}))+\d{8})$
	ok, _ := regexp.MatchString("^(((13[0-9]{1})|(15[0-9]{1})|(18[0-9]{1}))+\\d{8})$", mobile)
	if !ok {
		return false
	}
	return true
}
