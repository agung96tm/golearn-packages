package main

import (
	"github.com/hibiken/asynq"
	"log"
)

const Task1ScheduleTask = "task1"

// cron help: https://crontab.guru/#*_*_*_*_*

func (app application) registerSchedulers(scheduler *asynq.Scheduler) {
	if _, err := scheduler.Register(
		"* * * * *", // minutes
		asynq.NewTask(Task1ScheduleTask, nil),
	); err != nil {
		log.Fatal(err)
	}
}
