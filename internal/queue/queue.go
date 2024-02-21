package queue

import "github.com/hibiken/asynq"

type Queue struct {
	Client    *asynq.Client
	Scheduler *asynq.Scheduler
	Addr      string
}

func New(addr string) Queue {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: addr})

	return Queue{
		Client: client,
		Addr:   addr,
	}
}

func (q *Queue) Close() error {
	return q.Client.Close()
}
