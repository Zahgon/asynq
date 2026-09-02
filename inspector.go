package asynq

import (
	"time"

	"github.com/hibiken/asynq/internal/errors"
	"github.com/hibiken/asynq/internal/rdb"
	"github.com/redis/go-redis/v9"
)

type Inspector struct {
	rdb *rdb.RDB

	sharedConnection bool
}

func NewInspector(r RedisConnOpt) *Inspector { _ = "STUB: not implemented"; return nil }

func NewInspectorFromRedisClient(c redis.UniversalClient) *Inspector {
	_ = "STUB: not implemented"
	return nil
}

func (i *Inspector) Close() error { _ = "STUB: not implemented"; return nil }

func (i *Inspector) Queues() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (i *Inspector) Groups(queue string) ([]*GroupInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type GroupInfo struct {
	Group string

	Size int
}

type QueueInfo struct {
	Queue string

	MemoryUsage int64

	Latency time.Duration

	Size int

	Groups int

	Pending int

	Active int

	Scheduled int

	Retry int

	Archived int

	Completed int

	Aggregating int

	Processed int

	Failed int

	ProcessedTotal int

	FailedTotal int

	Paused bool

	Timestamp time.Time
}

func (i *Inspector) GetQueueInfo(queue string) (*QueueInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type DailyStats struct {
	Queue string

	Processed int

	Failed int

	Date time.Time
}

func (i *Inspector) History(queue string, n int) ([]*DailyStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var (
	ErrQueueNotFound = errors.New("queue not found")

	ErrQueueNotEmpty = errors.New("queue is not empty")

	ErrTaskNotFound = errors.New("task not found")
)

func (i *Inspector) DeleteQueue(queue string, force bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *Inspector) GetTaskInfo(queue, id string) (*TaskInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ListOption interface{}

type (
	pageSizeOpt int
	pageNumOpt  int
)

type listOption struct {
	pageSize int
	pageNum  int
}

const (
	defaultPageSize = 30

	defaultPageNum = 1
)

func composeListOptions(opts ...ListOption) listOption {
	_ = "STUB: not implemented"
	return *new(listOption)
}

func PageSize(n int) ListOption { _ = "STUB: not implemented"; return *new(ListOption) }

func Page(n int) ListOption { _ = "STUB: not implemented"; return *new(ListOption) }

func (i *Inspector) ListPendingTasks(queue string, opts ...ListOption) ([]*TaskInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *Inspector) ListActiveTasks(queue string, opts ...ListOption) ([]*TaskInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *Inspector) ListAggregatingTasks(queue, group string, opts ...ListOption) ([]*TaskInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *Inspector) ListScheduledTasks(queue string, opts ...ListOption) ([]*TaskInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *Inspector) ListRetryTasks(queue string, opts ...ListOption) ([]*TaskInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *Inspector) ListArchivedTasks(queue string, opts ...ListOption) ([]*TaskInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *Inspector) ListCompletedTasks(queue string, opts ...ListOption) ([]*TaskInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *Inspector) DeleteAllPendingTasks(queue string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (i *Inspector) DeleteAllScheduledTasks(queue string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (i *Inspector) DeleteAllRetryTasks(queue string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (i *Inspector) DeleteAllArchivedTasks(queue string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (i *Inspector) DeleteAllCompletedTasks(queue string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (i *Inspector) DeleteAllAggregatingTasks(queue, group string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (i *Inspector) UpdateTaskPayload(queue, id string, payload []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *Inspector) DeleteTask(queue, id string) error { _ = "STUB: not implemented"; return nil }

func (i *Inspector) RunAllScheduledTasks(queue string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (i *Inspector) RunAllRetryTasks(queue string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (i *Inspector) RunAllArchivedTasks(queue string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (i *Inspector) RunAllAggregatingTasks(queue, group string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (i *Inspector) RunTask(queue, id string) error { _ = "STUB: not implemented"; return nil }

func (i *Inspector) ArchiveAllPendingTasks(queue string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (i *Inspector) ArchiveAllScheduledTasks(queue string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (i *Inspector) ArchiveAllRetryTasks(queue string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (i *Inspector) ArchiveAllAggregatingTasks(queue, group string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (i *Inspector) ArchiveTask(queue, id string) error { _ = "STUB: not implemented"; return nil }

func (i *Inspector) CancelProcessing(id string) error { _ = "STUB: not implemented"; return nil }

func (i *Inspector) PauseQueue(queue string) error { _ = "STUB: not implemented"; return nil }

func (i *Inspector) UnpauseQueue(queue string) error { _ = "STUB: not implemented"; return nil }

func (i *Inspector) Servers() ([]*ServerInfo, error) { _ = "STUB: not implemented"; return nil, nil }

type ServerInfo struct {
	ID string

	Host string

	PID int

	Concurrency    int
	Queues         map[string]int
	StrictPriority bool

	Started time.Time

	Status string

	ActiveWorkers []*WorkerInfo
}

type WorkerInfo struct {
	TaskID string

	TaskType string

	TaskPayload []byte

	Queue string

	Started time.Time

	Deadline time.Time
}

func (i *Inspector) ClusterKeySlot(queue string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type ClusterNode struct {
	ID string

	Addr string
}

func (i *Inspector) ClusterNodes(queue string) ([]*ClusterNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type SchedulerEntry struct {
	ID string

	Spec string

	Task *Task

	Opts []Option

	Next time.Time

	Prev time.Time
}

func (i *Inspector) SchedulerEntries() ([]*SchedulerEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseOption(s string) (Option, error) { _ = "STUB: not implemented"; return *new(Option), nil }

func parseOptionFunc(s string) string { _ = "STUB: not implemented"; return "" }

func parseOptionArg(s string) string { _ = "STUB: not implemented"; return "" }

type SchedulerEnqueueEvent struct {
	TaskID string

	EnqueuedAt time.Time
}

func (i *Inspector) ListSchedulerEnqueueEvents(entryID string, opts ...ListOption) ([]*SchedulerEnqueueEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
