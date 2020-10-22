package cli

import (
	"github.com/hofstadter-io/hofmod-cli/schema"
)

// TODO add: --non-intreactive (-y) ;  --no-color

#CliPflags: [...schema.#Flag] & [
	// Kubernetes related flags
	{
		Name:    "kubeconfig"
		Long:    "kubeconfig"
		Short:   ""
		Type:    "string"
		Default: ""
		Help:    "path to the kubeconfig file"
	},
	{
		Name:    "kubecontext"
		Long:    "kubecontext"
		Short:   ""
		Type:    "string"
		Default: ""
		Help:    "name of the kubeconfig context to use"
	},
	{
		Name:    "namespace"
		Long:    "namespace"
		Short:   "n"
		Type:    "string"
		Default: ""
		Help:    "kubernetes namespace scope for this request"
	},
	
	// these are more cue specific with a dash of hof
	{
		Name:    "package"
		Long:    "package"
		Short:   "p"
		Type:    "string"
		Default: ""
		Help:    "the Cue package context to use"
	},
	{
		Name:    "verbose"
		Long:    "verbose"
		Short:   "v"
		Type:    "int"
		Default: ""
		Help:    "set the verbosity of output"
	},
	{
		Name:    "quiet"
		Long:    "quiet"
		Short:   ""
		Type:    "bool"
		Default: ""
		Help:    "turn off output and assume affirmative defaults at prompts"
	},
]
