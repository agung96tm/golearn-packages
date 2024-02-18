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
	models   models.Models
}

func main() {
	dbConfig := lib.DatabaseConfig{}
	flag.StringVar(&dbConfig.DSN, "db-dsn", "", "DSN Address")
	flag.Parse()

	// log
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// db
	db, err := lib.NewDB(dbConfig)
	if err != nil {
		errorLog.Fatalf("[database] cannot connect with error: %v", err)
	}

	app := application{
		errorLog: errorLog,
		infoLog:  infoLog,
		models:   models.NewModels(*db),
	}

	srv := http.Server{
		Addr:     ":8001",
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting Server on port :%d\n", 8001)
	errorLog.Fatal(srv.ListenAndServe())
}
