package asynq

import (
	"sync"
	"time"

	"github.com/hibiken/asynq/internal/log"
)

type syncer struct {
	logger *log.Logger

	requestsCh <-chan *syncRequest

	done chan struct{}

	interval time.Duration
}

type syncRequest struct {
	fn       func() error
	errMsg   string
	deadline time.Time
}

type syncerParams struct {
	logger     *log.Logger
	requestsCh <-chan *syncRequest
	interval   time.Duration
}

func newSyncer(params syncerParams) *syncer { _ = "STUB: not implemented"; return nil }

func (s *syncer) shutdown() { _ = "STUB: not implemented"; return }

func (s *syncer) start(wg *sync.WaitGroup) { _ = "STUB: not implemented"; return }
