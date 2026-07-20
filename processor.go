package asynq

import (
	"context"
	"sync"
	"time"

	"github.com/hibiken/asynq/internal/base"
	"github.com/hibiken/asynq/internal/errors"
	"github.com/hibiken/asynq/internal/log"
	"github.com/hibiken/asynq/internal/timeutil"
	"golang.org/x/time/rate"
)

type processor struct {
	logger *log.Logger
	broker base.Broker
	clock  timeutil.Clock

	handler   Handler
	baseCtxFn func() context.Context

	queueConfig map[string]int

	orderedQueues []string

	taskCheckInterval time.Duration
	retryDelayFunc    RetryDelayFunc
	isFailureFunc     func(error) bool

	errHandler      ErrorHandler
	shutdownTimeout time.Duration

	syncRequestCh chan<- *syncRequest

	errLogLimiter *rate.Limiter

	sema chan struct{}

	done chan struct{}
	once sync.Once

	quit chan struct{}

	abort chan struct{}

	cancelations *base.Cancelations

	starting chan<- *workerInfo
	finished chan<- *base.TaskMessage
}

type processorParams struct {
	logger            *log.Logger
	broker            base.Broker
	baseCtxFn         func() context.Context
	retryDelayFunc    RetryDelayFunc
	taskCheckInterval time.Duration
	isFailureFunc     func(error) bool
	syncCh            chan<- *syncRequest
	cancelations      *base.Cancelations
	concurrency       int
	queues            map[string]int
	strictPriority    bool
	errHandler        ErrorHandler
	shutdownTimeout   time.Duration
	starting          chan<- *workerInfo
	finished          chan<- *base.TaskMessage
}

func newProcessor(params processorParams) *processor { _ = "STUB: not implemented"; return nil }

func (p *processor) stop() { _ = "STUB: not implemented"; return }

func (p *processor) shutdown() { _ = "STUB: not implemented"; return }

func (p *processor) start(wg *sync.WaitGroup) { _ = "STUB: not implemented"; return }

func (p *processor) exec() { _ = "STUB: not implemented"; return }

func (p *processor) requeue(l *base.Lease, msg *base.TaskMessage) {
	_ = "STUB: not implemented"
	return
}

func (p *processor) handleSucceededMessage(l *base.Lease, msg *base.TaskMessage) {
	_ = "STUB: not implemented"
	return
}

func (p *processor) markAsComplete(l *base.Lease, msg *base.TaskMessage) {
	_ = "STUB: not implemented"
	return
}

func (p *processor) markAsDone(l *base.Lease, msg *base.TaskMessage) {
	_ = "STUB: not implemented"
	return
}

var SkipRetry = errors.New("skip retry for the task")

var RevokeTask = errors.New("revoke task")

func (p *processor) handleFailedMessage(ctx context.Context, l *base.Lease, msg *base.TaskMessage, err error) {
	_ = "STUB: not implemented"
	return
}

func (p *processor) retry(l *base.Lease, msg *base.TaskMessage, e error, isFailure bool) {
	_ = "STUB: not implemented"
	return
}

func (p *processor) archive(l *base.Lease, msg *base.TaskMessage, e error) {
	_ = "STUB: not implemented"
	return
}

func (p *processor) queues() []string { _ = "STUB: not implemented"; return nil }

func (p *processor) perform(ctx context.Context, task *Task) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func uniq(names []string, l int) []string { _ = "STUB: not implemented"; return nil }

func sortByPriority(qcfg map[string]int) []string { _ = "STUB: not implemented"; return nil }

type queue struct {
	name     string
	priority int
}

type byPriority []*queue

func (x byPriority) Len() int           { _ = "STUB: not implemented"; return 0 }
func (x byPriority) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
func (x byPriority) Swap(i, j int)      { _ = "STUB: not implemented"; return }

func normalizeQueues(queues map[string]int) map[string]int { _ = "STUB: not implemented"; return nil }

func gcd(xs ...int) int { _ = "STUB: not implemented"; return 0 }

func (p *processor) computeDeadline(msg *base.TaskMessage) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func IsPanicError(err error) bool { _ = "STUB: not implemented"; return false }
