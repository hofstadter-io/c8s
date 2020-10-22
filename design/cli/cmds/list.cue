package cmds

import (
	"github.com/hofstadter-io/hofmod-cli/schema"
)

#ListCommand: schema.#Command & {
	Name:  "list"
	Usage: "list [entrypoints]"
	Short: "list a set of charts"
	Long: Short

	Args: [{
		Name:     "entrypoints"
		Type:     "string"
		Help:     "cuelang entrypoints, defaults to current directory"
	}]
}

