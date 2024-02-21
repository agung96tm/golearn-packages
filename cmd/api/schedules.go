package main

import (
	"github.com/hibiken/asynq"
	"log"
)

const Task1ScheduleTask = "task1"

func (app application) scheduleRegister(scheduler *asynq.Scheduler) {
	if _, err := scheduler.Register(
		"@every 30s",
		asynq.NewTask(Task1ScheduleTask, nil),
	); err != nil {
		log.Fatal(err)
	}
}
