package validator

import (
	"fmt"
)

type Validator struct {
	ErrNonFields []string            `json:"err_non_fields"`
	ErrFields    map[string][]string `json:"err_fields"`
}

func (v *Validator) IsValid() bool {
	return len(v.ErrNonFields) == 0 && len(v.ErrFields) == 0
}

func (v *Validator) AddErrField(key string, message string) {
	if v.ErrFields == nil {
		v.ErrFields = make(map[string][]string)
	}

	if _, exists := v.ErrFields[key]; !exists {
		v.ErrFields[key] = make([]string, 0)
	}

	v.ErrFields[key] = append(v.ErrFields[key], message)
}

func (v *Validator) AddErrNonField(message string) {
	v.ErrNonFields = append(v.ErrNonFields, message)
}

type ErrValidator struct {
	Fields    map[string][]string `json:"fields,omitempty"`
	NonFields []string            `json:"non_fields,omitempty"`
}

func (e ErrValidator) Error() string {
	return fmt.Sprintf("%#v", e)
}
