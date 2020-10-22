package hof

import (
	g_cli "github.com/hofstadter-io/hofmod-cli/gen"
	d_cli "github.com/hofstadter-io/c8s/design/cli"
)

Cli: _ @gen(cli)
Cli: g_cli.#HofGenerator & {
  Outdir: "./"
  Cli: d_cli.#CLI
}
