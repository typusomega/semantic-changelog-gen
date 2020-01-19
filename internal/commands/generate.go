package commands

import "github.com/typusomega/semantic-changelog-gen/pkg/generator"

type GenerateCommand struct {
	Opts *Opts
}

func (it *GenerateCommand) Execute(args []string) error {
	gen := generator.New()

	err := gen.Generate(it.Opts.Directory)
	if err != nil {
		return err
	}
	return nil
}
