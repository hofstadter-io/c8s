package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/c8s/pkg/lint"
)

var lintLong = `lint a set of charts`

func LintRun(entrypoints []string) (err error) {

	// you can safely comment this print out
	// fmt.Println("not implemented")

	err = schema.Run(entrypoints)

	return err
}

var LintCmd = &cobra.Command{

	Use: "lint [entrypoints]",

	Short: "lint a set of charts",

	Long: lintLong,

	PreRun: func(cmd *cobra.Command, args []string) {

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		var entrypoints []string

		if 0 < len(args) {

		}

		err = LintRun(entrypoints)
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

	ohelp := LintCmd.HelpFunc()
	ousage := LintCmd.UsageFunc()
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

	LintCmd.SetHelpFunc(help)
	LintCmd.SetUsageFunc(usage)

}