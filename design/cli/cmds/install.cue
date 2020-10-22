package cmds

import (
	"github.com/hofstadter-io/hofmod-cli/schema"
)

#InstallCommand: schema.#Command & {
	Name:  "install"
	Usage: "install [entrypoints]"
	Short: "install a set of charts"
	Long: Short

	Args: [{
		Name:     "entrypoints"
		Type:     "string"
		Help:     "cuelang entrypoints, defaults to current directory"
	}]
}

