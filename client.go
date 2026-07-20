package asynq

import (
	"context"
	"time"

	"github.com/hibiken/asynq/internal/base"
	"github.com/hibiken/asynq/internal/errors"
	"github.com/redis/go-redis/v9"
)

type Client struct {
	broker base.Broker

	sharedConnection bool
}

func NewClient(r RedisConnOpt) *Client { _ = "STUB: not implemented"; return nil }

func NewClientFromRedisClient(c redis.UniversalClient) *Client {
	_ = "STUB: not implemented"
	return nil
}

type OptionType int

const (
	MaxRetryOpt OptionType = iota
	QueueOpt
	TimeoutOpt
	DeadlineOpt
	UniqueOpt
	ProcessAtOpt
	ProcessInOpt
	TaskIDOpt
	RetentionOpt
	GroupOpt
	HeaderOpt
)

type Option interface {
	String() string

	Type() OptionType

	Value() interface{}
}

type (
	retryOption     int
	queueOption     string
	taskIDOption    string
	timeoutOption   time.Duration
	deadlineOption  time.Time
	uniqueOption    time.Duration
	processAtOption time.Time
	processInOption time.Duration
	retentionOption time.Duration
	groupOption     string
	headerOption    [2]string
)

func MaxRetry(n int) Option { _ = "STUB: not implemented"; return *new(Option) }

func (n retryOption) String() string     { _ = "STUB: not implemented"; return "" }
func (n retryOption) Type() OptionType   { _ = "STUB: not implemented"; return *new(OptionType) }
func (n retryOption) Value() interface{} { _ = "STUB: not implemented"; return nil }

func Queue(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

func (name queueOption) String() string     { _ = "STUB: not implemented"; return "" }
func (name queueOption) Type() OptionType   { _ = "STUB: not implemented"; return *new(OptionType) }
func (name queueOption) Value() interface{} { _ = "STUB: not implemented"; return nil }

func TaskID(id string) Option { _ = "STUB: not implemented"; return *new(Option) }

func (id taskIDOption) String() string     { _ = "STUB: not implemented"; return "" }
func (id taskIDOption) Type() OptionType   { _ = "STUB: not implemented"; return *new(OptionType) }
func (id taskIDOption) Value() interface{} { _ = "STUB: not implemented"; return nil }

func Timeout(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func (d timeoutOption) String() string     { _ = "STUB: not implemented"; return "" }
func (d timeoutOption) Type() OptionType   { _ = "STUB: not implemented"; return *new(OptionType) }
func (d timeoutOption) Value() interface{} { _ = "STUB: not implemented"; return nil }

func Deadline(t time.Time) Option { _ = "STUB: not implemented"; return *new(Option) }

func (t deadlineOption) String() string { _ = "STUB: not implemented"; return "" }

func (t deadlineOption) Type() OptionType   { _ = "STUB: not implemented"; return *new(OptionType) }
func (t deadlineOption) Value() interface{} { _ = "STUB: not implemented"; return nil }

func Unique(ttl time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func (ttl uniqueOption) String() string     { _ = "STUB: not implemented"; return "" }
func (ttl uniqueOption) Type() OptionType   { _ = "STUB: not implemented"; return *new(OptionType) }
func (ttl uniqueOption) Value() interface{} { _ = "STUB: not implemented"; return nil }

func ProcessAt(t time.Time) Option { _ = "STUB: not implemented"; return *new(Option) }

func (t processAtOption) String() string { _ = "STUB: not implemented"; return "" }

func (t processAtOption) Type() OptionType   { _ = "STUB: not implemented"; return *new(OptionType) }
func (t processAtOption) Value() interface{} { _ = "STUB: not implemented"; return nil }

func ProcessIn(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func (d processInOption) String() string     { _ = "STUB: not implemented"; return "" }
func (d processInOption) Type() OptionType   { _ = "STUB: not implemented"; return *new(OptionType) }
func (d processInOption) Value() interface{} { _ = "STUB: not implemented"; return nil }

func Retention(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func (ttl retentionOption) String() string     { _ = "STUB: not implemented"; return "" }
func (ttl retentionOption) Type() OptionType   { _ = "STUB: not implemented"; return *new(OptionType) }
func (ttl retentionOption) Value() interface{} { _ = "STUB: not implemented"; return nil }

func Group(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

func (name groupOption) String() string     { _ = "STUB: not implemented"; return "" }
func (name groupOption) Type() OptionType   { _ = "STUB: not implemented"; return *new(OptionType) }
func (name groupOption) Value() interface{} { _ = "STUB: not implemented"; return nil }

func Header(key, value string) Option { _ = "STUB: not implemented"; return *new(Option) }

func (h headerOption) String() string { _ = "STUB: not implemented"; return "" }

func (h headerOption) Type() OptionType   { _ = "STUB: not implemented"; return *new(OptionType) }
func (h headerOption) Value() interface{} { _ = "STUB: not implemented"; return nil }

var ErrDuplicateTask = errors.New("task already exists")

var ErrTaskIDConflict = errors.New("task ID conflicts with another task")

type option struct {
	retry     int
	queue     string
	taskID    string
	timeout   time.Duration
	deadline  time.Time
	uniqueTTL time.Duration
	processAt time.Time
	retention time.Duration
	group     string
	headers   map[string]string
}

func composeOptions(opts ...Option) (option, error) {
	_ = "STUB: not implemented"
	return *new(option), nil
}

func isBlank(s string) bool { _ = "STUB: not implemented"; return false }

const (
	defaultMaxRetry = 25

	defaultTimeout = 30 * time.Minute
)

var (
	noTimeout  time.Duration = 0
	noDeadline time.Time     = time.Unix(0, 0)
)

func (c *Client) Close() error { _ = "STUB: not implemented"; return nil }

func (c *Client) Enqueue(task *Task, opts ...Option) (*TaskInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) EnqueueContext(ctx context.Context, task *Task, opts ...Option) (*TaskInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type BatchEnqueueResult struct {
	TaskInfo *TaskInfo
	Err      error
}

func (c *Client) BatchEnqueueContext(ctx context.Context, tasks []*Task, opts ...Option) []BatchEnqueueResult {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) Ping() error { _ = "STUB: not implemented"; return nil }

func (c *Client) enqueue(ctx context.Context, msg *base.TaskMessage, uniqueTTL time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) schedule(ctx context.Context, msg *base.TaskMessage, t time.Time, uniqueTTL time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) addToGroup(ctx context.Context, msg *base.TaskMessage, group string, uniqueTTL time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}
