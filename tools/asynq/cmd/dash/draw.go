package dash

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/hibiken/asynq"
)

var (
	baseStyle  = tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	labelStyle = baseStyle.Foreground(tcell.ColorLightGray)

	activeStyle      = baseStyle.Foreground(tcell.ColorBlue)
	pendingStyle     = baseStyle.Foreground(tcell.ColorGreen)
	aggregatingStyle = baseStyle.Foreground(tcell.ColorLightGreen)
	scheduledStyle   = baseStyle.Foreground(tcell.ColorYellow)
	retryStyle       = baseStyle.Foreground(tcell.ColorPink)
	archivedStyle    = baseStyle.Foreground(tcell.ColorPurple)
	completedStyle   = baseStyle.Foreground(tcell.ColorDarkGreen)
)

type drawer interface {
	Draw(state *State)
}

type dashDrawer struct {
	s    tcell.Screen
	opts Options
}

func (dd *dashDrawer) Draw(state *State) { _ = "STUB: not implemented"; return }

func drawQueueSizeGraphs(d *ScreenDrawer, state *State) { _ = "STUB: not implemented"; return }

func drawFooter(d *ScreenDrawer, state *State) { _ = "STUB: not implemented"; return }

func maxwidth(names []string) int { _ = "STUB: not implemented"; return 0 }

func rpad(s string, padding int) string { _ = "STUB: not implemented"; return "" }

func lpad(s string, padding int) string { _ = "STUB: not implemented"; return "" }

func byteCount(b int64) string { _ = "STUB: not implemented"; return "" }

var queueColumnConfigs = []*columnConfig[*asynq.QueueInfo]{
	{"Queue", alignLeft, func(q *asynq.QueueInfo) string { return q.Queue }},
	{"State", alignLeft, func(q *asynq.QueueInfo) string { return formatQueueState(q) }},
	{"Size", alignRight, func(q *asynq.QueueInfo) string { return strconv.Itoa(q.Size) }},
	{"Latency", alignRight, func(q *asynq.QueueInfo) string { return q.Latency.Round(time.Second).String() }},
	{"MemoryUsage", alignRight, func(q *asynq.QueueInfo) string { return byteCount(q.MemoryUsage) }},
	{"Processed", alignRight, func(q *asynq.QueueInfo) string { return strconv.Itoa(q.Processed) }},
	{"Failed", alignRight, func(q *asynq.QueueInfo) string { return strconv.Itoa(q.Failed) }},
	{"ErrorRate", alignRight, func(q *asynq.QueueInfo) string { return formatErrorRate(q.Processed, q.Failed) }},
}

func formatQueueState(q *asynq.QueueInfo) string { _ = "STUB: not implemented"; return "" }

func formatErrorRate(processed, failed int) string { _ = "STUB: not implemented"; return "" }

func formatNextProcessTime(t time.Time) string { _ = "STUB: not implemented"; return "" }

func formatPastTime(t time.Time) string { _ = "STUB: not implemented"; return "" }

func drawQueueTable(d *ScreenDrawer, style tcell.Style, state *State) {
	_ = "STUB: not implemented"
	return
}

func drawQueueSummary(d *ScreenDrawer, state *State) { _ = "STUB: not implemented"; return }

func groupPageSize(s tcell.Screen) int { _ = "STUB: not implemented"; return 0 }

func taskPageSize(s tcell.Screen) int { _ = "STUB: not implemented"; return 0 }

func shouldShowGroupTable(state *State) bool { _ = "STUB: not implemented"; return false }

func getTaskTableColumnConfig(taskState asynq.TaskState) []*columnConfig[*asynq.TaskInfo] {
	_ = "STUB: not implemented"
	return nil
}

var activeTaskTableColumns = []*columnConfig[*asynq.TaskInfo]{
	{"ID", alignLeft, func(t *asynq.TaskInfo) string { return t.ID }},
	{"Type", alignLeft, func(t *asynq.TaskInfo) string { return t.Type }},
	{"Retried", alignRight, func(t *asynq.TaskInfo) string { return strconv.Itoa(t.Retried) }},
	{"Max Retry", alignRight, func(t *asynq.TaskInfo) string { return strconv.Itoa(t.MaxRetry) }},
	{"Payload", alignLeft, func(t *asynq.TaskInfo) string { return formatByteSlice(t.Payload) }},
}

var pendingTaskTableColumns = []*columnConfig[*asynq.TaskInfo]{
	{"ID", alignLeft, func(t *asynq.TaskInfo) string { return t.ID }},
	{"Type", alignLeft, func(t *asynq.TaskInfo) string { return t.Type }},
	{"Retried", alignRight, func(t *asynq.TaskInfo) string { return strconv.Itoa(t.Retried) }},
	{"Max Retry", alignRight, func(t *asynq.TaskInfo) string { return strconv.Itoa(t.MaxRetry) }},
	{"Payload", alignLeft, func(t *asynq.TaskInfo) string { return formatByteSlice(t.Payload) }},
}

var aggregatingTaskTableColumns = []*columnConfig[*asynq.TaskInfo]{
	{"ID", alignLeft, func(t *asynq.TaskInfo) string { return t.ID }},
	{"Type", alignLeft, func(t *asynq.TaskInfo) string { return t.Type }},
	{"Payload", alignLeft, func(t *asynq.TaskInfo) string { return formatByteSlice(t.Payload) }},
	{"Group", alignLeft, func(t *asynq.TaskInfo) string { return t.Group }},
}

var scheduledTaskTableColumns = []*columnConfig[*asynq.TaskInfo]{
	{"ID", alignLeft, func(t *asynq.TaskInfo) string { return t.ID }},
	{"Type", alignLeft, func(t *asynq.TaskInfo) string { return t.Type }},
	{"Next Process Time", alignLeft, func(t *asynq.TaskInfo) string {
		return formatNextProcessTime(t.NextProcessAt)
	}},
	{"Payload", alignLeft, func(t *asynq.TaskInfo) string { return formatByteSlice(t.Payload) }},
}

var retryTaskTableColumns = []*columnConfig[*asynq.TaskInfo]{
	{"ID", alignLeft, func(t *asynq.TaskInfo) string { return t.ID }},
	{"Type", alignLeft, func(t *asynq.TaskInfo) string { return t.Type }},
	{"Retry", alignRight, func(t *asynq.TaskInfo) string { return fmt.Sprintf("%d/%d", t.Retried, t.MaxRetry) }},
	{"Last Failure", alignLeft, func(t *asynq.TaskInfo) string { return t.LastErr }},
	{"Last Failure Time", alignLeft, func(t *asynq.TaskInfo) string { return formatPastTime(t.LastFailedAt) }},
	{"Next Process Time", alignLeft, func(t *asynq.TaskInfo) string {
		return formatNextProcessTime(t.NextProcessAt)
	}},
	{"Payload", alignLeft, func(t *asynq.TaskInfo) string { return formatByteSlice(t.Payload) }},
}

var archivedTaskTableColumns = []*columnConfig[*asynq.TaskInfo]{
	{"ID", alignLeft, func(t *asynq.TaskInfo) string { return t.ID }},
	{"Type", alignLeft, func(t *asynq.TaskInfo) string { return t.Type }},
	{"Retry", alignRight, func(t *asynq.TaskInfo) string { return fmt.Sprintf("%d/%d", t.Retried, t.MaxRetry) }},
	{"Last Failure", alignLeft, func(t *asynq.TaskInfo) string { return t.LastErr }},
	{"Last Failure Time", alignLeft, func(t *asynq.TaskInfo) string { return formatPastTime(t.LastFailedAt) }},
	{"Payload", alignLeft, func(t *asynq.TaskInfo) string { return formatByteSlice(t.Payload) }},
}

var completedTaskTableColumns = []*columnConfig[*asynq.TaskInfo]{
	{"ID", alignLeft, func(t *asynq.TaskInfo) string { return t.ID }},
	{"Type", alignLeft, func(t *asynq.TaskInfo) string { return t.Type }},
	{"Completion Time", alignLeft, func(t *asynq.TaskInfo) string { return formatPastTime(t.CompletedAt) }},
	{"Payload", alignLeft, func(t *asynq.TaskInfo) string { return formatByteSlice(t.Payload) }},
	{"Result", alignLeft, func(t *asynq.TaskInfo) string { return formatByteSlice(t.Result) }},
}

func drawTaskTable(d *ScreenDrawer, state *State) { _ = "STUB: not implemented"; return }

func isNextTaskPageAvailable(s tcell.Screen, state *State) bool {
	_ = "STUB: not implemented"
	return false
}

func drawGroupTable(d *ScreenDrawer, state *State) { _ = "STUB: not implemented"; return }

type number interface {
	int | int64 | float64
}

func min[V number](x, y V) V { _ = "STUB: not implemented"; return *new(V) }

var taskStates = []asynq.TaskState{
	asynq.TaskStateActive,
	asynq.TaskStatePending,
	asynq.TaskStateAggregating,
	asynq.TaskStateScheduled,
	asynq.TaskStateRetry,
	asynq.TaskStateArchived,
	asynq.TaskStateCompleted,
}

func nextTaskState(current asynq.TaskState) asynq.TaskState {
	_ = "STUB: not implemented"
	return *new(asynq.TaskState)
}

func prevTaskState(current asynq.TaskState) asynq.TaskState {
	_ = "STUB: not implemented"
	return *new(asynq.TaskState)
}

func getTaskCount(queue *asynq.QueueInfo, taskState asynq.TaskState) int {
	_ = "STUB: not implemented"
	return 0
}

func drawTaskStateBreakdown(d *ScreenDrawer, style tcell.Style, state *State) {
	_ = "STUB: not implemented"
	return
}

func drawTaskModal(d *ScreenDrawer, state *State) { _ = "STUB: not implemented"; return }

func isPrintable(data []byte) bool { _ = "STUB: not implemented"; return false }

func formatByteSlice(data []byte) string { _ = "STUB: not implemented"; return "" }

type modalRowDrawer struct {
	d        *ScreenDrawer
	width    int
	maxWidth int
}

func (d *modalRowDrawer) Print(s string, style tcell.Style) { _ = "STUB: not implemented"; return }

func withModal(d *ScreenDrawer, rowPrintFns []func(d *modalRowDrawer)) {
	_ = "STUB: not implemented"
	return
}

func adjustWidth(s string, width int) string { _ = "STUB: not implemented"; return "" }

func truncate(s string, max int) string { _ = "STUB: not implemented"; return "" }

func drawDebugInfo(d *ScreenDrawer, state *State) { _ = "STUB: not implemented"; return }

func drawHelp(d *ScreenDrawer) { _ = "STUB: not implemented"; return }
