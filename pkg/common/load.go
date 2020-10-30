package common

import (
	"path/filepath"

	"cuelang.org/go/cue/load"

	"github.com/hofstadter-io/c8s/box"
)

var Config load.Config = load.Config {
	Overlay: map[string]load.Source{},
	// Dir: "./.c8s",
}

// Load loads the entire embedded schema
func Load() error {
	fns := box.List()

	for _, fn := range fns {
		content := box.Get(fn)
		ifn := filepath.Join("cue.mod/pkg", fn)
		Config.Overlay[ifn] = load.FromBytes(content)
	}

	Config.Overlay["cue.mod/module.cue"] = load.FromString(`module: "github.com/hofstadter-io/c8s"`)

	return nil
}
