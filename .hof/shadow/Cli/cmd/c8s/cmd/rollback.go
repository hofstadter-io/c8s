package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rollbackLong = `rollback a set of charts`

func RollbackRun(entrypoints string) (err error) {

	// you can safely comment this print out
	fmt.Println("not implemented")

	return err
}

var RollbackCmd = &cobra.Command{

	Use: "rollback [entrypoints]",

	Short: "rollback a set of charts",

	Long: rollbackLong,

	PreRun: func(cmd *cobra.Command, args []string) {

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		var entrypoints string

		if 0 < len(args) {

			entrypoints = args[0]

		}

		err = RollbackRun(entrypoints)
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

	ohelp := RollbackCmd.HelpFunc()
	ousage := RollbackCmd.UsageFunc()
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

	RollbackCmd.SetHelpFunc(help)
	RollbackCmd.SetUsageFunc(usage)

}
