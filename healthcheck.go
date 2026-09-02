package asynq

import (
	"sync"
	"time"

	"github.com/hibiken/asynq/internal/base"
	"github.com/hibiken/asynq/internal/log"
)

type healthchecker struct {
	logger *log.Logger
	broker base.Broker

	done chan struct{}

	interval time.Duration

	healthcheckFunc func(error)
}

type healthcheckerParams struct {
	logger          *log.Logger
	broker          base.Broker
	interval        time.Duration
	healthcheckFunc func(error)
}

func newHealthChecker(params healthcheckerParams) *healthchecker {
	_ = "STUB: not implemented"
	return nil
}

func (hc *healthchecker) shutdown() { _ = "STUB: not implemented"; return }

func (hc *healthchecker) start(wg *sync.WaitGroup) { _ = "STUB: not implemented"; return }
