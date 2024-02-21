package main

import (
	"encoding/json"
	"github.com/hibiken/asynq"
	"time"
)

const EmailDeliveryTask = "email:deliver"

type EmailDeliveryPayload struct {
	PostID uint `json:"post_id"`
}

func (app application) runEmailDeliveryTask(postID uint) error {
	payload, err := json.Marshal(EmailDeliveryPayload{PostID: postID})
	if err != nil {
		return err
	}

	if _, err = app.queue.Client.Enqueue(
		asynq.NewTask(EmailDeliveryTask, payload),
		asynq.ProcessIn(1*time.Minute),
	); err != nil {
		return err
	}

	return nil
}
