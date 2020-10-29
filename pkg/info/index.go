package info

import (
	"fmt"

	"github.com/hofstadter-io/c8s/cmd/c8s/flags"
)

func Run(args []string, iFlags flags.InfoFlagpole) error {
	fmt.Println("info", args, iFlags)

  return nil
}

