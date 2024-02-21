package main

import (
	"flag"
	"log"
	"os"
)

var serveAs string

func main() {
	flag.StringVar(&serveAs, "serve", "app", "")
	flag.Parse()

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	app := application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	switch serveAs {
	case "app":
		errorLog.Fatal(app.serveApp())
	default:
		errorLog.Fatal(app.serveWorker())
	}
}
