package asynq

import (
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/hibiken/asynq/internal/log"
	"github.com/hibiken/asynq/internal/rdb"
	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	id string

	state *serverState

	heartbeatInterval time.Duration
	logger            *log.Logger
	client            *Client
	rdb               *rdb.RDB
	cron              *cron.Cron
	location          *time.Location
	done              chan struct{}
	wg                sync.WaitGroup
	preEnqueueFunc    func(task *Task, opts []Option)
	postEnqueueFunc   func(info *TaskInfo, err error)
	errHandler        func(task *Task, opts []Option, err error)

	mu sync.Mutex

	idmap map[string]cron.EntryID
}

const defaultHeartbeatInterval = 10 * time.Second

func NewScheduler(r RedisConnOpt, opts *SchedulerOpts) *Scheduler {
	_ = "STUB: not implemented"
	return nil
}

func NewSchedulerFromRedisClient(c redis.UniversalClient, opts *SchedulerOpts) *Scheduler {
	_ = "STUB: not implemented"
	return nil
}

func newScheduler(opts *SchedulerOpts) *Scheduler { _ = "STUB: not implemented"; return nil }

func generateSchedulerID() string { _ = "STUB: not implemented"; return "" }

type SchedulerOpts struct {
	HeartbeatInterval time.Duration

	Logger Logger

	LogLevel LogLevel

	Location *time.Location

	PreEnqueueFunc func(task *Task, opts []Option)

	PostEnqueueFunc func(info *TaskInfo, err error)

	EnqueueErrorHandler func(task *Task, opts []Option, err error)
}

type enqueueJob struct {
	id              uuid.UUID
	cronspec        string
	task            *Task
	opts            []Option
	location        *time.Location
	logger          *log.Logger
	client          *Client
	rdb             *rdb.RDB
	preEnqueueFunc  func(task *Task, opts []Option)
	postEnqueueFunc func(info *TaskInfo, err error)
	errHandler      func(task *Task, opts []Option, err error)
}

func (j *enqueueJob) Run() { _ = "STUB: not implemented"; return }

func (s *Scheduler) Register(cronspec string, task *Task, opts ...Option) (entryID string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Scheduler) Unregister(entryID string) error { _ = "STUB: not implemented"; return nil }

func (s *Scheduler) Run() error { _ = "STUB: not implemented"; return nil }

func (s *Scheduler) Start() error { _ = "STUB: not implemented"; return nil }

func (s *Scheduler) start() error { _ = "STUB: not implemented"; return nil }

func (s *Scheduler) Shutdown() { _ = "STUB: not implemented"; return }

func (s *Scheduler) runHeartbeater() { _ = "STUB: not implemented"; return }

func (s *Scheduler) beat() { _ = "STUB: not implemented"; return }

func stringifyOptions(opts []Option) []string { _ = "STUB: not implemented"; return nil }

func (s *Scheduler) clearHistory() { _ = "STUB: not implemented"; return }

func (s *Scheduler) Ping() error { _ = "STUB: not implemented"; return nil }
