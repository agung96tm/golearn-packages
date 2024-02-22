package main

import (
	"flag"
	"github.com/agung96tm/golearn-packages/internal/models"
	"github.com/agung96tm/golearn-packages/lib"
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
	sessionManager *scs.SessionManager
	db             lib.Database
	models         models.Model
	storage        lib.Storage
}

var dbDSN string

func main() {
	flag.StringVar(&dbDSN, "db-dsn", "", "postgres DSN")
	flag.Parse()

	// log
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// template
	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	// sessions
	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour

	// db
	db, err := lib.NewDB(dbDSN)
	if err != nil {
		errorLog.Fatal(err)
	}

	storage := lib.NewStorage()

	app := &application{
		errorLog:       errorLog,
		infoLog:        infoLog,
		templateCache:  templateCache,
		formDecoder:    form.NewDecoder(),
		sessionManager: sessionManager,
		db:             *db,
		storage:        storage,
		models:         models.New(db, &storage),
	}

	srv := http.Server{
		Addr:     ":8000",
		ErrorLog: app.errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting Server on port :%d\n", 8000)
	errorLog.Fatal(srv.ListenAndServe())
}
