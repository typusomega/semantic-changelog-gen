// Package builder contains the builder extracting a git repository's log into a Changelog
package builder

import (
	"github.com/typusomega/semantic-changelog-gen/pkg/changelog"
	"github.com/typusomega/semantic-changelog-gen/pkg/git"
)

// A Builder builds a changelog.
type Builder interface {
	Build() (*changelog.Changelog, error)
}

// New create a new instance of a Builder.
func New(repository git.Repository) Builder {
	return &builder{
		repository: repository,
	}
}

// Build builds a changelog from a git history.
func (it *builder) Build() (*changelog.Changelog, error) {
	chlog := changelog.NewChangelog()

	log, err := it.repository.GetLog()
	if err != nil {
		return chlog, err
	}

	release := chlog.NewRelease("tbd")
	for _, commit := range log {
		if commit.IsTagged() {
			release = chlog.NewRelease(commit.Tag)
		}

		release.AddEntry(commit)
	}

	return chlog, err
}

type builder struct {
	repository git.Repository
}
