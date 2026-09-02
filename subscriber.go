package asynq

import (
	"sync"
	"time"

	"github.com/hibiken/asynq/internal/base"
	"github.com/hibiken/asynq/internal/log"
)

type subscriber struct {
	logger *log.Logger
	broker base.Broker

	done chan struct{}

	cancelations *base.Cancelations

	retryTimeout time.Duration
}

type subscriberParams struct {
	logger       *log.Logger
	broker       base.Broker
	cancelations *base.Cancelations
}

func newSubscriber(params subscriberParams) *subscriber { _ = "STUB: not implemented"; return nil }

func (s *subscriber) shutdown() { _ = "STUB: not implemented"; return }

func (s *subscriber) start(wg *sync.WaitGroup) { _ = "STUB: not implemented"; return }
