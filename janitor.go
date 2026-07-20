package asynq

import (
	"sync"
	"time"

	"github.com/hibiken/asynq/internal/base"
	"github.com/hibiken/asynq/internal/log"
)

type janitor struct {
	logger *log.Logger
	broker base.Broker

	done chan struct{}

	queues []string

	avgInterval time.Duration

	batchSize int
}

type janitorParams struct {
	logger    *log.Logger
	broker    base.Broker
	queues    []string
	interval  time.Duration
	batchSize int
}

func newJanitor(params janitorParams) *janitor { _ = "STUB: not implemented"; return nil }

func (j *janitor) shutdown() { _ = "STUB: not implemented"; return }

func (j *janitor) start(wg *sync.WaitGroup) { _ = "STUB: not implemented"; return }

func (j *janitor) exec() { _ = "STUB: not implemented"; return }
