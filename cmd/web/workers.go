package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
)

func (app application) workers() asynq.Handler {
	mux := asynq.NewServeMux()

	// task
	mux.HandleFunc(EmailArticleCreateTask, app.handleEmailArticleCreateTask)

	// scheduler
	mux.HandleFunc(Task1ScheduleTask, app.handleScheduleTask1)

	return mux
}

func (app application) handleScheduleTask1(ctx context.Context, t *asynq.Task) error {
	app.infoLog.Printf("Sending Scheduling Task 1")
	return nil
}

func (app application) handleEmailArticleCreateTask(ctx context.Context, t *asynq.Task) error {
	var payload EmailArticleCreateTaskPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	app.infoLog.Printf("Sending Email to Admin: about article withid=%d", payload.ArticleID)
	return nil
}
