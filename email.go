package jsonvalidator

import (
	"github.com/asaskevich/govalidator"
	"github.com/tidwall/gjson"
)

//EmailConstraint Email validation constraint
type EmailConstraint struct {
	message string
}

//Validate Validates an email
func (e *EmailConstraint) Validate(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations) {
	if isEmpty(value) {
		return
	}

	if !isString(value) || !govalidator.IsEmail(value.String()) {
		violations.Add(field, e.message)
	}
}

//Email Creates a new email validation constraint
func Email(message string) *EmailConstraint {
	return &EmailConstraint{message: message}
}
