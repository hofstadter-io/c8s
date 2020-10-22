package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/lib/mod"

	"github.com/hofstadter-io/c8s/cmd/c8s/cmd/mod"
)

var modLong = `Dependency management for c8s modules based on go mods.

mod file format:

  module = "<module-path>"

  c8s = "version"

  require (
    <module-path> <version>
    ...
  )

  replace <module path> => <alt path> <version>
    or
  replace <module path> => <local path>
  ...`

func ModPersistentPreRun(args []string) (err error) {

	mod.InitLangs()

	return err
}

var ModCmd = &cobra.Command{

	Use: "mod",

	Aliases: []string{
		"m",
	},

	Short: "dependency management for c8s modules",

	Long: modLong,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = ModPersistentPreRun(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},

	PreRun: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	extra := func(cmd *cobra.Command) bool {

		return false
	}

	ohelp := ModCmd.HelpFunc()
	ousage := ModCmd.UsageFunc()
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

	ModCmd.SetHelpFunc(help)
	ModCmd.SetUsageFunc(usage)

	ModCmd.AddCommand(cmdmod.InfoCmd)
	ModCmd.AddCommand(cmdmod.InitCmd)
	ModCmd.AddCommand(cmdmod.VendorCmd)

}
