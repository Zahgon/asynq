package cmd

import (
	"time"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/hibiken/asynq"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(taskCmd)
	taskCmd.AddCommand(taskListCmd)
	taskListCmd.Flags().StringP("queue", "q", "", "queue to inspect (required)")
	taskListCmd.Flags().StringP("state", "s", "", "state of the tasks; one of { active | pending | aggregating | scheduled | retry | archived | completed } (required)")
	taskListCmd.Flags().Int("page", 1, "page number")
	taskListCmd.Flags().Int("size", 30, "page size")
	taskListCmd.Flags().StringP("group", "g", "", "group to inspect (required for listing aggregating tasks)")
	taskListCmd.MarkFlagRequired("queue")
	taskListCmd.MarkFlagRequired("state")

	taskCmd.AddCommand(taskCancelCmd)

	taskCmd.AddCommand(taskInspectCmd)
	taskInspectCmd.Flags().StringP("queue", "q", "", "queue to which the task belongs (required)")
	taskInspectCmd.Flags().StringP("id", "i", "", "id of the task (required)")
	taskInspectCmd.MarkFlagRequired("queue")
	taskInspectCmd.MarkFlagRequired("id")

	taskCmd.AddCommand(taskArchiveCmd)
	taskArchiveCmd.Flags().StringP("queue", "q", "", "queue to which the task belongs (required)")
	taskArchiveCmd.Flags().StringP("id", "i", "", "id of the task (required)")
	taskArchiveCmd.MarkFlagRequired("queue")
	taskArchiveCmd.MarkFlagRequired("id")

	taskCmd.AddCommand(taskDeleteCmd)
	taskDeleteCmd.Flags().StringP("queue", "q", "", "queue to which the task belongs (required)")
	taskDeleteCmd.Flags().StringP("id", "i", "", "id of the task (required)")
	taskDeleteCmd.MarkFlagRequired("queue")
	taskDeleteCmd.MarkFlagRequired("id")

	taskCmd.AddCommand(taskRunCmd)
	taskRunCmd.Flags().StringP("queue", "q", "", "queue to which the task belongs (required)")
	taskRunCmd.Flags().StringP("id", "i", "", "id of the task (required)")
	taskRunCmd.MarkFlagRequired("queue")
	taskRunCmd.MarkFlagRequired("id")

	taskCmd.AddCommand(taskEnqueueCmd)
	taskEnqueueCmd.Flags().StringP("type_name", "t", "", "type name to enqueue the task as (required)")
	taskEnqueueCmd.Flags().StringP("payload", "l", "", "payload to enqueue (required)")

	taskEnqueueCmd.Flags().Int("retry", 0, "maximum retries")
	taskEnqueueCmd.Flags().String("queue", "", "queue to enqueue the task to")
	taskEnqueueCmd.Flags().String("id", "", "id to enqueue the task as")
	taskEnqueueCmd.Flags().String("timeout", "", "timeout for the task (how long it can run); must be parseable as a time.Duration")
	taskEnqueueCmd.Flags().String("deadline", "", "deadline for the task; must be in RFC3339 format")
	taskEnqueueCmd.Flags().String("unique", "", "unique period for the task (duration within which it is guaranteed to be unique); must be parseable as a time.Duration")
	taskEnqueueCmd.Flags().String("process_at", "", "process at time for the task; must be in RFC3339 format")
	taskEnqueueCmd.Flags().String("process_in", "", "process in window for the task; must be parseable as a time.Duration")
	taskEnqueueCmd.Flags().String("retention", "", "retention window for the task; must be parseable as a time.Duration")
	taskEnqueueCmd.Flags().String("group", "", "group for the task")
	taskEnqueueCmd.MarkFlagRequired("type_name")
	taskEnqueueCmd.MarkFlagRequired("payload")

	taskCmd.AddCommand(taskArchiveAllCmd)
	taskArchiveAllCmd.Flags().StringP("queue", "q", "", "queue to which the tasks belong (required)")
	taskArchiveAllCmd.Flags().StringP("state", "s", "", "state of the tasks; one of { pending | aggregating | scheduled | retry } (required)")
	taskArchiveAllCmd.MarkFlagRequired("queue")
	taskArchiveAllCmd.MarkFlagRequired("state")
	taskArchiveAllCmd.Flags().StringP("group", "g", "", "group to which the tasks belong (required for archiving aggregating tasks)")

	taskCmd.AddCommand(taskDeleteAllCmd)
	taskDeleteAllCmd.Flags().StringP("queue", "q", "", "queue to which the tasks belong (required)")
	taskDeleteAllCmd.Flags().StringP("state", "s", "", "state of the tasks; one of { pending | aggregating | scheduled | retry | archived | completed } (required)")
	taskDeleteAllCmd.MarkFlagRequired("queue")
	taskDeleteAllCmd.MarkFlagRequired("state")
	taskDeleteAllCmd.Flags().StringP("group", "g", "", "group to which the tasks belong (required for deleting aggregating tasks)")

	taskCmd.AddCommand(taskRunAllCmd)
	taskRunAllCmd.Flags().StringP("queue", "q", "", "queue to which the tasks belong (required)")
	taskRunAllCmd.Flags().StringP("state", "s", "", "state of the tasks; one of { scheduled | retry | archived } (required)")
	taskRunAllCmd.MarkFlagRequired("queue")
	taskRunAllCmd.MarkFlagRequired("state")
	taskRunAllCmd.Flags().StringP("group", "g", "", "group to which the tasks belong (required for running aggregating tasks)")
}

var taskCmd = &cobra.Command{
	Use:   "task <command> [flags]",
	Short: "Manage tasks",
	Example: heredoc.Doc(`
		$ asynq task list --queue=myqueue --state=scheduled
		$ asynq task inspect --queue=myqueue --id=7837f142-6337-4217-9276-8f27281b67d1
		$ asynq task delete --queue=myqueue --id=7837f142-6337-4217-9276-8f27281b67d1
		$ asynq task deleteall --queue=myqueue --state=archived`),
}

var taskListCmd = &cobra.Command{
	Use:     "list --queue=<queue> --state=<state> [flags]",
	Aliases: []string{"ls"},
	Short:   "List tasks",
	Long: heredoc.Doc(`
	List tasks of the given state from the specified queue.

	The --queue and --state flags are required.

	Note: For aggregating tasks, additional --group flag is required.

	List opeartion paginates the result set. By default, the command fetches the first 30 tasks.
	Use --page and --size flags to specify the page number and size.`),
	Example: heredoc.Doc(`
		$ asynq task list --queue=myqueue --state=pending
		$ asynq task list --queue=myqueue --state=aggregating --group=mygroup
		$ asynq task list --queue=myqueue --state=scheduled --page=2`),
	RunE: taskList,
}

var taskInspectCmd = &cobra.Command{
	Use:   "inspect --queue=<queue> --id=<task_id>",
	Short: "Display detailed information on the specified task",
	Args:  cobra.NoArgs,
	RunE:  taskInspect,
	Example: heredoc.Doc(`
		$ asynq task inspect --queue=myqueue --id=f1720682-f5a6-4db1-8953-4f48ae541d0f`),
}

var taskCancelCmd = &cobra.Command{
	Use:   "cancel <task_id> [<task_id>...]",
	Short: "Cancel one or more active tasks",
	Args:  cobra.MinimumNArgs(1),
	RunE:  taskCancel,
	Example: heredoc.Doc(`
		$ asynq task cancel f1720682-f5a6-4db1-8953-4f48ae541d0f`),
}

var taskArchiveCmd = &cobra.Command{
	Use:   "archive --queue=<queue> --id=<task_id>",
	Short: "Archive a task with the given id",
	Args:  cobra.NoArgs,
	RunE:  taskArchive,
	Example: heredoc.Doc(`
		$ asynq task archive --queue=myqueue --id=f1720682-f5a6-4db1-8953-4f48ae541d0f`),
}

var taskDeleteCmd = &cobra.Command{
	Use:     "delete --queue=<queue> --id=<task_id>",
	Aliases: []string{"remove", "rm"},
	Short:   "Delete a task with the given id",
	Args:    cobra.NoArgs,
	RunE:    taskDelete,
	Example: heredoc.Doc(`
		$ asynq task delete --queue=myqueue --id=f1720682-f5a6-4db1-8953-4f48ae541d0f`),
}

var taskRunCmd = &cobra.Command{
	Use:   "run --queue=<queue> --id=<task_id>",
	Short: "Run a task with the given id",
	Args:  cobra.NoArgs,
	RunE:  taskRun,
	Example: heredoc.Doc(`
		$ asynq task run --queue=myqueue --id=f1720682-f5a6-4db1-8953-4f48ae541d0f`),
}

var taskEnqueueCmd = &cobra.Command{
	Use:   "enqueue --type_name=footype --payload=barpayload",
	Short: "Enqueue a task",
	Args:  cobra.NoArgs,
	RunE:  taskEnqueue,
	Example: heredoc.Doc(`
		$ asynq task enqueue -t footype -l barpayload
		$ asynq task enqueue -t footask -l barpayload --retry 3 --id f1720682-f5a6-4db1-8953-4f48ae541d0f --queue bazqueue --timeout 100s --deadline 2024-12-14T01:23:45Z --unique 100s --process_at 2024-12-14T01:22:05Z --process_in 100s --retention 5h --group baygroup`),
}

var taskArchiveAllCmd = &cobra.Command{
	Use:   "archiveall --queue=<queue> --state=<state>",
	Short: "Archive all tasks in the given state",
	Args:  cobra.NoArgs,
	RunE:  taskArchiveAll,
	Example: heredoc.Doc(`
		$ asynq task archiveall --queue=myqueue --state=retry
		$ asynq task archiveall --queue=myqueue --state=aggregating --group=mygroup`),
}

var taskDeleteAllCmd = &cobra.Command{
	Use:   "deleteall --queue=<queue> --state=<state>",
	Short: "Delete all tasks in the given state",
	Args:  cobra.NoArgs,
	RunE:  taskDeleteAll,
	Example: heredoc.Doc(`
		$ asynq task deleteall --queue=myqueue --state=archived
		$ asynq task deleteall --queue=myqueue --state=aggregating --group=mygroup`),
}

var taskRunAllCmd = &cobra.Command{
	Use:   "runall --queue=<queue> --state=<state>",
	Short: "Run all tasks in the given state",
	Args:  cobra.NoArgs,
	RunE:  taskRunAll,
	Example: heredoc.Doc(`
		$ asynq task runall --queue=myqueue --state=retry
		$ asynq task runall --queue=myqueue --state=aggregating --group=mygroup`),
}

func taskList(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func listActiveTasks(qname string, pageNum, pageSize int) error {
	_ = "STUB: not implemented"
	return nil
}

func listPendingTasks(qname string, pageNum, pageSize int) error {
	_ = "STUB: not implemented"
	return nil
}

func listScheduledTasks(qname string, pageNum, pageSize int) error {
	_ = "STUB: not implemented"
	return nil
}

func formatProcessAt(processAt time.Time) string { _ = "STUB: not implemented"; return "" }

func listRetryTasks(qname string, pageNum, pageSize int) error {
	_ = "STUB: not implemented"
	return nil
}

func listArchivedTasks(qname string, pageNum, pageSize int) error {
	_ = "STUB: not implemented"
	return nil
}

func listCompletedTasks(qname string, pageNum, pageSize int) error {
	_ = "STUB: not implemented"
	return nil
}

func listAggregatingTasks(qname, group string, pageNum, pageSize int) error {
	_ = "STUB: not implemented"
	return nil
}

func taskCancel(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func taskInspect(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func printTaskInfo(info *asynq.TaskInfo) { _ = "STUB: not implemented"; return }

func formatNextProcessAt(processAt time.Time) string { _ = "STUB: not implemented"; return "" }

func formatPastTime(t time.Time) string { _ = "STUB: not implemented"; return "" }

func taskArchive(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func taskDelete(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func taskRun(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func taskEnqueue(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func taskArchiveAll(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func taskDeleteAll(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func taskRunAll(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }
