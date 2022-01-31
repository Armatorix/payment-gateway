package validator

import (
	"time"

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

func ValidateExpiration(fl validator.FieldLevel) bool {
	year := fl.Field().FieldByName("Year").Int()
	month := fl.Field().FieldByName("Month").Int()
	return time.Now().Before(time.Date(int(year), time.Month(month), 1, 0, 0, 0, 0, time.Local))
}

func New() (*Validator, error) {
	v := validator.New()
	err := v.RegisterValidation("luhn", ValidateLuhn)
	if err != nil {
		return nil, err
	}
	err = v.RegisterValidation("not_expired", ValidateExpiration)
	if err != nil {
		return nil, err
	}
	return &Validator{v}, nil
}
