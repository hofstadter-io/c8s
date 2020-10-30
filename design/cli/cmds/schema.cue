package cmds

import (
	"github.com/hofstadter-io/hofmod-cli/schema"
)

#SchemaCommand: schema.#Command & {
	Name:  "schema"
	Usage: "schema [...resources]"
	Short: "get information about kubernetes resource schemas"
	Long: Short

	Flags: [
		{
			Name:    "details"
			Long:    "details"
			Short:   "D"
			Type:    "bool"
			Default: "false"
			Help:    "print nested definition details for schemas"
		},
		{
			Name:    "extended"
			Long:    "extended"
			Short:   "E"
			Type:    "bool"
			Default: "false"
			Help:    "print extended schemas"
		},
		{
			Name:    "comments"
			Long:    "comments"
			Short:   "C"
			Type:    "bool"
			Default: "false"
			Help:    "print comments with schemas"
		},
	]
}
