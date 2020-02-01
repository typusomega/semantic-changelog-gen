package commands

import "github.com/typusomega/semantic-changelog-gen/pkg/formatter"

// Opts define the command arguments of the generateCommand.
type Opts struct {
	OutputOptions
	// GitRepository is the directory containing the git repository to.
	GitRepository string `short:"d" long:"dir" description:"The directory containing the git repository to" value-name:"DIR" default:"."`
}

// OutputOptions define the set of all output parameters.
type OutputOptions struct {
	// OutputFile is path to the output file.
	OutputFile string `short:"o"  long:"out" description:"The path to the output file" value-name:"FILE" default:"./changelog.md"`
	// Format is the format to be used.
	OutputFormat formatter.Format `short:"f"  long:"format" description:"The output format to be used. Possible values: 'markdown', 'custom'" value-name:"FORMAT" default:"markdown"`
	// OutputTemplateFile is the Go template to be used for formatting.
	OutputTemplateFile string `short:"t"  long:"template" description:"The path to the Go template to be used for formatting. Only taken into account when 'format' is 'custom'" value-name:"TEMPLATE-FILE-PATH"`
}
