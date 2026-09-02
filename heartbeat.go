package asynq

import (
	"sync"
	"time"

	"github.com/hibiken/asynq/internal/base"
	"github.com/hibiken/asynq/internal/log"
	"github.com/hibiken/asynq/internal/timeutil"
)

type heartbeater struct {
	logger *log.Logger
	broker base.Broker
	clock  timeutil.Clock

	done chan struct{}

	interval time.Duration

	host           string
	pid            int
	serverID       string
	concurrency    int
	queues         map[string]int
	strictPriority bool

	started time.Time
	workers map[string]*workerInfo

	state *serverState

	starting <-chan *workerInfo
	finished <-chan *base.TaskMessage
}

type heartbeaterParams struct {
	logger         *log.Logger
	broker         base.Broker
	interval       time.Duration
	concurrency    int
	queues         map[string]int
	strictPriority bool
	state          *serverState
	starting       <-chan *workerInfo
	finished       <-chan *base.TaskMessage
}

func newHeartbeater(params heartbeaterParams) *heartbeater { _ = "STUB: not implemented"; return nil }

func (h *heartbeater) shutdown() { _ = "STUB: not implemented"; return }

type workerInfo struct {
	msg *base.TaskMessage

	started time.Time

	deadline time.Time

	lease *base.Lease
}

func (h *heartbeater) start(wg *sync.WaitGroup) { _ = "STUB: not implemented"; return }

func (h *heartbeater) beat() { _ = "STUB: not implemented"; return }
