package cmds

import (
	"github.com/hofstadter-io/hofmod-cli/schema"
)

#UpgradeCommand: schema.#Command & {
	Name:  "upgrade"
	Usage: "upgrade [entrypoints]"
	Short: "upgrade a set of charts"
	Long: Short

	Args: [{
		Name:     "entrypoints"
		Type:     "string"
		Help:     "cuelang entrypoints, defaults to current directory"
	}]
}

