package cmds

import (
	"github.com/hofstadter-io/hofmod-cli/schema"
)

#ModCmdImports: [
	{Path: "github.com/hofstadter-io/hof/lib/mod", ...},
]

#ModCommand: schema.#Command & {
	Name:  "mod"
	Usage: "mod"
	Aliases: ["m"]
	Short: "dependency management for c8s modules"
	Long: """
  Dependency management for c8s modules based on go mods.

  mod file format:

    module = "<module-path>"

    c8s = "version"

    require (
      <module-path> <version>
      ...
    )

    replace <module path> => <alt path> <version>
      or
    replace <module path> => <local path>
    ...
  """

	OmitRun: true

	Imports: #ModCmdImports

	PersistentPrerun: true
	PersistentPrerunBody: """
    mod.InitLangs()
  """

	Commands: [{
		Name:  "info"
		Usage: "info [module]"
		Short: "print info about modules"
		Long: Short

		Args: [{
			Name: "module"
			Type: "string"
			Help: "name of the module to print info about"
		}]

		Imports: #ModCmdImports

		Body: """
		msg, err := mod.LangInfo("c8s")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(msg)
		"""

	}, {
		Name:  "init"
		Usage: "init <module>"
		Short: "initialize a new module in the current directory"
		Long:  Short

		Args: [{
			Name:     "module"
			Type:     "string"
			Required: true
			Help:     "module full name with path"
		}]

		Imports: #ModCmdImports

		Body: """
		err = mod.Init("c8s", module)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		"""

	}, {
		Name:  "vendor"
		Usage: "vendor"
		Short: "make a vendored copy of dependencies"
		Long:  Short

		Imports: #ModCmdImports

		Body: """
		err = mod.ProcessLangs("vendor", []string{"c8s"})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		"""
	}]

}

