package main

import (
	"bytes"
	"errors"
	"fmt"
	appForm "github.com/agung96tm/golearn-packages/internal/form"
	"github.com/go-playground/form/v4"
	"github.com/go-playground/validator/v10"
	"github.com/justinas/nosurf"
	"net/http"
	"runtime/debug"
	"time"
)

func (app *application) newTemplateData(r *http.Request) *templateData {
	return &templateData{
		CSRFToken:   nosurf.Token(r),
		Flash:       app.sessionManager.PopString(r.Context(), "flash"),
		CurrentYear: time.Now().Year(),
	}
}

func (app *application) PostForm(r *http.Request, dst appForm.IForm) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	err = app.formDecoder.Decode(dst, r.PostForm)
	if err != nil {
		var invalidDecoderError *form.InvalidDecoderError
		if errors.As(err, &invalidDecoderError) {
			panic(err)
		}
		return err
	}

	err = app.validator.Validate.Struct(dst)
	if err != nil {
		var errFields = make(map[string][]string)

		var errs validator.ValidationErrors
		errors.As(err, &errs)
		for _, e := range errs {
			key := e.Field()
			if _, exists := errFields[key]; !exists {
				errFields[key] = make([]string, 0)
			}
			errFields[key] = append(errFields[key], e.Translate(app.validator.Trans))
		}

		dst.SetErrFields(errFields)
	}

	return nil
}

func (app *application) render(w http.ResponseWriter, status int, page string, data *templateData) {
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, err)
		return
	}

	buf := new(bytes.Buffer)
	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(status)
	buf.WriteTo(w)
}

func (app *application) redirect(w http.ResponseWriter, r *http.Request, url string) {
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	if app.debug {
		http.Error(w, trace, http.StatusInternalServerError)
		return
	}

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
