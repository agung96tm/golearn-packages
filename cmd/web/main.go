package main

import (
	"github.com/agung96tm/golearn-packages/internal/validator"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type application struct {
	errorLog       *log.Logger
	infoLog        *log.Logger
	templateCache  map[string]*template.Template
	debug          bool
	formDecoder    *form.Decoder
	validator      *validator.Validator
	sessionManager *scs.SessionManager
}

func main() {
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// template
	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	// validator
	v, err := validator.NewValidator()
	if err != nil {
		errorLog.Fatal(err)
	}

	// sessions
	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour

	app := &application{
		errorLog:       errorLog,
		infoLog:        infoLog,
		templateCache:  templateCache,
		formDecoder:    form.NewDecoder(),
		validator:      v,
		sessionManager: sessionManager,
	}

	srv := http.Server{
		Addr:     ":8000",
		ErrorLog: app.errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting Server on port :%d\n", 8000)
	errorLog.Fatal(srv.ListenAndServe())
}
