package commands

type Opts struct {
	Directory string `short:"d" long:"directory" description:"The directory containing the git repository to generate the changelog from" value-name:"DIR" default:"."`
}
