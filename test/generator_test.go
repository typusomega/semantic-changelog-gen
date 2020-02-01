package test

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xorcare/golden"

	bldr "github.com/typusomega/semantic-changelog-gen/pkg/builder"
	"github.com/typusomega/semantic-changelog-gen/pkg/formatter"
	"github.com/typusomega/semantic-changelog-gen/pkg/git"
)

func TestMarkdown(t *testing.T) {
	builder := bldr.New(git.NewRepository("../", git.NewParser()))

	changelog, err := builder.Build()
	assert.NoError(t, err)

	fmt, err := formatter.NewTemplateFormatter(formatter.WithFormat(formatter.MarkdownFormat))
	assert.NoError(t, err)

	formattedChangelog, err := fmt.Format(changelog)
	assert.NoError(t, err)

	golden.Assert(t, []byte(formattedChangelog))
}

func TestTemplate(t *testing.T) {
	builder := bldr.New(git.NewRepository("../", git.NewParser()))

	changelog, err := builder.Build()
	assert.NoError(t, err)

	file, err := ioutil.ReadFile("testdata/test_template.gohtml")
	assert.NoError(t, err)

	fmt, err := formatter.NewTemplateFormatter(formatter.WithFormat(formatter.CustomFormat), formatter.WithTemplate(string(file)))
	assert.NoError(t, err)

	formattedChangelog, err := fmt.Format(changelog)
	assert.NoError(t, err)

	golden.Assert(t, []byte(formattedChangelog))
}
