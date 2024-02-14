package validator

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslator "github.com/go-playground/validator/v10/translations/en"
	"reflect"
	"strings"
)

type Validator struct {
	Validate *validator.Validate
	Trans    ut.Translator
}

func NewValidator() (*Validator, error) {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		nameTag := fld.Tag.Get("json")
		if nameTag != "" {
			name := strings.SplitN(nameTag, ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		}
		return fld.Name
	})

	newEn := en.New()
	uni := ut.New(newEn, newEn)
	trans, _ := uni.GetTranslator("en")
	err := enTranslator.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return nil, err
	}

	return &Validator{
		Validate: validate,
		Trans:    trans,
	}, nil
}

func (v *Validator) Form(form any) {

}
