package main

import (
	"github.com/agung96tm/golearn-packages/internal/validator"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog  *log.Logger
	infoLog   *log.Logger
	validator *validator.Validator
}

func main() {
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// validator
	v, err := validator.NewValidator()
	if err != nil {
		errorLog.Fatal(err)
	}

	app := application{
		errorLog:  errorLog,
		infoLog:   infoLog,
		validator: v,
	}

	srv := http.Server{
		Addr:     ":8001",
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting Server on port :%d\n", 8001)
	errorLog.Fatal(srv.ListenAndServe())
}
