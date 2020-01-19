package main

import (
	"os"

	"github.com/typusomega/semantic-changelog-gen/internal/commands"

	"github.com/jessevdk/go-flags"
)

var opts = commands.Opts{}

func main() {
	parser := flags.NewParser(&opts, flags.Default)

	_, err := parser.AddCommand("generate", "Generate changelog", "", &commands.GenerateCommand{Opts: &opts})
	if err != nil {
		panic(err)
	}

	_, err = parser.Parse()
	if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
