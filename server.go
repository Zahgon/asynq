package asynq

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/hibiken/asynq/internal/base"
	"github.com/hibiken/asynq/internal/log"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	logger *log.Logger

	broker base.Broker

	sharedConnection bool

	state *serverState

	wg            sync.WaitGroup
	forwarder     *forwarder
	processor     *processor
	syncer        *syncer
	heartbeater   *heartbeater
	subscriber    *subscriber
	recoverer     *recoverer
	healthchecker *healthchecker
	janitor       *janitor
	aggregator    *aggregator
}

type serverState struct {
	mu    sync.Mutex
	value serverStateValue
}

type serverStateValue int

const (
	srvStateNew serverStateValue = iota

	srvStateActive

	srvStateStopped

	srvStateClosed
)

var serverStates = []string{
	"new",
	"active",
	"stopped",
	"closed",
}

func (s serverStateValue) String() string { _ = "STUB: not implemented"; return "" }

type Config struct {
	Concurrency int

	BaseContext func() context.Context

	TaskCheckInterval time.Duration

	RetryDelayFunc RetryDelayFunc

	IsFailure func(error) bool

	Queues map[string]int

	StrictPriority bool

	ErrorHandler ErrorHandler

	Logger Logger

	LogLevel LogLevel

	ShutdownTimeout time.Duration

	HealthCheckFunc func(error)

	HealthCheckInterval time.Duration

	DelayedTaskCheckInterval time.Duration

	GroupGracePeriod time.Duration

	GroupMaxDelay time.Duration

	GroupMaxSize int

	GroupAggregator GroupAggregator

	JanitorInterval time.Duration

	JanitorBatchSize int
}

type GroupAggregator interface {
	Aggregate(group string, tasks []*Task) *Task
}

type GroupAggregatorFunc func(group string, tasks []*Task) *Task

func (fn GroupAggregatorFunc) Aggregate(group string, tasks []*Task) *Task {
	_ = "STUB: not implemented"
	return nil
}

type ErrorHandler interface {
	HandleError(ctx context.Context, task *Task, err error)
}

type ErrorHandlerFunc func(ctx context.Context, task *Task, err error)

func (fn ErrorHandlerFunc) HandleError(ctx context.Context, task *Task, err error) {
	_ = "STUB: not implemented"
	return
}

type RetryDelayFunc func(n int, e error, t *Task) time.Duration

type Logger interface {
	Debug(args ...interface{})

	Info(args ...interface{})

	Warn(args ...interface{})

	Error(args ...interface{})

	Fatal(args ...interface{})
}

type LogLevel int32

const (
	level_unspecified LogLevel = iota

	DebugLevel

	InfoLevel

	WarnLevel

	ErrorLevel

	FatalLevel
)

func (l *LogLevel) String() string { _ = "STUB: not implemented"; return "" }

func (l *LogLevel) Set(val string) error { _ = "STUB: not implemented"; return nil }

func toInternalLogLevel(l LogLevel) log.Level { _ = "STUB: not implemented"; return *new(log.Level) }

func DefaultRetryDelayFunc(n int, e error, t *Task) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func defaultIsFailureFunc(err error) bool { _ = "STUB: not implemented"; return false }

var defaultQueueConfig = map[string]int{
	base.DefaultQueueName: 1,
}

const (
	defaultTaskCheckInterval = 1 * time.Second

	defaultShutdownTimeout = 8 * time.Second

	defaultHealthCheckInterval = 15 * time.Second

	defaultDelayedTaskCheckInterval = 5 * time.Second

	defaultGroupGracePeriod = 1 * time.Minute

	defaultJanitorInterval = 8 * time.Second

	defaultJanitorBatchSize = 100
)

func NewServer(r RedisConnOpt, cfg Config) *Server { _ = "STUB: not implemented"; return nil }

func NewServerFromRedisClient(c redis.UniversalClient, cfg Config) *Server {
	_ = "STUB: not implemented"
	return nil
}

type Handler interface {
	ProcessTask(context.Context, *Task) error
}

type HandlerFunc func(context.Context, *Task) error

func (fn HandlerFunc) ProcessTask(ctx context.Context, task *Task) error {
	_ = "STUB: not implemented"
	return nil
}

var ErrServerClosed = errors.New("asynq: Server closed")

func (srv *Server) Run(handler Handler) error { _ = "STUB: not implemented"; return nil }

func (srv *Server) Start(handler Handler) error { _ = "STUB: not implemented"; return nil }

func (srv *Server) start() error { _ = "STUB: not implemented"; return nil }

func (srv *Server) Shutdown() { _ = "STUB: not implemented"; return }

func (srv *Server) Stop() { _ = "STUB: not implemented"; return }

func (srv *Server) Ping() error { _ = "STUB: not implemented"; return nil }
