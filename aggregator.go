package asynq

import (
	"sync"
	"time"

	"github.com/hibiken/asynq/internal/base"
	"github.com/hibiken/asynq/internal/log"
)

type aggregator struct {
	logger *log.Logger
	broker base.Broker
	client *Client

	done chan struct{}

	queues []string

	gracePeriod time.Duration
	maxDelay    time.Duration
	maxSize     int

	ga GroupAggregator

	interval time.Duration

	sema chan struct{}
}

type aggregatorParams struct {
	logger          *log.Logger
	broker          base.Broker
	queues          []string
	gracePeriod     time.Duration
	maxDelay        time.Duration
	maxSize         int
	groupAggregator GroupAggregator
}

const (
	maxConcurrentAggregationChecks = 3

	defaultAggregationCheckInterval = 7 * time.Second
)

func newAggregator(params aggregatorParams) *aggregator { _ = "STUB: not implemented"; return nil }

func (a *aggregator) shutdown() { _ = "STUB: not implemented"; return }

func (a *aggregator) start(wg *sync.WaitGroup) { _ = "STUB: not implemented"; return }

func (a *aggregator) exec(t time.Time) { _ = "STUB: not implemented"; return }

func (a *aggregator) aggregate(t time.Time) { _ = "STUB: not implemented"; return }
