package jsonvalidator

import (
	"errors"

	"github.com/tidwall/gjson"
)

//Violations Validation errors
type Violations map[string]string

//Add Adds a new error for a given field
func (vs Violations) Add(field string, message string) {
	if nil == vs {
		vs = make(map[string]string, 0)
	}

	if !vs.Has(field) {
		vs[field] = message
	}
}

//Has Checks if a field is in the violation
func (vs Violations) Has(field string) bool {
	_, has := vs[field]
	return has
}

//Constraint Validation constraint
type Constraint interface {
	Validate(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations)
}

//ValidatorFunc Validation function
type ValidatorFunc func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations)

//Rule Rule
type Rule struct {
	validator ValidatorFunc
}

//Validate Validates
func (r *Rule) Validate(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations) {
	if nil == r.validator {
		return
	}
	r.validator(field, value, parent, source, violations)
}

//NewRule NewRule
func NewRule(validator ValidatorFunc) *Rule {
	return &Rule{validator: validator}
}

//Validator Validator service
type Validator struct {
}

//Validate Validates a given data againts a set of rules
func (v *Validator) Validate(input []byte, constraints map[string][]Constraint) (*Violations, error) {

	if !gjson.ValidBytes(input) {
		return nil, errors.New("Invalid JSON document. ")
	}

	data := gjson.ParseBytes(input)
	violations := &Violations{}

	for field, rules := range constraints {
		for _, rule := range rules {
			value := data.Get(field)
			rule.Validate(field, &value, &data, &data, violations)
		}
	}

	return violations, nil
}
