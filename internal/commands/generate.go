package commands

import (
	fmt "fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/jessevdk/go-flags"

	bldr "github.com/typusomega/semantic-changelog-gen/pkg/builder"
	"github.com/typusomega/semantic-changelog-gen/pkg/git"
)

// NewGenerateCommand creates a new instance of a generateCommand
func NewGenerateCommand(opts *Opts) flags.Commander {
	return &generateCommand{opts: opts}
}

// Execute runs the command
func (it *generateCommand) Execute(args []string) error {
	repositoryAbsoultePath, err := filepath.Abs(it.opts.GitRepository)
	if err != nil {
		return err
	}

	println(fmt.Sprintf("generating changelog for git repository '%s'...", repositoryAbsoultePath))
	builder := bldr.New(git.NewRepository(repositoryAbsoultePath, git.NewParser()))

	println("building...")
	changelog, err := builder.Build()
	if err != nil {
		return err
	}

	println("formatting...")
	fmter, err := it.opts.GetFormatter()
	if err != nil {
		return err
	}

	formattedChangelog, err := fmter.Format(changelog)
	if err != nil {
		return err
	}

	outfileAbsoultePath, err := filepath.Abs(it.opts.OutputFile)
	if err != nil {
		return err
	}

	println(fmt.Sprintf("writing changelog to '%s'...", outfileAbsoultePath))
	err = ioutil.WriteFile(outfileAbsoultePath, []byte(formattedChangelog), 0644)
	if err != nil {
		return err
	}

	println("done.")

	return err
}

type generateCommand struct {
	opts *Opts
}
