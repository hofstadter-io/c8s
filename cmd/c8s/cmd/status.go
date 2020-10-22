package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var statusLong = `get the status a set of charts`

func StatusRun(entrypoints string) (err error) {

	// you can safely comment this print out
	fmt.Println("not implemented")

	return err
}

var StatusCmd = &cobra.Command{

	Use: "status [entrypoints]",

	Short: "get the status a set of charts",

	Long: statusLong,

	PreRun: func(cmd *cobra.Command, args []string) {

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		var entrypoints string

		if 0 < len(args) {

			entrypoints = args[0]

		}

		err = StatusRun(entrypoints)
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

	ohelp := StatusCmd.HelpFunc()
	ousage := StatusCmd.UsageFunc()
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

	StatusCmd.SetHelpFunc(help)
	StatusCmd.SetUsageFunc(usage)

}
