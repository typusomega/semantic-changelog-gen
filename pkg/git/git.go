// Package git contains the git interface.
package git

import (
	"os/exec"

	"github.com/typusomega/semantic-changelog-gen/pkg/changelog"
)

// A Repository represents a Git repository containing semantic commits.
type Repository interface {
	GetLog() ([]*changelog.SemanticCommit, error)
}

// NewRepository creates a new instace of a Repository
func NewRepository(directory string, parser Parser) Repository {
	return &repository{
		directory: directory,
		parser:    parser,
	}
}

type repository struct {
	directory string
	parser    Parser
}

func (it *repository) GetLog() ([]*changelog.SemanticCommit, error) {
	cmd := exec.Command("git", "log", "--tags", "--decorate", "HEAD")
	cmd.Dir = it.directory

	out, err := cmd.StdoutPipe()
	if err != nil {
		return make([]*changelog.SemanticCommit, 0), nil
	}

	err = cmd.Start()
	if err != nil {
		return make([]*changelog.SemanticCommit, 0), nil
	}

	log, err := it.parser.ParseLog(out)
	if err != nil {
		return nil, err
	}

	return log, cmd.Wait()
}
