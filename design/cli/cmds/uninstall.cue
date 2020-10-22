package cmds

import (
	"github.com/hofstadter-io/hofmod-cli/schema"
)

#UninstallCommand: schema.#Command & {
	Name:  "uninstall"
	Usage: "uninstall [entrypoints]"
	Short: "uninstall a set of charts"
	Long: Short

	Args: [{
		Name:     "entrypoints"
		Type:     "string"
		Help:     "cuelang entrypoints, defaults to current directory"
	}]
}

