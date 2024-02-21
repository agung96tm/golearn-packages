package main

import (
	"encoding/json"
	"github.com/hibiken/asynq"
	"time"
)

const EmailArticleCreateTask = "email:post-create"

type EmailArticleCreateTaskPayload struct {
	ArticleID uint `json:"article_id"`
}

func (app application) runEmailArticleCreateTask(articleID uint) {
	payload, err := json.Marshal(EmailArticleCreateTaskPayload{ArticleID: articleID})
	if err != nil {
		panic(err)
	}

	if _, err = app.queue.Client.Enqueue(
		asynq.NewTask(EmailArticleCreateTask, payload),
		asynq.ProcessIn(1*time.Minute),
		asynq.MaxRetry(3),
		asynq.Timeout(3*time.Minute),
		asynq.Queue("priority"),
	); err != nil {
		panic(err)
	}
}
