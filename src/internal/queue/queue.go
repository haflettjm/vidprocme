package queue

import (
	"vidprocme/internal/config"
	"vidprocme/internal/utils"

	"go.uber.org/zap"
)

type Queue struct {
	logger *zap.Logger
	Jobs   []utils.Job
	config *config.Config
}

func New(logger *zap.Logger, config *config.Config) *Queue {
	return &Queue{
		logger: logger,
		Jobs:   []utils.Job{},
		config: config,
	}
}

func (q *Queue) Enqueue(job utils.Job) {
	q.Jobs = append(q.Jobs, job)
}

func (q *Queue) Dequeue() utils.Job {
	if len(q.Jobs) == 0 {
		return utils.Job{}
	}
	job := q.Jobs[0]
	q.Jobs = q.Jobs[1:]
	return job
}

func (q *Queue) updateJob(job utils.Job) {
	q.logger.Info("Updating job", zap.String("job_id", job.ID))
	for i, j := range q.Jobs {
		if j.ID == job.ID {
			q.Jobs[i] = job
			return
		}
	}
}
