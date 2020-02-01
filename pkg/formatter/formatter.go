// Package formatter contains everything related to changelog Formatters.
package formatter

import (
	"bytes"
	"text/template"

	"github.com/joomcode/errorx"

	"github.com/typusomega/semantic-changelog-gen/pkg/changelog"
)

// Formatter formats Changelogs.
type Formatter interface {
	// Format formats the given changelog.
	Format(changelog *changelog.Changelog) (string, error)
}

// NewTemplateFormatter creates a new instance of a TemplateFormatter
func NewTemplateFormatter(opts ...Option) (Formatter, error) {
	options := &Options{Format: MarkdownFormat}
	for _, opt := range opts {
		opt(options)
	}

	tpl, err := options.GetTemplate()
	if err != nil {
		return nil, err
	}

	return &TemplateFormatter{
		template: tpl,
	}, nil
}

// Format renders the given changelog in Markdown given.
func (it *TemplateFormatter) Format(chlog *changelog.Changelog) (string, error) {
	if chlog == nil {
		return "", errorx.IllegalArgument.New("changelog must not be nil")
	}

	tpl, err := template.New("changelog").Parse(it.template)
	if err != nil {
		return "", errorx.IllegalArgument.Wrap(err, "invalid template")
	}

	bufferString := bytes.NewBufferString("")
	err = tpl.Execute(bufferString, chlog)
	if err != nil {
		return "", errorx.IllegalArgument.Wrap(err, "could not execute template")
	}

	return bufferString.String(), nil
}

// A TemplateFormatter is a Formatter rendering markdown
type TemplateFormatter struct {
	template string
}
