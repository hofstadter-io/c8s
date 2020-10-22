package cmd_test

import (
	"testing"

	"github.com/hofstadter-io/hof/lib/yagu"
	"github.com/hofstadter-io/hof/script/runtime"

	"github.com/hofstadter-io/c8s/cmd/c8s/cmd"
)

func TestScriptLintCliTests(t *testing.T) {
	// setup some directories

	dir := "lint"

	workdir := ".workdir/cli/" + dir
	yagu.Mkdir(workdir)

	runtime.Run(t, runtime.Params{
		Setup: func(env *runtime.Env) error {
			// add any environment variables for your tests here

			return nil
		},
		Funcs: map[string]func(ts *runtime.Script, args []string) error{
			"__c8s": cmd.CallTS,
		},
		Dir:         "hls/cli/lint",
		WorkdirRoot: workdir,
	})
}
