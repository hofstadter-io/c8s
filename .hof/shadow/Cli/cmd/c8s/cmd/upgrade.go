package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var upgradeLong = `upgrade a set of charts`

func UpgradeRun(entrypoints string) (err error) {

	// you can safely comment this print out
	fmt.Println("not implemented")

	return err
}

var UpgradeCmd = &cobra.Command{

	Use: "upgrade [entrypoints]",

	Short: "upgrade a set of charts",

	Long: upgradeLong,

	PreRun: func(cmd *cobra.Command, args []string) {

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		var entrypoints string

		if 0 < len(args) {

			entrypoints = args[0]

		}

		err = UpgradeRun(entrypoints)
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

	ohelp := UpgradeCmd.HelpFunc()
	ousage := UpgradeCmd.UsageFunc()
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

	UpgradeCmd.SetHelpFunc(help)
	UpgradeCmd.SetUsageFunc(usage)

}
