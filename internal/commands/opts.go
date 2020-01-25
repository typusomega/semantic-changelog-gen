package commands

// Opts define the command arguments of the generateCommand.
type Opts struct {
	// GitRepository is the directory containing the git repository to.
	GitRepository string `short:"d" long:"dir" description:"The directory containing the git repository to" value-name:"DIR" default:"."`
	// OutputFile is path to the output file.
	OutputFile string `short:"o"  long:"out" description:"The path to the output file" value-name:"FILE" default:"./changelog.md"`
}
