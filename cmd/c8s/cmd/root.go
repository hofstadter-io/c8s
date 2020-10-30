package cmd

import (
	"fmt"
	"os"

	"log"
	"runtime/pprof"

	"strings"

	"github.com/hofstadter-io/hof/script/runtime"
	"github.com/spf13/cobra"

	"github.com/hofstadter-io/c8s/cmd/c8s/flags"
)

var c8sLong = `The Cuelang for Kubernetes package manager`

func init() {

	RootCmd.PersistentFlags().StringVarP(&(flags.RootPflags.Kubeconfig), "kubeconfig", "", "", "path to the kubeconfig file")
	RootCmd.PersistentFlags().StringVarP(&(flags.RootPflags.Kubecontext), "kubecontext", "", "", "name of the kubeconfig context to use")
	RootCmd.PersistentFlags().StringVarP(&(flags.RootPflags.Namespace), "namespace", "n", "", "kubernetes namespace scope for this request")
	RootCmd.PersistentFlags().StringVarP(&(flags.RootPflags.Package), "package", "p", "", "the Cue package context to use")
	RootCmd.PersistentFlags().IntVarP(&(flags.RootPflags.Verbose), "verbose", "v", 0, "set the verbosity of output")
	RootCmd.PersistentFlags().BoolVarP(&(flags.RootPflags.Quiet), "quiet", "", false, "turn off output and assume affirmative defaults at prompts")
}

func RootPersistentPreRun(args []string) (err error) {

	return err
}

func RootPersistentPostRun(args []string) (err error) {

	WaitPrintUpdateAvailable()

	return err
}

var RootCmd = &cobra.Command{

	Use: "c8s",

	Short: "The Cuelang for Kubernetes package manager",

	Long: c8sLong,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = RootPersistentPreRun(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},

	PreRun: func(cmd *cobra.Command, args []string) {

	},

	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = RootPersistentPostRun(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func RootInit() {
	extra := func(cmd *cobra.Command) bool {

		return false
	}

	help := func(cmd *cobra.Command, args []string) {
		if extra(cmd) {
			return
		}
		fu := RootCmd.Flags().FlagUsages()
		rh := strings.Replace(RootCustomHelp, "<<flag-usage>>", fu, 1)
		fmt.Println(rh)
	}
	usage := func(cmd *cobra.Command) error {
		if extra(cmd) {
			return nil
		}
		fu := RootCmd.Flags().FlagUsages()
		rh := strings.Replace(RootCustomHelp, "<<flag-usage>>", fu, 1)
		fmt.Println(rh)
		return fmt.Errorf("unknown c8s command")
	}

	RootCmd.SetHelpFunc(help)
	RootCmd.SetUsageFunc(usage)

	RootCmd.AddCommand(UpdateCmd)

	RootCmd.AddCommand(VersionCmd)

	RootCmd.AddCommand(CompletionCmd)

	RootCmd.AddCommand(InitCmd)
	RootCmd.AddCommand(ModCmd)
	RootCmd.AddCommand(InfoCmd)
	RootCmd.AddCommand(ListCmd)
	RootCmd.AddCommand(InstallCmd)
	RootCmd.AddCommand(UpgradeCmd)
	RootCmd.AddCommand(RollbackCmd)
	RootCmd.AddCommand(UninstallCmd)
	RootCmd.AddCommand(HistoryCmd)
	RootCmd.AddCommand(StatusCmd)
	RootCmd.AddCommand(LintCmd)
	RootCmd.AddCommand(TestCmd)
	RootCmd.AddCommand(SchemaCmd)
	RootCmd.AddCommand(EnvCmd)
	RootCmd.AddCommand(HackCmd)

}

func RunExit() {
	if err := RunErr(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func RunInt() int {
	if err := RunErr(); err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}

func RunErr() error {

	if fn := os.Getenv("C8S_CPU_PROFILE"); fn != "" {
		f, err := os.Create(fn)
		if err != nil {
			log.Fatal("Could not create file for CPU profile:", err)
		}
		defer f.Close()

		err = pprof.StartCPUProfile(f)
		if err != nil {
			log.Fatal("Could not start CPU profile process:", err)
		}

		defer pprof.StopCPUProfile()
	}

	RootInit()
	return RootCmd.Execute()
}

func CallTS(ts *runtime.Script, args []string) error {
	RootCmd.SetArgs(args)

	err := RootCmd.Execute()
	ts.Check(err)

	return err
}

const RootCustomHelp = `c8s - Kubernetes resource and package manager

  Learn more at https://docs.hofstadter.io/c8s

Usage:
  c8s [flags] [command] [args]

Initialize a new chart and work with modules:
  init                  create an empty module or initialize an existing directory to one
  mod                   dependency management for c8s modules

Manage c8s charts and workspaces:
  list                  list a set of charts
  install               install a set of charts
  upgrade               upgrade a set of charts
  rollback              rollback a set of charts
  uninstall             uninstall a set of charts
  history               get the history for charts
  status                get the status a set of charts
  lint                  lint a set of charts
  test                  test a set of charts

Get information about kubernetes resource schemas:
  info                  get information about kubernetes resource schemas

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
`
