package helper

import "github.com/go-playground/validator/v10"

type (
	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}

	Xvalidator struct {
		validator *validator.Validate
	}
)

func NewValidator() *Xvalidator {
	return &Xvalidator{
		validator: validator.New(),
	}
}


var validate = validator.New()

func (v Xvalidator) Validator(data interface{}) []ErrorResponse{
	var validationErrors []ErrorResponse

	err := validate.Struct(data)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			error := ErrorResponse{
				Error: 	 true,
				FailedField: err.Field(),
				Tag:         err.Tag(),
				Value:       err.Param(),
			}
			validationErrors = append(validationErrors, error)
		}
		
	}

	return validationErrors
	
}