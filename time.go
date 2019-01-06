package jsonvalidator

import (
	"github.com/asaskevich/govalidator"
	"github.com/tidwall/gjson"
)

//DateTimeFormat Creates a new constraint for validating date and time againts a given format
func DateTimeFormat(format, message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if IsEmpty(value) {
			return
		}

		if !IsString(value) || !govalidator.IsTime(value.String(), format) {
			violations.Add(field, message)
		}
	})
}

//DateTime Creates a new constraint for validating date and time
func DateTime(message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if IsEmpty(value) {
			return
		}

		if !IsString(value) || !govalidator.IsTime(value.String(), validator.DateTimeFormat) {
			violations.Add(field, message)
		}
	})
}

//Date Creates a new constraint for validate date
func Date(message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if IsEmpty(value) {
			return
		}

		if !IsString(value) || !govalidator.IsTime(value.String(), validator.DateFormat) {
			violations.Add(field, message)
		}
	})
}

//Time Creates a new constraint for validating time
func Time(message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if IsEmpty(value) {
			return
		}

		if !IsString(value) || !govalidator.IsTime(value.String(), validator.TimeFormat) {
			violations.Add(field, message)
		}
	})
}
