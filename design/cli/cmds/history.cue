package cmds

import (
	"github.com/hofstadter-io/hofmod-cli/schema"
)

#HistoryCommand: schema.#Command & {
	Name:  "history"
	Usage: "history [entrypoints]"
	Short: "get the history for charts"
	Long: Short

	Args: [{
		Name:     "entrypoints"
		Type:     "string"
		Help:     "cuelang entrypoints, defaults to current directory"
	}]
}

