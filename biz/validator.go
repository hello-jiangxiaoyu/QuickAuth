package internal

import (
	"QuickAuth/pkg/utils"
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
)

func GetValidator() (*validator.Validate, error) {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return nil, errors.New("validator.Engin err")
	}

	if err := v.RegisterValidation("safe_url", ValidSafeURL); err != nil {
		return nil, err
	}

	if err := v.RegisterValidation("safe_string", ValidSafeString); err != nil {
		return nil, err
	}

	if err := v.RegisterValidation("safe_string_list", ValidSafeStrings); err != nil {
		return nil, err
	}

	return v, nil
}

func ValidSafeURL(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	if value == "" {
		return true
	}
	return utils.IsSafeURLValid(value)
}

func ValidSafeString(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	return utils.IsSafeStringValid(value, 255)
}

func ValidSafeStrings(fl validator.FieldLevel) bool {
	value := fl.Field()
	switch value.Kind() {
	case reflect.Slice:
		if value.Len() > 64 {
			return false
		}

		for i := 0; i < value.Len(); i++ {
			if !utils.IsSafeStringValid(value.Index(i).String(), 255) {
				return false
			}
		}
	default:
		return false
	}

	return true
}
