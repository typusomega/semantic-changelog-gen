package commands

import (
	fmt "fmt"
	"io/ioutil"

	"github.com/jessevdk/go-flags"

	"github.com/typusomega/semantic-changelog-gen/pkg/builder"
	"github.com/typusomega/semantic-changelog-gen/pkg/formatter"
	"github.com/typusomega/semantic-changelog-gen/pkg/parser"
)

// NewGenerateCommand creates a new instance of a generateCommand
func NewGenerateCommand(opts *Opts) flags.Commander {
	return &generateCommand{opts: opts}
}

// Execute runs the command
func (it *generateCommand) Execute(args []string) error {
	println(fmt.Sprintf("generating changelog for git repository '%s'...", it.opts.GitRepository))
	bldr := builder.New(builder.Options{
		Repository: it.opts.GitRepository,
		Parser:     parser.New(),
	})

	println("building...")
	changelog, err := bldr.Build()
	if err != nil {
		return err
	}

	println("formatting...")
	fmter := formatter.NewMarkdownFormatter()
	formattedChangelog, err := fmter.Format(changelog)
	if err != nil {
		return err
	}

	println(fmt.Sprintf("writing changelog to '%s'...", it.opts.OutputFile))
	err = ioutil.WriteFile(it.opts.OutputFile, []byte(formattedChangelog), 0644)
	if err != nil {
		return err
	}

	println("done.")

	return err
}

type generateCommand struct {
	opts *Opts
}
