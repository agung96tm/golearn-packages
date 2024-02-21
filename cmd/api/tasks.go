package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
)

const EmailDeliveryTask = "email:deliver"

type EmailDeliveryPayload struct {
	PostID uint `json:"post_id"`
}

func RunEmailDeliveryTask(postID uint) (*asynq.Task, error) {
	payload, err := json.Marshal(EmailDeliveryPayload{PostID: postID})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(EmailDeliveryTask, payload), nil
}

func (app application) handleEmailDeliveryTask(ctx context.Context, t *asynq.Task) error {
	var p EmailDeliveryPayload

	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	app.infoLog.Printf("Sending Email to Admin: post_id=%d", p.PostID)
	return nil
}
