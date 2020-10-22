package cmds

import (
	"github.com/hofstadter-io/hofmod-cli/schema"
)

#InfoCommand: schema.#Command & {
	Name:  "info"
	Usage: "info [...resources]"
	Short: "get information about kubernetes resource schemas"
	Long: Short
}
