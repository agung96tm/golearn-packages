package main

import (
	"github.com/agung96tm/golearn-packages/internal/queue"
	"github.com/hibiken/asynq"
	"log"
	"net/http"
	"time"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	queue    queue.Queue
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
	err := srv.ListenAndServe()

	if err != nil {
		_ = app.queue.Close()
	}
	return err
}

func (app application) serveWorker() error {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: app.queue.Addr},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6,
				"priority": 4,
				"default":  3,
				"low":      1,
			},
		},
	)

	return srv.Run(app.workers())
}

func (app application) serveScheduler() error {
	scheduler := asynq.NewScheduler(
		asynq.RedisClientOpt{Addr: app.queue.Addr},
		&asynq.SchedulerOpts{
			Location: time.Local,
		},
	)

	app.registerSchedulers(scheduler)

	return scheduler.Run()
}
