package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
)

func (app application) workers() asynq.Handler {
	mux := asynq.NewServeMux()

	mux.HandleFunc(EmailDeliveryTask, app.handleEmailDeliveryTask)
	mux.HandleFunc(Task1ScheduleTask, app.handleScheduleTask1)

	return mux
}

func (app application) handleScheduleTask1(ctx context.Context, t *asynq.Task) error {
	app.infoLog.Printf("Sending Scheduling Task 1")
	return nil
}

func (app application) handleEmailDeliveryTask(ctx context.Context, t *asynq.Task) error {
	var p EmailDeliveryPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	app.infoLog.Printf("Sending Email to Admin: post_id=%d", p.PostID)
	return nil
}
