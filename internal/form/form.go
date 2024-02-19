package form

import (
	"errors"
)

type Form struct {
	ErrFields    map[string][]string
	ErrNonFields []string
}

func (f *Form) IsValid() bool {
	return len(f.ErrFields) == 0 && len(f.ErrNonFields) == 0
}

func (f *Form) AddErrField(key string, message string) {
	if f.ErrFields == nil {
		f.ErrFields = make(map[string][]string)
	}
	if _, exists := f.ErrFields[key]; !exists {
		f.ErrFields[key] = append(f.ErrFields[key], message)
	}
}

func (f *Form) AddErrFields(fields map[string][]string) {
	f.ErrFields = fields
}

func (f *Form) AddErrNonField(message string) {
	f.ErrNonFields = append(f.ErrNonFields, message)
}

var ErrForm = errors.New("invalid form")
