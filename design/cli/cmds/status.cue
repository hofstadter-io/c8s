package cmds

import (
	"github.com/hofstadter-io/hofmod-cli/schema"
)

#StatusCommand: schema.#Command & {
	Name:  "status"
	Usage: "status [entrypoints]"
	Short: "get the status a set of charts"
	Long: Short

	Args: [{
		Name:     "entrypoints"
		Type:     "string"
		Help:     "cuelang entrypoints, defaults to current directory"
	}]
}

