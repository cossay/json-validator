package jsonvalidator

import (
	"github.com/asaskevich/govalidator"

	"github.com/tidwall/gjson"
)

//IsIntConstraint Constraint for checking if a given value is an integer
type IsIntConstraint struct {
	message string
}

//Validate Validates that a given value is an integer
func (i *IsIntConstraint) Validate(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations) {
	if !value.Exists() || govalidator.IsInt(value.String()) {
		return
	}

	violations.Add(field, i.message)
}

//IsInt Creates new validation constraint to check if a given value is an integer
func IsInt(message string) *IsIntConstraint {
	return &IsIntConstraint{message: message}
}

//LessThanConstraint Constraint for checking if a given value is a number less than a given limit
type LessThanConstraint struct {
	message string
	limit   float64
}

//Validate Validates that a given number is less than or equal to a given number
func (lt *LessThanConstraint) Validate(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations) {

	if !value.Exists() || !govalidator.IsNumeric(value.String()) {
		return
	}

	if value.Float() >= lt.limit {
		violations.Add(field, lt.message)
	}
}

//LessThan Creates new validation constraint to check if a given number is less than  agiven number
func LessThan(limit float64, message string) *LessThanConstraint {
	return &LessThanConstraint{limit: limit, message: message}
}

//LessThanOrEqualConstraint Constraint for checking if a given value is a number less than or equal to a given limit
type LessThanOrEqualConstraint struct {
	message string
	limit   float64
}

//Validate Validates that a given number is less than or equal to a given number
func (lte *LessThanOrEqualConstraint) Validate(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations) {

	if !value.Exists() || !govalidator.IsNumeric(value.String()) {
		return
	}

	if value.Float() > lte.limit {
		violations.Add(field, lte.message)
	}
}

//LessThanOrEqual Creates new validation constraint to check if a given number is less than or equal to agiven number
func LessThanOrEqual(limit float64, message string) *LessThanOrEqualConstraint {
	return &LessThanOrEqualConstraint{limit: limit, message: message}
}

//GreaterThanConstraint Constraint for checking if a given value is a number greater than a given limit
type GreaterThanConstraint struct {
	message string
	limit   float64
}

//Validate Validates that a given number is greater than a given number
func (gt *GreaterThanConstraint) Validate(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations) {

	if !value.Exists() || !govalidator.IsNumeric(value.String()) {
		return
	}

	if value.Float() <= gt.limit {
		violations.Add(field, gt.message)
	}
}

//GreaterThan Creates new validation constraint to check if a given number is greater than  agiven number
func GreaterThan(limit float64, message string) *GreaterThanConstraint {
	return &GreaterThanConstraint{limit: limit, message: message}
}

//GreaterThanOrEqualConstraint Constraint for checking if a given value is a number greater than or equal to a given limit
type GreaterThanOrEqualConstraint struct {
	message string
	limit   float64
}

//Validate Validates that a given number is greater or equal to a given number
func (lte *GreaterThanOrEqualConstraint) Validate(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations) {

	if !value.Exists() || !govalidator.IsNumeric(value.String()) {
		return
	}

	if value.Float() < lte.limit {
		violations.Add(field, lte.message)
	}
}

//GreaterThanOrEqual Creates new validation constraint to check if a given number is greater than or equal to agiven number
func GreaterThanOrEqual(limit float64, message string) *GreaterThanOrEqualConstraint {
	return &GreaterThanOrEqualConstraint{limit: limit, message: message}
}

//EqualToConstraint Constraint for checking if a given value is a number equal to a given limit
type EqualToConstraint struct {
	message string
	limit   float64
}

//Validate Validates that a given value is equal to a given number
func (et *EqualToConstraint) Validate(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations) {

	if !value.Exists() || !govalidator.IsNumeric(value.String()) {
		return
	}

	if value.Float() != et.limit {
		violations.Add(field, et.message)
	}
}

//EqualTo Creates new validation constraint to check if a given number is equal to agiven number
func EqualTo(limit float64, message string) *EqualToConstraint {
	return &EqualToConstraint{limit: limit, message: message}
}
