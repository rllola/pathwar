package main

import "os"

func Example() {
	os.Args = []string{"-h"}
	flagOutput = os.Stdout
	main()
	// Output:
	// USAGE
	//   pathwar [global flags] <subcommand> [flags] [args...]
	//
	// More info here: https://github.com/pathwar/pathwar/wiki/CLI
	//
	// SUBCOMMANDS
	//   api      manage the Pathwar API
	//   compose  manage a challenge
	//   agent    manage an agent node (multiple challenges)
	//   sso      manage SSO tokens
	//   misc     misc contains advanced commands
	//   version  show version
	//
	// FLAGS
	//   -debug false  debug mode
}
