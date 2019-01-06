package jsonvalidator

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/tidwall/gjson"
)

//TimeConstraint Time validation constraint
type TimeConstraint struct {
	format  string
	message string
}

//Validate Validats that a given value is a time matching a given format
func (t *TimeConstraint) Validate(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations) {
	if IsEmpty(value) {
		return
	}

	if !IsString(value) || !govalidator.IsTime(value.String(), t.format) {
		violations.Add(field, t.message)
	}
}

//TypeTime Creates a new time constraint
func TypeTime(format, message string) *TimeConstraint {
	return &TimeConstraint{format: format, message: message}
}

//TypeDate Date
func TypeDate(message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations) {
		if IsEmpty(value) {
			return
		}

		if !IsString(value) || !govalidator.IsTime(value.String(), time.RFC3339) {
			violations.Add(field, message)
		}
	})
}
