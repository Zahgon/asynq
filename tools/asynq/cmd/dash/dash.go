package dash

import (
	"time"

	"github.com/hibiken/asynq"
)

type viewType int

const (
	viewTypeQueues viewType = iota
	viewTypeQueueDetails
	viewTypeHelp
)

type State struct {
	queues []*asynq.QueueInfo
	tasks  []*asynq.TaskInfo
	groups []*asynq.GroupInfo
	err    error

	queueTableRowIdx int
	taskTableRowIdx  int
	groupTableRowIdx int
	taskState        asynq.TaskState
	taskID           string

	selectedQueue *asynq.QueueInfo
	selectedGroup *asynq.GroupInfo
	selectedTask  *asynq.TaskInfo

	pageNum int

	view     viewType
	prevView viewType
}

func (s *State) DebugString() string { _ = "STUB: not implemented"; return "" }

type Options struct {
	DebugMode    bool
	PollInterval time.Duration
	RedisConnOpt asynq.RedisConnOpt
}

func Run(opts Options) { _ = "STUB: not implemented"; return }
