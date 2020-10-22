package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var historyLong = `get the history for charts`

func HistoryRun(entrypoints string) (err error) {

	// you can safely comment this print out
	fmt.Println("not implemented")

	return err
}

var HistoryCmd = &cobra.Command{

	Use: "history [entrypoints]",

	Short: "get the history for charts",

	Long: historyLong,

	PreRun: func(cmd *cobra.Command, args []string) {

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		var entrypoints string

		if 0 < len(args) {

			entrypoints = args[0]

		}

		err = HistoryRun(entrypoints)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	extra := func(cmd *cobra.Command) bool {

		return false
	}

	ohelp := HistoryCmd.HelpFunc()
	ousage := HistoryCmd.UsageFunc()
	help := func(cmd *cobra.Command, args []string) {
		if extra(cmd) {
			return
		}
		ohelp(cmd, args)
	}
	usage := func(cmd *cobra.Command) error {
		if extra(cmd) {
			return nil
		}
		return ousage(cmd)
	}

	HistoryCmd.SetHelpFunc(help)
	HistoryCmd.SetUsageFunc(usage)

}
