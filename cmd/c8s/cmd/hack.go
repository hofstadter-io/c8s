package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var hackLong = `development command`

func HackRun(args []string) (err error) {

	// you can safely comment this print out
	fmt.Println("not implemented")

	return err
}

var HackCmd = &cobra.Command{

	Use: "hack ...",

	Hidden: true,

	Aliases: []string{
		"h",
		"x",
	},

	Short: "development command",

	Long: hackLong,

	PreRun: func(cmd *cobra.Command, args []string) {

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = HackRun(args)
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

	ohelp := HackCmd.HelpFunc()
	ousage := HackCmd.UsageFunc()
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

	HackCmd.SetHelpFunc(help)
	HackCmd.SetUsageFunc(usage)

}
