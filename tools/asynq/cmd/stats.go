package cmd

import (
	"time"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/hibiken/asynq/internal/rdb"
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "View current state",
	Long: heredoc.Doc(`
	  Stats shows the overview of tasks and queues at that instant.

	  The command shows the following:
	    * Number of tasks in each state
	    * Number of tasks in each queue
	    * Aggregate data for the current day
	    * Basic information about the running redis instance`),
	Args: cobra.NoArgs,
	RunE: stats,
}

var jsonFlag bool

func init() {
	rootCmd.AddCommand(statsCmd)
	statsCmd.Flags().BoolVar(&jsonFlag, "json", false, "Output stats in JSON format.")

}

type AggregateStats struct {
	Active      int       `json:"active"`
	Pending     int       `json:"pending"`
	Aggregating int       `json:"aggregating"`
	Scheduled   int       `json:"scheduled"`
	Retry       int       `json:"retry"`
	Archived    int       `json:"archived"`
	Completed   int       `json:"completed"`
	Processed   int       `json:"processed"`
	Failed      int       `json:"failed"`
	Timestamp   time.Time `json:"timestamp"`
}

type FullStats struct {
	Aggregate  AggregateStats    `json:"aggregate"`
	QueueStats []*rdb.Stats      `json:"queues"`
	RedisInfo  map[string]string `json:"redis"`
}

func stats(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func printStatsByState(s *AggregateStats) { _ = "STUB: not implemented"; return }

func numDigits(n int) int { _ = "STUB: not implemented"; return 0 }

func maxWidthOf(vals ...int) int { _ = "STUB: not implemented"; return 0 }

func maxInt(a, b int) int { _ = "STUB: not implemented"; return 0 }

func printStatsByQueue(stats []*rdb.Stats) { _ = "STUB: not implemented"; return }

func queueTitle(s *rdb.Stats) string { _ = "STUB: not implemented"; return "" }

func printSuccessFailureStats(s *AggregateStats) { _ = "STUB: not implemented"; return }

func printInfo(info map[string]string) { _ = "STUB: not implemented"; return }

func printClusterInfo(info map[string]string) { _ = "STUB: not implemented"; return }

func toInterfaceSlice(strs []string) []interface{} { _ = "STUB: not implemented"; return nil }
