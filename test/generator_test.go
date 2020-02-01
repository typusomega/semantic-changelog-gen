package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xorcare/golden"

	bldr "github.com/typusomega/semantic-changelog-gen/pkg/builder"
	"github.com/typusomega/semantic-changelog-gen/pkg/formatter"
	"github.com/typusomega/semantic-changelog-gen/pkg/git"
)

func TestChangelogGenerator(t *testing.T) {
	builder := bldr.New(git.NewRepository("../", git.NewParser()))

	changelog, err := builder.Build()
	assert.Nil(t, err)

	formattedChangelog, err := formatter.NewMarkdownFormatter().Format(changelog)
	assert.Nil(t, err)

	golden.Assert(t, []byte(formattedChangelog))
}
