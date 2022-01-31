package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/jancajthaml-go/luhn"
)

type Validator struct {
	validate *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validate.Struct(i)
}

func ValidateLuhn(fl validator.FieldLevel) bool {
	return luhn.Validate(fl.Field().String())
}

func New() (*Validator, error) {
	v := validator.New()
	err := v.RegisterValidation("luhn", ValidateLuhn)
	if err != nil {
		return nil, err
	}
	return &Validator{v}, nil
}
