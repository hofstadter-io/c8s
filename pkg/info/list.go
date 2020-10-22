package info

import (
	"fmt"

	"github.com/hofstadter-io/c8s/box"
)

func List() {
	fmt.Println("All k8s schemas")
	fns := box.List()
	for _, fn := range fns {
		fmt.Println(" ", fn)
	}
}
