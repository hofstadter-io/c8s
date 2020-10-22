package cli

import (
	"github.com/hofstadter-io/hofmod-cli/schema"

	"github.com/hofstadter-io/c8s/design/cli/cmds"
)

#Outdir: "./cmd/c8s"

#CLI: schema.#Cli & {
	Name:    "c8s"
	Package: "github.com/hofstadter-io/c8s/cmd/c8s"

	Usage:      "c8s"
	Short:      "The Cuelang for Kubernetes package manager"
	Long:       Short
	CustomHelp: #RootCustomHelp

	OmitRun: true

	PersistentPrerun:     true
	PersistentPostrun: true

	Pflags: #CliPflags

	//Commands: [...schema.#Command]
	Commands: [

		// chart and modules
		cmds.#InitCommand,
		cmds.#ModCommand,

		// manage charts
		cmds.#ListCommand,
		cmds.#InstallCommand,
		cmds.#UpgradeCommand,
		cmds.#RollbackCommand,
		cmds.#UninstallCommand,
		cmds.#HistoryCommand,
		cmds.#StatusCommand,
		cmds.#LintCommand,
		cmds.#TestCommand,

		// repo commands t.b.d.

		// other commands
		cmds.#EnvCommand,

		// hacks down this way
		{
			Hidden: true
			Name:   "hack"
			Usage:  "hack ..."
			Aliases:  ["h", "x"]
			Short:  "development command"
			Long:   Short
		},
	]

	//
	// Addons
	//
	Releases: #CliReleases
	Updates:  true

	EnablePProf: true
}
