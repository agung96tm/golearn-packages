package main

import (
	"github.com/hibiken/asynq"
)

func (app application) workers() asynq.Handler {
	mux := asynq.NewServeMux()

	mux.HandleFunc(EmailDeliveryTask, app.handleEmailDeliveryTask)

	return mux
}
