package main

import (
	"fmt"
	"os"
	router "vidprocme/internal/api"
	"vidprocme/internal/config"

	"github.com/google/uuid"
	"github.com/uber-go/zap"
)

type Handler struct{}

type Queue struct {
	logger *zap.Logger
	Jobs   []Job
}

type Job struct {
	ID       string
	Data     string
	Status   string
	Worker   string
	WorkerID string
	Topic    string
}

type Scheduler struct {
	logger *zap.Logger
	queue  *Queue
	cfg    *config.Config
}

func newQueue(logger *zap.Logger) *Queue {
	return &Queue{
		logger: logger,
	}
}

func newJob(data string, topic string) Job {
	return Job{
		ID:       uuid.New().String(),
		Data:     "",
		Status:   "pending",
		Worker:   "",
		WorkerID: "",
		Topic:    topic,
	}
}

func (q *Queue) addJob(job Job) {
	q.Jobs = append(q.Jobs, job)
}

func (q *Queue) updateJob(job Job) {
	for i, j := range q.Jobs {
		if j.ID == job.ID {
			q.Jobs[i] = job
			return
		}
	}
}

func newScheduler(logger *zap.Logger, queue *Queue, cfg *config.Config) *Scheduler {
	return &Scheduler{
		logger: logger,
		queue:  queue,
		cfg:    cfg,
	}
}

func main() {
	// Initialize the configuration
	cfg := &config.Config{}
	if err := cfg.Load(); err != nil {
		fmt.Println("Error loading configuration:", err)
		os.Exit(1)
	}

	logger := zap.NewExample()
	queue := newQueue(logger)
	sched := newScheduler(logger, queue, cfg)

	fmt.Println("Server started")
	router.StartServer(cfg, queue, sched, logger)

}
