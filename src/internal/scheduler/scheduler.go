package scheduler

import (
	"vidprocme/internal/config"
	"vidprocme/internal/queue"

	"go.uber.org/zap"
)

type Scheduler struct {
	logger *zap.Logger
	queue  *queue.Queue
	cfg    *config.Config
}

func New(logger *zap.Logger, queue *queue.Queue, cfg *config.Config) *Scheduler {
	return &Scheduler{
		logger: logger,
		queue:  queue,
		cfg:    cfg,
	}
}

func (s *Scheduler) Start() error {
	s.logger.Info("Starting scheduler")

	return nil
}
