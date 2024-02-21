package queue

import "github.com/hibiken/asynq"

type Queue struct {
	Client *asynq.Client
}

func New(addr string) Queue {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: addr})

	return Queue{
		Client: client,
	}
}

func (q *Queue) Close() error {
	return q.Client.Close()
}
