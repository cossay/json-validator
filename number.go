package jsonvalidator

import (
	"github.com/asaskevich/govalidator"

	"github.com/tidwall/gjson"
)

//IsNumber Checks if a given value is a number
func IsNumber(value *gjson.Result) bool {
	return value.Type == gjson.Number
}

//IsInt Checks if a given value is an integer
func IsInt(value *gjson.Result) bool {
	return IsNumber(value) && govalidator.IsInt(value.String())
}

//TypeInt Creates new validation constraint to check if a given value is an integer
func TypeInt(message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if IsEmpty(value) || IsInt(value) {
			return
		}

		violations.Add(field, message)
	})
}

//LessThan Creates new validation constraint to check if a given number is less than  agiven number
func LessThan(limit float64, message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if IsEmpty(value) || !IsNumber(value) {
			return
		}

		if value.Float() >= limit {
			violations.Add(field, message)
		}
	})
}

//LessThanOrEqual Creates new validation constraint to check if a given number is less than or equal to agiven number
func LessThanOrEqual(limit float64, message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if IsEmpty(value) || !IsNumber(value) {
			return
		}

		if value.Float() > limit {
			violations.Add(field, message)
		}
	})
}

//GreaterThan Creates new validation constraint to check if a given number is greater than  agiven number
func GreaterThan(limit float64, message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if IsEmpty(value) || !IsNumber(value) {
			return
		}

		if value.Float() <= limit {
			violations.Add(field, message)
		}
	})
}

//GreaterThanOrEqual Creates new validation constraint to check if a given number is greater than or equal to agiven number
func GreaterThanOrEqual(limit float64, message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if IsEmpty(value) || !IsNumber(value) {
			return
		}

		if value.Float() < limit {
			violations.Add(field, message)
		}
	})
}

//EqualTo Creates new validation constraint to check if a given number is equal to agiven number
func EqualTo(limit float64, message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if IsEmpty(value) || !IsNumber(value) {
			return
		}

		if value.Float() != limit {
			violations.Add(field, message)
		}
	})
}
