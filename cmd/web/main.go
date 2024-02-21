package main

import (
	"flag"
	"github.com/agung96tm/golearn-packages/internal/queue"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	"log"
	"os"
	"time"
)

var serveAs string

func main() {
	flag.StringVar(&serveAs, "serve", "app", "")
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

	app := &application{
		errorLog:       errorLog,
		infoLog:        infoLog,
		templateCache:  templateCache,
		formDecoder:    form.NewDecoder(),
		sessionManager: sessionManager,
		queue:          queue.New("127.0.0.1:6379"),
	}

	switch serveAs {
	case "app":
		errorLog.Fatal(app.serveApp())
	case "worker":
		errorLog.Fatal(app.serveWorker())
	case "scheduler":
		errorLog.Fatal(app.serveScheduler())
	}
}
