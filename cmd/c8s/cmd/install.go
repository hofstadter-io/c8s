package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var installLong = `install a set of charts`

func InstallRun(entrypoints string) (err error) {

	// you can safely comment this print out
	fmt.Println("not implemented")

	return err
}

var InstallCmd = &cobra.Command{

	Use: "install [entrypoints]",

	Short: "install a set of charts",

	Long: installLong,

	PreRun: func(cmd *cobra.Command, args []string) {

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		var entrypoints string

		if 0 < len(args) {

			entrypoints = args[0]

		}

		err = InstallRun(entrypoints)
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

	ohelp := InstallCmd.HelpFunc()
	ousage := InstallCmd.UsageFunc()
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

	InstallCmd.SetHelpFunc(help)
	InstallCmd.SetUsageFunc(usage)

}
