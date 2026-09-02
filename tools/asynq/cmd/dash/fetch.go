package dash

import (
	"github.com/gdamore/tcell/v2"
	"github.com/hibiken/asynq"
)

type fetcher interface {
	Fetch(state *State)
}

type dataFetcher struct {
	inspector *asynq.Inspector
	opts      Options
	s         tcell.Screen

	errorCh  chan<- error
	queueCh  chan<- *asynq.QueueInfo
	taskCh   chan<- *asynq.TaskInfo
	queuesCh chan<- []*asynq.QueueInfo
	groupsCh chan<- []*asynq.GroupInfo
	tasksCh  chan<- []*asynq.TaskInfo
}

func (f *dataFetcher) Fetch(state *State) { _ = "STUB: not implemented"; return }

func (f *dataFetcher) fetchQueues() { _ = "STUB: not implemented"; return }

func fetchQueues(i *asynq.Inspector, queuesCh chan<- []*asynq.QueueInfo, errorCh chan<- error, opts Options) {
	_ = "STUB: not implemented"
	return
}

func fetchQueueInfo(i *asynq.Inspector, qname string, queueCh chan<- *asynq.QueueInfo, errorCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

func (f *dataFetcher) fetchGroups(qname string) { _ = "STUB: not implemented"; return }

func fetchGroups(i *asynq.Inspector, qname string, groupsCh chan<- []*asynq.GroupInfo, errorCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

func (f *dataFetcher) fetchAggregatingTasks(qname, group string, pageSize, pageNum int) {
	_ = "STUB: not implemented"
	return
}

func fetchAggregatingTasks(i *asynq.Inspector, qname, group string, pageSize, pageNum int,
	tasksCh chan<- []*asynq.TaskInfo, errorCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

func (f *dataFetcher) fetchTasks(qname string, taskState asynq.TaskState, pageSize, pageNum int) {
	_ = "STUB: not implemented"
	return
}

func fetchTasks(i *asynq.Inspector, qname string, taskState asynq.TaskState, pageSize, pageNum int,
	tasksCh chan<- []*asynq.TaskInfo, errorCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

func (f *dataFetcher) fetchTaskInfo(qname, taskID string) { _ = "STUB: not implemented"; return }

func fetchTaskInfo(i *asynq.Inspector, qname, taskID string, taskCh chan<- *asynq.TaskInfo, errorCh chan<- error) {
	_ = "STUB: not implemented"
	return
}
