package utils

import (
	"github.com/google/uuid"
)

type Job struct {
	ID       string
	Data     string
	Status   string
	Worker   string
	WorkerID string
	Topic    string
}

func (Job) newJob(data string, topic string) Job {
	return Job{
		ID:       uuid.New().String(),
		Data:     "",
		Status:   "pending",
		Worker:   "",
		WorkerID: "",
		Topic:    topic,
	}
}
