package cmds

import (
	"github.com/hofstadter-io/hofmod-cli/schema"
)

#RollbackCommand: schema.#Command & {
	Name:  "rollback"
	Usage: "rollback [entrypoints]"
	Short: "rollback a set of charts"
	Long: Short

	Args: [{
		Name:     "entrypoints"
		Type:     "string"
		Help:     "cuelang entrypoints, defaults to current directory"
	}]
}

