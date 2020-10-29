package schema

import (
	"fmt"
	"strings"

	"cuelang.org/go/cue"
	"github.com/mattn/go-zglob"

	"github.com/hofstadter-io/c8s/box"
	"github.com/hofstadter-io/c8s/cmd/c8s/flags"
)

func Run(args []string, iFlags flags.SchemaFlagpole) error {
	fmt.Println("flags", iFlags)
	// list code paths
	if len(args) == 0 {
		return List(args)
	} else if len(args) > 0 && args[0] == "list" {
		return List(args[1:])

	// details & extended code paths
	} else {
		if iFlags.Extended {
			return Extended(args, iFlags)
		}
		return Details(args, iFlags)
	}
}

// List prints the k8s resource schema(s) filenames
func List(args []string) error {
	// get all the files
	fns := box.List()

	// print all
	if len(args) == 0 {
		for _, fn := range fns {
			fmt.Println(" ", fn)
		}
		return nil
	}

	// print by glob
	for _, fn := range fns {
		for _, arg := range args {
			if !strings.Contains(arg, "*") {
				arg = "**/" + arg + "/**"
			}

			match, err := zglob.Match(arg, fn)
			if err != nil {
				return err
			}

			if match {
				fmt.Println(" ", fn)
			}
		}
	}

	return nil
}

// Details prints the Cue value for matching k8s resource schemas
func Details(args []string, iFlags flags.SchemaFlagpole) error {
	// TODO, lookup from a map

  return nil
}

func Extended(args []string, iFlags flags.SchemaFlagpole) error {
	// get all the files
	fns := box.List()

	// print by glob
	for _, fn := range fns {
		for _, arg := range args {
			if !strings.Contains(arg, "*") {
				arg = "**/" + arg + "/**"
			}

			match, err := zglob.Match(arg, fn)
			if err != nil {
				return err
			}

			if match {
				err = SchemaDetail(fn, iFlags)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

var r cue.Runtime

func SchemaDetail(fn string, iFlags flags.SchemaFlagpole) error {
	b := box.Get(fn)
	fmt.Printf("\n%s\n===========================================\n%s\n\n", fn, string(b))

	return nil
}

/*
Flags for:
- just defs / names
- with(out) comments
- recurse defs
- details (what is printing now)

Pick out the top resources typically used
- make mapping
- easy to use
*/
