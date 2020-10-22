package cli

import (
	"github.com/hofstadter-io/c8s/design/cli/cmds"
)

#RootCustomHelp: """
c8s - Kubernetes resource and package manager

  Learn more at https://docs.hofstadter.io/c8s

Usage:
  c8s [flags] [command] [args]

Initialize a new chart and work with modules:
  \(cmds.#InitCommand.Help)
  \(cmds.#ModCommand.Help)

Manage c8s charts and workspaces:
  \(cmds.#ListCommand.Help)
  \(cmds.#InstallCommand.Help)
  \(cmds.#UpgradeCommand.Help)
  \(cmds.#RollbackCommand.Help)
  \(cmds.#UninstallCommand.Help)
  \(cmds.#HistoryCommand.Help)
  \(cmds.#StatusCommand.Help)
  \(cmds.#LintCommand.Help)
  \(cmds.#TestCommand.Help)

Get information about kubernetes resource schemas:
  \(cmds.#InfoCommand.Help)

	(repo commands t.b.d.)

You can also chat with us on https://gitter.im/hofstadter-io

Additional commands:
  env                   print information about c8s environment
  help                  help about any command
  update                check for new versions and run self-updates
  version               print detailed version information
  completion            generate completion helpers for your terminal

Flags:
<<flag-usage>>
Use "c8s [command] --help / -h" for more information about a command.

"""
