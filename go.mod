module github.com/hofstadter-io/c8s

go 1.15

replace cuelang.org/go => ../../cue/cue

replace github.com/hofstadter-io/hof => ../hof

require (
	github.com/hofstadter-io/hof v0.3.8
	github.com/parnurzeal/gorequest v0.2.16
	github.com/spf13/cobra v1.1.1
	moul.io/http2curl v1.0.0 // indirect
)
