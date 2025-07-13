package queue

import (
	"vidprocme/src/internal/config"
	"vidprocme/src/internal/utils"

	"github.com/uber-go/zap"
)

type Queue struct {
	logger *zap.Logger
	Jobs   []utils.Job
	config config.Config
}

type Scheduler struct {
	logger *zap.Logger
	queue  *Queue
	cfg    *config.Config
}

func (Queue) Init(logger *zap.Logger) *Queue {
	return &Queue{
		logger: logger,
		Jobs:   []utils.Job{},
	}
}

func (q *Queue) addJob(job utils.Job) {
	q.Jobs = append(q.Jobs, job)
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
