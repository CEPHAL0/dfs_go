package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

type XValidator struct {
	validator *validator.Validate
}

type GlobalErrorHandlerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

var validate = validator.New()

func (v XValidator) Validate(data interface{}) []ValidationError {
	validationErrors := []ValidationError{}

	errs := validate.Struct(data)

	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ValidationError

			elem.FailedField = err.Field()
			elem.Tag = err.Tag()
			elem.Value = err.Value()
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func Validate(model interface{}) error {
	var returnErr error
	myValidator := &XValidator{validator: validate}

	errs := myValidator.Validate(model)

	if len(errs) > 0 && errs[0].Error {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
			returnErr = errors.New(strings.Join(errMsgs, "and "))
		}
	}
	return returnErr
}
