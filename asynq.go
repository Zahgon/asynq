package asynq

import (
	"context"
	"crypto/tls"
	"net/url"
	"time"

	"github.com/hibiken/asynq/internal/base"
)

type Task struct {
	typename string

	payload []byte

	headers map[string]string

	opts []Option

	w *ResultWriter
}

func (t *Task) Type() string               { _ = "STUB: not implemented"; return "" }
func (t *Task) Payload() []byte            { _ = "STUB: not implemented"; return nil }
func (t *Task) Headers() map[string]string { _ = "STUB: not implemented"; return nil }
func (t *Task) Options() []Option          { _ = "STUB: not implemented"; return nil }

func (t *Task) ResultWriter() *ResultWriter { _ = "STUB: not implemented"; return nil }

func NewTask(typename string, payload []byte, opts ...Option) *Task {
	_ = "STUB: not implemented"
	return nil
}

func NewTaskWithHeaders(typename string, payload []byte, headers map[string]string, opts ...Option) *Task {
	_ = "STUB: not implemented"
	return nil
}

func newTask(typename string, payload []byte, w *ResultWriter) *Task {
	_ = "STUB: not implemented"
	return nil
}

type TaskInfo struct {
	ID string

	Queue string

	Type string

	Payload []byte

	Headers map[string]string

	State TaskState

	MaxRetry int

	Retried int

	LastErr string

	LastFailedAt time.Time

	Timeout time.Duration

	Deadline time.Time

	Group string

	NextProcessAt time.Time

	IsOrphaned bool

	Retention time.Duration

	CompletedAt time.Time

	Result []byte
}

func fromUnixTimeOrZero(t int64) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func newTaskInfo(msg *base.TaskMessage, state base.TaskState, nextProcessAt time.Time, result []byte) *TaskInfo {
	_ = "STUB: not implemented"
	return nil
}

type TaskState int

const (
	TaskStateActive TaskState = iota + 1

	TaskStatePending

	TaskStateScheduled

	TaskStateRetry

	TaskStateArchived

	TaskStateCompleted

	TaskStateAggregating
)

func (s TaskState) String() string { _ = "STUB: not implemented"; return "" }

type RedisConnOpt interface {
	MakeRedisClient() interface{}
}

type RedisClientOpt struct {
	Network string

	Addr string

	Username string

	Password string

	DB int

	DialTimeout time.Duration

	ReadTimeout time.Duration

	WriteTimeout time.Duration

	PoolSize int

	TLSConfig *tls.Config
}

func (opt RedisClientOpt) MakeRedisClient() interface{} { _ = "STUB: not implemented"; return nil }

type RedisFailoverClientOpt struct {
	MasterName string

	SentinelAddrs []string

	SentinelUsername string

	SentinelPassword string

	Username string

	Password string

	DB int

	DialTimeout time.Duration

	ReadTimeout time.Duration

	WriteTimeout time.Duration

	PoolSize int

	TLSConfig *tls.Config
}

func (opt RedisFailoverClientOpt) MakeRedisClient() interface{} {
	_ = "STUB: not implemented"
	return nil
}

type RedisClusterClientOpt struct {
	Addrs []string

	MaxRedirects int

	Username string

	Password string

	DialTimeout time.Duration

	ReadTimeout time.Duration

	WriteTimeout time.Duration

	TLSConfig *tls.Config
}

func (opt RedisClusterClientOpt) MakeRedisClient() interface{} {
	_ = "STUB: not implemented"
	return nil
}

func ParseRedisURI(uri string) (RedisConnOpt, error) {
	_ = "STUB: not implemented"
	return *new(RedisConnOpt), nil
}

func parseRedisURI(u *url.URL) (RedisConnOpt, error) {
	_ = "STUB: not implemented"
	return *new(RedisConnOpt), nil
}

func parseRedisSocketURI(u *url.URL) (RedisConnOpt, error) {
	_ = "STUB: not implemented"
	return *new(RedisConnOpt), nil
}

func parseRedisSentinelURI(u *url.URL) (RedisConnOpt, error) {
	_ = "STUB: not implemented"
	return *new(RedisConnOpt), nil
}

type ResultWriter struct {
	id     string
	qname  string
	broker base.Broker
	ctx    context.Context
}

func (w *ResultWriter) Write(data []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (w *ResultWriter) TaskID() string { _ = "STUB: not implemented"; return "" }
