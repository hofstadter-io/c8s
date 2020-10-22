package cmds

import (
	"github.com/hofstadter-io/hofmod-cli/schema"
)

#TestCommand: schema.#Command & {
	Name:  "test"
	Usage: "test [entrypoints]"
	Short: "test a set of charts"
	Long: Short

	Args: [{
		Name:     "entrypoints"
		Type:     "string"
		Help:     "cuelang entrypoints, defaults to current directory"
	}]
}

