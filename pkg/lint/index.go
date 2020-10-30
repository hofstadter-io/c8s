package schema

import (
	"fmt"
	"os"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/errors"
	"cuelang.org/go/cue/load"

	"github.com/hofstadter-io/c8s/pkg/common"
)

var r cue.Runtime

func Run(args []string) error {

	err := common.Load()
	if err != nil {
		return err
	}

	hasErr := false

	bis := load.Instances(args, &common.Config)
	for _, bi := range bis {
		if bi.Err != nil {
			fmt.Println("Load Err:", bi.Err)
			hasErr = true
			continue
		}

		I, err := r.Build(bi)
		if err != nil {
			fmt.Println("Build Err:", err)
			hasErr = true
			continue
		}

		V := I.Value()
		val := V

		// TODO, use flags to subselect like Cue
		//path := "deploy"
		//val := V.LookupPath(cue.ParsePath(path))

		err = val.Validate(
			cue.Concrete(true),
			cue.Final(),
			cue.All(),
		)
		if err != nil {
			errors.Print(os.Stderr, err, nil)
			hasErr = true
			continue
		}

	}

	if !hasErr {
		fmt.Println("OK")
	}
	return nil
}

