package validator

import "github.com/asaskevich/govalidator"

var defaultValidator *Validator

type Validator struct{}

// i 必须为结构体 或者结构体指针
func (Validator) Validate(i interface{}) error {
	ok, err := govalidator.ValidateStruct(i)
	if !ok {
		return err
	}
	return nil
}

func Instance() *Validator {
	if defaultValidator == nil {
		defaultValidator = new(Validator)
	}

	return defaultValidator
}

// add more
