package asynq

import (
	"sync"
	"time"

	"github.com/hibiken/asynq/internal/base"
	"github.com/hibiken/asynq/internal/log"
)

type forwarder struct {
	logger *log.Logger
	broker base.Broker

	done chan struct{}

	queues []string

	avgInterval time.Duration
}

type forwarderParams struct {
	logger   *log.Logger
	broker   base.Broker
	queues   []string
	interval time.Duration
}

func newForwarder(params forwarderParams) *forwarder { _ = "STUB: not implemented"; return nil }

func (f *forwarder) shutdown() { _ = "STUB: not implemented"; return }

func (f *forwarder) start(wg *sync.WaitGroup) { _ = "STUB: not implemented"; return }

func (f *forwarder) exec() { _ = "STUB: not implemented"; return }
