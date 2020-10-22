package cmdmod

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/lib/mod"
)

var infoLong = `print info about modules`

func InfoRun(module string) (err error) {

	msg, err := mod.LangInfo("c8s")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(msg)

	return err
}

var InfoCmd = &cobra.Command{

	Use: "info [module]",

	Short: "print info about modules",

	Long: infoLong,

	PreRun: func(cmd *cobra.Command, args []string) {

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		var module string

		if 0 < len(args) {

			module = args[0]

		}

		err = InfoRun(module)
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
