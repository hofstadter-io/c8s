package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var envLong = `Environment variables:

| Name                               | Description                                                                       |
|------------------------------------|-----------------------------------------------------------------------------------|
| $C8S_CACHE_HOME                    | set an alternative location for storing cached files.                             |
| $C8S_CONFIG_HOME                   | set an alternative location for storing c8s configuration.                        |
| $C8S_DATA_HOME                     | set an alternative location for storing c8s data.                                 |
| $C8S_DRIVER                        | set the backend storage driver. Values are: configmap, secret, memory, postgres   |
| $C8S_DRIVER_SQL_CONNECTION_STRING  | set the connection string the SQL storage driver should use.                      |
| $KUBECONFIG                        | set an alternative Kubernetes configuration file (default "~/.kube/config")       |

c8s stores cache, configuration, and data based on the following configuration order:

- If a c8s_*_HOME environment variable is set, it will be used
- Otherwise, on systems supporting the XDG base directory specification, the XDG variables will be used
- When no other location is set a default location will be used based on the operating system

By default, the default directories depend on the Operating System. The defaults are listed below:

| Operating System | Cache Path                | Configuration Path             | Data Path               |
|------------------|---------------------------|--------------------------------|-------------------------|
| Linux            | $HOME/.cache/c8s          | $HOME/.config/c8s              | $HOME/.local/share/c8s  |
| macOS            | $HOME/Library/Caches/c8s  | $HOME/Library/Preferences/c8s  | $HOME/Library/c8s       |
| Windows          | %TEMP%\c8s                | %APPDATA%\c8s                  | %APPDATA%\c8s           |

Flags will override the respective environment variables`

func EnvRun(args []string) (err error) {

	// you can safely comment this print out
	fmt.Println("not implemented")

	return err
}

var EnvCmd = &cobra.Command{

	Use: "env",

	Aliases: []string{
		"environment",
	},

	Short: "print information about c8s environment",

	Long: envLong,

	PreRun: func(cmd *cobra.Command, args []string) {

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = EnvRun(args)
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

	ohelp := EnvCmd.HelpFunc()
	ousage := EnvCmd.UsageFunc()
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

	EnvCmd.SetHelpFunc(help)
	EnvCmd.SetUsageFunc(usage)

}
