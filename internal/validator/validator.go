package validator

type Validator struct {
	ErrNonFields []string            `json:"err_non_fields"`
	ErrFields    map[string][]string `json:"err_fields"`
}

func (v *Validator) IsValid() bool {
	return len(v.ErrNonFields) == 0 && len(v.ErrFields) == 0
}

func (v *Validator) SetErrField(key string, message string) {
	if v.ErrFields == nil {
		v.ErrFields = make(map[string][]string)
	}

	if _, exists := v.ErrFields[key]; !exists {
		v.ErrFields[key] = make([]string, 0)
	}

	v.ErrFields[key] = append(v.ErrFields[key], message)
}

func (v *Validator) SetErrNonField(message string) {
	v.ErrNonFields = append(v.ErrNonFields, message)
}
