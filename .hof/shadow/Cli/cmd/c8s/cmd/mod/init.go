package cmdmod

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/lib/mod"
)

var initLong = `initialize a new module in the current directory`

func InitRun(module string) (err error) {

	err = mod.Init("c8s", module)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return err
}

var InitCmd = &cobra.Command{

	Use: "init <module>",

	Short: "initialize a new module in the current directory",

	Long: initLong,

	PreRun: func(cmd *cobra.Command, args []string) {

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		if 0 >= len(args) {
			fmt.Println("missing required argument: 'module'")
			cmd.Usage()
			os.Exit(1)
		}

		var module string

		if 0 < len(args) {

			module = args[0]

		}

		err = InitRun(module)
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

	ohelp := InitCmd.HelpFunc()
	ousage := InitCmd.UsageFunc()
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

	InitCmd.SetHelpFunc(help)
	InitCmd.SetUsageFunc(usage)

}
