package main

import (
	"flag"
	"github.com/agung96tm/golearn-packages/internal/queue"
	"log"
	"os"
)

var serveAs string

func main() {
	flag.StringVar(&serveAs, "serve", "app", "")
	flag.Parse()

	// log
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	app := application{
		errorLog: errorLog,
		infoLog:  infoLog,
		queue:    queue.New("127.0.0.1:6379"),
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
