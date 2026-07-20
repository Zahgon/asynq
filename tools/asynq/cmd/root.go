package cmd

import (
	"crypto/tls"
	"fmt"
	"io"
	"time"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/hibiken/asynq"
	"github.com/hibiken/asynq/internal/base"
	"github.com/hibiken/asynq/internal/rdb"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

var (
	uri      string
	db       int
	password string
	username string

	useRedisCluster bool
	clusterAddrs    string
	tlsServerName   string
	insecure        bool
	useTLS          bool
)

var rootCmd = &cobra.Command{
	Use:     "asynq <command> <subcommand> [flags]",
	Short:   "Asynq CLI",
	Long:    `Command line tool to inspect tasks and queues managed by Asynq`,
	Version: base.Version,

	SilenceUsage:  true,
	SilenceErrors: true,

	Example: heredoc.Doc(`
		$ asynq stats
		$ asynq queue pause myqueue
		$ asynq task list --queue=myqueue --state=archived`),
	Annotations: map[string]string{
		"help:feedback": heredoc.Doc(`
			Open an issue at https://github.com/hibiken/asynq/issues/new/choose`),
	},
}

var versionOutput = fmt.Sprintf("asynq version %s\n", base.Version)

var versionCmd = &cobra.Command{
	Use:    "version",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(versionOutput)
	},
}

func Execute() { _ = "STUB: not implemented"; return }

func isRootCmd(cmd *cobra.Command) bool { _ = "STUB: not implemented"; return false }

type displayLine struct {
	name string
	desc string
	pad  int
}

func (l *displayLine) String() string { _ = "STUB: not implemented"; return "" }

type displayLines []*displayLine

func (dls displayLines) String() string { _ = "STUB: not implemented"; return "" }

func capitalize(s string) string { _ = "STUB: not implemented"; return "" }

func rootHelpFunc(cmd *cobra.Command, args []string) { _ = "STUB: not implemented"; return }

func rootUsageFunc(cmd *cobra.Command) error { _ = "STUB: not implemented"; return nil }

func printSubcommandSuggestions(cmd *cobra.Command, arg string) { _ = "STUB: not implemented"; return }

func adjustPadding(lines ...*displayLine) { _ = "STUB: not implemented"; return }

func rpad(s string, padding int) string { _ = "STUB: not implemented"; return "" }

func lpad(s string, padding int) string { _ = "STUB: not implemented"; return "" }

func indent(text string, space int) string { _ = "STUB: not implemented"; return "" }

func dedent(text string) string { _ = "STUB: not implemented"; return "" }

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.SetHelpFunc(rootHelpFunc)
	rootCmd.SetUsageFunc(rootUsageFunc)

	rootCmd.AddCommand(versionCmd)
	rootCmd.SetVersionTemplate(versionOutput)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file to set flag defaut values (default is $HOME/.asynq.yaml)")
	rootCmd.PersistentFlags().StringVarP(&uri, "uri", "u", "127.0.0.1:6379", "Redis server URI")
	rootCmd.PersistentFlags().IntVarP(&db, "db", "n", 0, "Redis database number (default is 0)")
	rootCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "Password to use when connecting to redis server")
	rootCmd.PersistentFlags().StringVarP(&username, "username", "U", "", "Username to use when connecting to Redis (ACL username)")
	rootCmd.PersistentFlags().BoolVar(&useRedisCluster, "cluster", false, "Connect to redis cluster")
	rootCmd.PersistentFlags().StringVar(&clusterAddrs, "cluster_addrs",
		"127.0.0.1:7000,127.0.0.1:7001,127.0.0.1:7002,127.0.0.1:7003,127.0.0.1:7004,127.0.0.1:7005",
		"List of comma-separated redis server addresses")
	rootCmd.PersistentFlags().BoolVar(&useTLS, "tls", false, "Enable TLS connection")
	rootCmd.PersistentFlags().StringVar(&tlsServerName, "tls_server",
		"", "Server name for TLS validation")
	rootCmd.PersistentFlags().BoolVar(&insecure, "insecure",
		false, "Allow insecure TLS connection by skipping cert validation")

	viper.BindPFlag("uri", rootCmd.PersistentFlags().Lookup("uri"))
	viper.BindPFlag("db", rootCmd.PersistentFlags().Lookup("db"))
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))
	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("cluster", rootCmd.PersistentFlags().Lookup("cluster"))
	viper.BindPFlag("cluster_addrs", rootCmd.PersistentFlags().Lookup("cluster_addrs"))
	viper.BindPFlag("tls", rootCmd.PersistentFlags().Lookup("tls"))
	viper.BindPFlag("tls_server", rootCmd.PersistentFlags().Lookup("tls_server"))
	viper.BindPFlag("insecure", rootCmd.PersistentFlags().Lookup("insecure"))
}

func initConfig() { _ = "STUB: not implemented"; return }

func createRDB() *rdb.RDB { _ = "STUB: not implemented"; return nil }

func createClient() *asynq.Client { _ = "STUB: not implemented"; return nil }

func createInspector() *asynq.Inspector { _ = "STUB: not implemented"; return nil }

func getRedisConnOpt() asynq.RedisConnOpt {
	_ = "STUB: not implemented"
	return *new(asynq.RedisConnOpt)
}

func getTLSConfig() *tls.Config { _ = "STUB: not implemented"; return nil }

func printTable(cols []string, printRows func(w io.Writer, tmpl string)) {
	_ = "STUB: not implemented"
	return
}

func sprintBytes(payload []byte) string { _ = "STUB: not implemented"; return "" }

func isPrintable(data []byte) bool { _ = "STUB: not implemented"; return false }

func getDuration(cmd *cobra.Command, arg string) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

func getTime(cmd *cobra.Command, arg string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}
