package jsonvalidator

import (
	"errors"
	"time"

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
	Validate(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator)
}

//ValidatorFunc Validation function
type ValidatorFunc func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator)

//Rule Rule
type Rule struct {
	validator ValidatorFunc
}

//Validate Validates
func (r *Rule) Validate(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
	if nil == r.validator {
		return
	}
	r.validator(field, value, parent, source, violations, validator)
}

//NewRule NewRule
func NewRule(validator ValidatorFunc) *Rule {
	return &Rule{validator: validator}
}

//Validator Validator service
type Validator struct {
	DateTimeFormat string
	TimeFormat     string
	DateFormat     string
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
			rule.Validate(field, &value, &data, &data, violations, v)
		}
	}

	return violations, nil
}

//NewValidator Creates a new validator with default configuration
func NewValidator() *Validator {
	return &Validator{
		DateTimeFormat: time.RFC3339,
		TimeFormat:     "15:04:05",
		DateFormat:     "2006-01-02",
	}
}
