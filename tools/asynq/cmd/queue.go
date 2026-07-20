package cmd

import (
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/hibiken/asynq"
	"github.com/spf13/cobra"
)

const separator = "================================================="

func init() {
	rootCmd.AddCommand(queueCmd)
	queueCmd.AddCommand(queueListCmd)
	queueCmd.AddCommand(queueInspectCmd)
	queueCmd.AddCommand(queueHistoryCmd)
	queueHistoryCmd.Flags().IntP("days", "x", 10, "show data from last x days")

	queueCmd.AddCommand(queuePauseCmd)
	queueCmd.AddCommand(queueUnpauseCmd)
	queueCmd.AddCommand(queueRemoveCmd)
	queueRemoveCmd.Flags().BoolP("force", "f", false, "remove the queue regardless of its size")
}

var queueCmd = &cobra.Command{
	Use:   "queue <command> [flags]",
	Short: "Manage queues",
	Example: heredoc.Doc(`
	  $ asynq queue ls
	  $ asynq queue inspect myqueue
	  $ asynq queue pause myqueue`),
}

var queueListCmd = &cobra.Command{
	Use:     "list",
	Short:   "List queues",
	Aliases: []string{"ls"},
	RunE:    queueList,
}

var queueInspectCmd = &cobra.Command{
	Use:   "inspect <queue> [<queue>...]",
	Short: "Display detailed information on one or more queues",
	Args:  cobra.MinimumNArgs(1),
	RunE:  queueInspect,
	Example: heredoc.Doc(`
		$ asynq queue inspect myqueue
		$ asynq queue inspect queue1 queue2 queue3`),
}

var queueHistoryCmd = &cobra.Command{
	Use:   "history <queue> [<queue>...]",
	Short: "Display historical aggregate data from one or more queues",
	Args:  cobra.MinimumNArgs(1),
	RunE:  queueHistory,
	Example: heredoc.Doc(`
		$ asynq queue history myqueue
		$ asynq queue history queue1 queue2 queue3
		$ asynq queue history myqueue --days=90`),
}

var queuePauseCmd = &cobra.Command{
	Use:   "pause <queue> [<queue>...]",
	Short: "Pause one or more queues",
	Args:  cobra.MinimumNArgs(1),
	RunE:  queuePause,
	Example: heredoc.Doc(`
		$ asynq queue pause myqueue
		$ asynq queue pause queue1 queue2 queue3`),
}

var queueUnpauseCmd = &cobra.Command{
	Use:     "resume <queue> [<queue>...]",
	Short:   "Resume (unpause) one or more queues",
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"unpause"},
	RunE:    queueUnpause,
	Example: heredoc.Doc(`
		$ asynq queue resume myqueue
		$ asynq queue resume queue1 queue2 queue3`),
}

var queueRemoveCmd = &cobra.Command{
	Use:     "remove <queue> [<queue>...]",
	Short:   "Remove one or more queues",
	Aliases: []string{"rm", "delete"},
	Args:    cobra.MinimumNArgs(1),
	RunE:    queueRemove,
	Example: heredoc.Doc(`
		$ asynq queue rm myqueue
		$ asynq queue rm queue1 queue2 queue3
		$ asynq queue rm myqueue --force`),
}

func queueList(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func queueInspect(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func printQueueInfo(info *asynq.QueueInfo) { _ = "STUB: not implemented"; return }

func queueHistory(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func printDailyStats(stats []*asynq.DailyStats) { _ = "STUB: not implemented"; return }

func queuePause(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func queueUnpause(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func queueRemove(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }
