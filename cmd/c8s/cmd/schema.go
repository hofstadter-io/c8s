package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/c8s/cmd/c8s/flags"
	"github.com/hofstadter-io/c8s/pkg/schema"
)

var schemaLong = `get information about kubernetes resource schemas`

func init() {

	SchemaCmd.Flags().BoolVarP(&(flags.SchemaFlags.Details), "details", "D", false, "print nested definition details for schemas")
	SchemaCmd.Flags().BoolVarP(&(flags.SchemaFlags.Extended), "extended", "E", false, "print extended schemas")
	SchemaCmd.Flags().BoolVarP(&(flags.SchemaFlags.Comments), "comments", "C", false, "print comments with schemas")
}

func SchemaRun(args []string) (err error) {

	// you can safely comment this print out
	// fmt.Println("not implemented")

	err = schema.Run(args, flags.SchemaFlags)

	return err
}

var SchemaCmd = &cobra.Command{

	Use: "schema [...resources]",

	Short: "get information about kubernetes resource schemas",

	Long: schemaLong,

	PreRun: func(cmd *cobra.Command, args []string) {

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = SchemaRun(args)
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

	ohelp := SchemaCmd.HelpFunc()
	ousage := SchemaCmd.UsageFunc()
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

	SchemaCmd.SetHelpFunc(help)
	SchemaCmd.SetUsageFunc(usage)

}
