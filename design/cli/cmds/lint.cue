package cmds

import (
	"github.com/hofstadter-io/hofmod-cli/schema"
)

#LintCommand: schema.#Command & {
	Name:  "lint"
	Usage: "lint [entrypoints]"
	Short: "lint a set of charts"
	Long: Short

	Args: [{
		Name:     "entrypoints"
		Type:     "[]string"
		Help:     "cuelang entrypoints, defaults to current directory"
	}]
}

