package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/agung96tm/golearn-packages/internal/validator"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func (app application) readIDParam(r *http.Request) (uint, error) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return uint(id), nil
}

func (app application) readJSON(w http.ResponseWriter, r *http.Request, dest any) error {
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(dest)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("body contains unknown key %s", fieldName)
		case err.Error() == "http: request body too large":
			return fmt.Errorf("body must not be larger than %d bytes", maxBytes)
		case errors.As(err, &invalidUnmarshalError):
			panic(err)
		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single JSON Value")
	}

	return nil
}

func (app application) writeJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func (app application) errorResponse(w http.ResponseWriter, status int, message any) {
	if err, ok := message.(error); ok {
		message = err.Error()
	}

	env := struct {
		Message any `json:"message"`
	}{Message: message}

	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.errorLog.Println(err)
		w.WriteHeader(500)
	}
}

func (app application) serverErrorResponse(w http.ResponseWriter, err error) {
	app.errorLog.Println(err)
	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, http.StatusInternalServerError, message)
}

func (app application) badRequestResponse(w http.ResponseWriter, err error) {
	if errValidator, ok := err.(validator.ErrValidator); ok {
		err := app.writeJSON(w, http.StatusBadRequest, errValidator, nil)
		if err != nil {
			app.serverErrorResponse(w, err)
		}
		return
	}
	app.errorResponse(w, http.StatusBadRequest, err)
}

func (app application) notFoundResponse(w http.ResponseWriter, message any) {
	app.errorResponse(w, http.StatusNotFound, message)
}
