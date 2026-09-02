package asynq

import (
	"sync"
	"time"

	"github.com/hibiken/asynq/internal/base"
	"github.com/hibiken/asynq/internal/errors"
	"github.com/hibiken/asynq/internal/log"
)

type recoverer struct {
	logger         *log.Logger
	broker         base.Broker
	retryDelayFunc RetryDelayFunc
	isFailureFunc  func(error) bool

	done chan struct{}

	queues []string

	interval time.Duration
}

type recovererParams struct {
	logger         *log.Logger
	broker         base.Broker
	queues         []string
	interval       time.Duration
	retryDelayFunc RetryDelayFunc
	isFailureFunc  func(error) bool
}

func newRecoverer(params recovererParams) *recoverer { _ = "STUB: not implemented"; return nil }

func (r *recoverer) shutdown() { _ = "STUB: not implemented"; return }

func (r *recoverer) start(wg *sync.WaitGroup) { _ = "STUB: not implemented"; return }

var ErrLeaseExpired = errors.New("asynq: task lease expired")

func (r *recoverer) recover() { _ = "STUB: not implemented"; return }

func (r *recoverer) recoverLeaseExpiredTasks() { _ = "STUB: not implemented"; return }

func (r *recoverer) recoverStaleAggregationSets() { _ = "STUB: not implemented"; return }

func (r *recoverer) retry(msg *base.TaskMessage, err error) { _ = "STUB: not implemented"; return }

func (r *recoverer) archive(msg *base.TaskMessage, err error) { _ = "STUB: not implemented"; return }
