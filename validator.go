package jsonvalidator

import (
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/tidwall/gjson"
)

//Violations Validation errors
type Violations struct {
	sync.RWMutex
	errors map[string]string
}

//MarshalJSON JSON marshaller
func (vs *Violations) MarshalJSON() ([]byte, error) {
	return json.Marshal(vs.errors)
}

//String Returns string representation of violoations
func (vs *Violations) String() string {
	j, _ := vs.MarshalJSON()
	return string(j)
}

//Add Adds a new error for a given field
func (vs *Violations) Add(field string, message string) {
	if !vs.Has(field) {
		vs.Lock()
		vs.errors[field] = message
		vs.Unlock()
	}
}

//Has Checks if a field is in the violation
func (vs *Violations) Has(field string) bool {
	vs.RLock()
	_, has := vs.errors[field]
	vs.RUnlock()
	return has
}

//Get Returns violation for a given field
func (vs *Violations) Get(field string) string {
	if !vs.Has(field) {
		return ""
	}

	vs.RLock()
	message := vs.errors[field]
	vs.RUnlock()
	return message
}

//First Return first violation
func (vs *Violations) First() string {
	vs.RLock()
	f := ""
	for _, message := range vs.errors {
		f = message
		break
	}
	vs.RUnlock()

	return f
}

//Errors Returns all violations
func (vs *Violations) Errors() map[string]string {
	return vs.errors
}

//NewViolations Creates a new violations
func NewViolations() *Violations {
	return &Violations{errors: make(map[string]string)}
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
	if nil == r.validator || violations.Has(field) {
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
func (v *Validator) Validate(document []byte, constraints map[string][]Constraint) *Violations {
	data := gjson.ParseBytes(document)
	violations := NewViolations()

	for field, rules := range constraints {
		for _, rule := range rules {
			value := data.Get(field)
			rule.Validate(field, &value, &data, &data, violations, v)
		}
	}

	return violations
}

//CheckDocument Checks validity of a JSON document
func (v *Validator) CheckDocument(document []byte) error {
	if !gjson.ValidBytes(document) {
		return errors.New("Invalid JSON document")
	}

	return nil
}

//NewValidator Creates a new validator with default configuration
func NewValidator() *Validator {
	return &Validator{
		DateTimeFormat: time.RFC3339,
		TimeFormat:     "15:04:05",
		DateFormat:     "2006-01-02",
	}
}
