package cmds

import (
	"github.com/hofstadter-io/hofmod-cli/schema"
)

#InitCommand: schema.#Command & {
	Name:  "init"
	Usage: "init <module> [dirname]"
	Short: "create an empty module or initialize an existing directory to one"
	Long:  """
	create an empty module or initialize an existing directory to one

	  module name or path should look like github.com/hofstadter-io/hof
	"""

	Args: [{
		Name:     "module"
		Type:     "string"
		Required: true
		Help:     "module url or path (github.com/hofstadter-io/hof"
	},
	{
		Name:     "name"
		Type:     "string"
		Help:     "module name, defaults to last part of module"
	}]
}

