package cmd

import (
	"time"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.AddCommand(serverListCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server <command> [flags]",
	Short: "Manage servers",
	Example: heredoc.Doc(`
		$ asynq server list`),
}

var serverListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List servers",
	Long: `Server list (asynq server ls) shows all running worker servers
pulling tasks from the given redis instance.

The command shows the following for each server:
* Host and PID of the process in which the server is running
* Number of active workers out of worker pool
* Queue configuration
* State of the worker server ("active" | "stopped")
* Time the server was started

A "active" server is pulling tasks from queues and processing them.
A "stopped" server is no longer pulling new tasks from queues`,
	RunE: serverList,
}

func serverList(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func formatQueues(qmap map[string]int) string { _ = "STUB: not implemented"; return "" }

func timeAgo(since time.Time) string { _ = "STUB: not implemented"; return "" }
