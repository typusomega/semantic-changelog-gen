package main

import (
	"github.com/typusomega/semantic-changelog-gen/internal/commands"

	"github.com/jessevdk/go-flags"
)

func main() {
	var opts commands.Opts
	parser := flags.NewParser(&opts, flags.Default)

	_, err := parser.AddCommand("generate", "Build changelog", "", commands.NewGenerateCommand(&opts))
	if err != nil {
		panic(err)
	}

	_, err = parser.Parse()
	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			panic(err)
		}
	}
}
