package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var infoLong = `get information about kubernetes resource schemas`

func InfoRun(args []string) (err error) {

	// you can safely comment this print out
	fmt.Println("not implemented")

	return err
}

var InfoCmd = &cobra.Command{

	Use: "info [...resources]",

	Short: "get information about kubernetes resource schemas",

	Long: infoLong,

	PreRun: func(cmd *cobra.Command, args []string) {

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = InfoRun(args)
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

	ohelp := InfoCmd.HelpFunc()
	ousage := InfoCmd.UsageFunc()
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

	InfoCmd.SetHelpFunc(help)
	InfoCmd.SetUsageFunc(usage)

}
