package utils

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

func ValidateReq(req interface{}) interface{} {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return errorMessage(err, req)
	}

	return nil
}

func errorMessage(err error, req interface{}) interface{} {
	msg := ""
	var errorValidation = map[string]string{}
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			field, _ := reflect.TypeOf(req).Elem().FieldByName(err.Field())

			var key string
			for _, source := range []string{"json", "param", "query"} {
				key = field.Tag.Get(source)
				if len(key) > 0 && key != "-" {
					break
				}
			}

			switch err.Tag() {
			case "required", "required_if", "required_unless":
				errorValidation[field.Tag.Get("json")] = field.Tag.Get("required")
			case "min":
				errorValidation[key] = field.Tag.Get("min")
			}
		}
	}

	for _, validate := range errorValidation {
		msg = validate
		break
	}

	return struct {
		Message         string            `json:"message"`
		ErrorValidation map[string]string `json:"error_validation"`
	}{
		Message:         msg,
		ErrorValidation: errorValidation,
	}
}
