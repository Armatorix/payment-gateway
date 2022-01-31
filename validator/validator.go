package validator

import (
	"time"

	"github.com/Armatorix/payment-gateway/x/xtime"
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

func ValidateExpiration(fl validator.StructLevel) {
	my := fl.Current().Interface().(xtime.MonthYear)

	if !time.Now().Before(time.Date(int(my.Year), time.Month(my.Month), 1, 0, 0, 0, 0, time.Local)) {
		fl.ReportError(my, "Month/Year", "Expiration", "", "")
	}
}

func New() (*Validator, error) {
	v := validator.New()
	err := v.RegisterValidation("luhn", ValidateLuhn)
	if err != nil {
		return nil, err
	}

	v.RegisterStructValidation(ValidateExpiration, xtime.MonthYear{})
	if err != nil {
		return nil, err
	}
	return &Validator{v}, nil
}
