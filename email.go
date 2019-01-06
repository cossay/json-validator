package jsonvalidator

import (
	"github.com/asaskevich/govalidator"
	"github.com/tidwall/gjson"
)

//Email Creates a new email validation constraint
func Email(message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if IsEmpty(value) {
			return
		}

		if !IsString(value) || !govalidator.IsEmail(value.String()) {
			violations.Add(field, message)
		}
	})
}

//ExistingEmail Creates a constraint for validating a given value as an email of an existing domain
func ExistingEmail(message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if IsEmpty(value) {
			return
		}

		if !IsString(value) || !govalidator.IsExistingEmail(value.String()) {
			violations.Add(field, message)
		}
	})
}
