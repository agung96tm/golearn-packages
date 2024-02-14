package form

type IForm interface {
	SetErrFields(map[string][]string)
	IsValid() bool
}

type Form struct {
	ErrFields map[string][]string
}

func (f *Form) IsValid() bool {
	return len(f.ErrFields) == 0
}

func (f *Form) SetErrFields(fields map[string][]string) {
	f.ErrFields = fields
}
