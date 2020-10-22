package cmds

import (
	"github.com/hofstadter-io/hofmod-cli/schema"
)

#EnvLong: """
Environment variables:

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
| Windows          | %TEMP%\\c8s                | %APPDATA%\\c8s                  | %APPDATA%\\c8s           |

Flags will override the respective environment variables
"""

#EnvCommand: schema.#Command & {
	Name:  "env"
	Usage: "env"
	Aliases: ["environment"]
	Short: "print information about c8s environment"
	Long:  #EnvLong

}
