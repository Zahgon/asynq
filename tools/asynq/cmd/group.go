package cmd

import (
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(groupCmd)
	groupCmd.AddCommand(groupListCmd)
	groupListCmd.Flags().StringP("queue", "q", "", "queue to inspect")
	groupListCmd.MarkFlagRequired("queue")
}

var groupCmd = &cobra.Command{
	Use:   "group <command> [flags]",
	Short: "Manage groups",
	Example: heredoc.Doc(`
		$ asynq group list --queue=myqueue`),
}

var groupListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List groups",
	Args:    cobra.NoArgs,
	RunE:    groupLists,
}

func groupLists(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }
