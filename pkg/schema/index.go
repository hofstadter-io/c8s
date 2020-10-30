package schema

import (
	"bytes"
	"fmt"
	"sort"
	"text/template"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/format"
	"cuelang.org/go/cue/load"

	"github.com/hofstadter-io/c8s/cmd/c8s/flags"
	"github.com/hofstadter-io/c8s/pkg/common"
)

var r cue.Runtime

var tplString string = `
import (
{{ range .Values }}  {{.Key}} "{{.Dir}}"
{{ end }}
)

{{ range .Values }}{{.Val}}: {{.Key}}.{{.Path}}
{{ end }}

`

func Run(args []string, iFlags flags.SchemaFlagpole) error {
	if len(args) == 0 {
		printResources()
		return nil
	}

	err := common.Load()
	if err != nil {
		return err
	}

	ms := []mapping{}
	for _, arg := range args {
		m, ok := mappings[arg]
		if !ok {
			return fmt.Errorf("Unkown resource %q", arg)
		}
		ms = append(ms, m)
	}
	data := map[string]interface{}{}
	data["Values"] = ms

	tmpl, err := template.New("test").Parse(tplString)
	if err != nil {
		return err
	}
	var b bytes.Buffer
	err = tmpl.Execute(&b, data)
	if err != nil {
		fmt.Println("err:", err)
		return err
	}

	fmt.Println("===============")
	fmt.Println(b.String())
	fmt.Println("===============")

	common.Config.Stdin = &b

	bis := load.Instances([]string{"-"}, &common.Config)
	for _, bi := range bis {
		if bi.Err != nil {
			fmt.Println("Load Err:", bi.Err)
			continue
		}

		I, err := r.Build(bi)
		if err != nil {
			fmt.Println("Build Err:", err)
			continue
		}

		V := I.Value()
		val := V

		//path := "deploy"
		//val := V.LookupPath(cue.ParsePath(path))


		syn := val.Syntax(
			cue.Final(),
			cue.Raw(),
			cue.Schema(),
			cue.Attributes(false),
			cue.Concrete(false),
			cue.Definitions(true),
			cue.Docs(true),
			cue.Hidden(true),
			cue.Optional(true),
			cue.ResolveReferences(true),
		)
		out, err := format.Node(
			syn,
			format.TabIndent(false),
			format.UseSpaces(2),
			format.Simplify(),
		)
		if err != nil {
			fmt.Println("Print Err:", err)
			continue
		}

		fmt.Println(string(out))
	}

	return nil
}

func printResources() {
		fmt.Println("Known resources:\n--------------")
		rm := map[string][]string{}

		for k, v := range mappings {
			curr := rm[v.Val]
			curr = append(curr, k)
			rm[v.Val] = curr
		}

		// collect and sort keys
		rs := []string{}
		for r, _ := range rm {
			rs = append(rs, r)
		}
		sort.Strings(rs)

		for _, r := range rs {
			rs := rm[r]
			if len(rs) > 1 {
				sort.Strings(rs)
				fmt.Printf("%s: %v\n", r, rs)
			} else {
				fmt.Println(r)
			}
		}
}
