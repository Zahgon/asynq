package cmd

import (
	"time"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cronCmd)
	cronCmd.AddCommand(cronListCmd)
	cronCmd.AddCommand(cronHistoryCmd)
	cronHistoryCmd.Flags().Int("page", 1, "page number")
	cronHistoryCmd.Flags().Int("size", 30, "page size")
}

var cronCmd = &cobra.Command{
	Use:   "cron <command> [flags]",
	Short: "Manage cron",
	Example: heredoc.Doc(`
		$ asynq cron ls
		$ asynq cron history 7837f142-6337-4217-9276-8f27281b67d1`),
}

var cronListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List cron entries",
	RunE:    cronList,
}

var cronHistoryCmd = &cobra.Command{
	Use:   "history <entry_id> [<entry_id>...]",
	Short: "Show history of each cron tasks",
	Args:  cobra.MinimumNArgs(1),
	RunE:  cronHistory,
	Example: heredoc.Doc(`
		$ asynq cron history 7837f142-6337-4217-9276-8f27281b67d1
		$ asynq cron history 7837f142-6337-4217-9276-8f27281b67d1 bf6a8594-cd03-4968-b36a-8572c5e160dd
		$ asynq cron history 7837f142-6337-4217-9276-8f27281b67d1 --size=100
		$ asynq cron history 7837f142-6337-4217-9276-8f27281b67d1 --page=2`),
}

func cronList(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func nextEnqueue(nextEnqueueAt time.Time) string { _ = "STUB: not implemented"; return "" }

func prevEnqueue(prevEnqueuedAt time.Time) string { _ = "STUB: not implemented"; return "" }

func cronHistory(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }
