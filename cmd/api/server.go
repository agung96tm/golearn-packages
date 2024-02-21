package main

import (
	"github.com/hibiken/asynq"
	"log"
	"net/http"
	"time"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	runAs    string
}

func (app application) serveApp() error {
	srv := &http.Server{
		Addr: ":8001",

		ErrorLog: app.errorLog,
		Handler:  app.routes(),

		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.infoLog.Printf("Starting Server on port :%d\n", 8001)
	return srv.ListenAndServe()
}

func (app application) serveWorker() error {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "127.0.0.1:6379"},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)

	return srv.Run(app.workers())
}
