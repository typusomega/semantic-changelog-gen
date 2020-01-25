package builder

import (
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"

	"github.com/typusomega/semantic-changelog-gen/pkg/changelog"
	"github.com/typusomega/semantic-changelog-gen/pkg/parser"
)

// A Builder builds a changelog.
type Builder interface {
	Build() (*changelog.Changelog, error)
}

// New create a new instance of a Builder.
// This builder is stateful and should not be used for more than one changelog.
func New(options Options) Builder {
	return &generator{
		repository: options.Repository,
		parser:     options.Parser,
		tags:       make(map[plumbing.Hash]*object.Tag),
	}
}

// Build builds a changelog from a git history.
func (it *generator) Build() (*changelog.Changelog, error) {
	chlog := changelog.NewChangelog()

	repository, err := it.open()
	if err != nil {
		return nil, err
	}

	tags, err := repository.TagObjects()
	if err != nil {
		return nil, err
	}

	err = tags.ForEach(func(tag *object.Tag) error {
		it.tags[tag.Target] = tag
		return nil
	})
	if err != nil {
		return nil, err
	}

	reference, err := repository.Head()
	if err != nil {
		return nil, err
	}

	log, err := repository.Log(&git.LogOptions{
		From: reference.Hash(),
	})
	if err != nil {
		return nil, err
	}

	release := chlog.NewRelease("tbd")
	err = log.ForEach(func(commit *object.Commit) error {
		if tag, isTagged := it.tags[commit.Hash]; isTagged {
			release = chlog.NewRelease(tag.Name)
		}

		entry, err := it.parser.Parse(commit.Message)
		if err != nil {
			return err
		}

		release.AddEntry(entry)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return chlog, err
}

func (it *generator) open() (*git.Repository, error) {
	return git.PlainOpen(it.repository)
}

type generator struct {
	repository string
	parser     parser.Parser
	tags       map[plumbing.Hash]*object.Tag
}

// Options contain the configuration parameters of a Builder.
type Options struct {
	Repository string
	Parser     parser.Parser
}
