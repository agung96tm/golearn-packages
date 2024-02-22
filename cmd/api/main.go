package main

import (
	"flag"
	"github.com/agung96tm/golearn-packages/internal/models"
	"github.com/agung96tm/golearn-packages/lib"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	db       lib.Database
	models   models.Model
}

var dbDSN string

func main() {
	flag.StringVar(&dbDSN, "db-dsn", "", "postgres DSN")
	flag.Parse()

	// log
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// db
	db, err := lib.NewDB(dbDSN)
	if err != nil {
		errorLog.Fatal(err)
	}

	app := application{
		errorLog: errorLog,
		infoLog:  infoLog,
		db:       *db,
		models:   models.New(db),
	}

	srv := http.Server{
		Addr:     ":8001",
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting Server on port :%d\n", 8001)
	errorLog.Fatal(srv.ListenAndServe())
}
