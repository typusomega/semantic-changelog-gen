// Package formatter contains everything related to changelog Formatters.
package formatter

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/joomcode/errorx"

	"github.com/typusomega/semantic-changelog-gen/pkg/changelog"
)

// Formatter formats Changelogs.
type Formatter interface {
	// Format formats the given changelog.
	Format(changelog *changelog.Changelog) (string, error)
}

// NewCustomTemplateFormatter creates a new instance of a CustomTemplateFormatter
func NewCustomTemplateFormatter(goTemplate string) (Formatter, error) {
	if strings.Trim(goTemplate, " \t\r\n") == "" {
		return nil, errorx.IllegalArgument.New("template must not be empty")
	}
	return &CustomTemplateFormatter{
		template: goTemplate,
	}, nil
}

func format(chlog *changelog.Changelog, templ string, options interface{}) (string, error) {
	if chlog == nil {
		return "", errorx.IllegalArgument.New("changelog must not be nil")
	}

	tpl, err := template.New("changelog").Parse(templ)
	if err != nil {
		return "", errorx.IllegalArgument.Wrap(err, "invalid template")
	}

	bufferString := bytes.NewBufferString("")
	err = tpl.Execute(bufferString,
		struct {
			Releases []*changelog.Release
			Options  interface{}
		}{Releases: chlog.Releases, Options: options})
	if err != nil {
		return "", errorx.IllegalArgument.Wrap(err, "could not execute template")
	}

	return bufferString.String(), nil
}

// Format renders the given changelog with the CustomTemplateFormatter's template.
func (it *CustomTemplateFormatter) Format(chlog *changelog.Changelog) (string, error) {
	return format(chlog, it.template, nil)
}

// A CustomTemplateFormatter is a Formatter rendering changelogs with Go templates
type CustomTemplateFormatter struct {
	template string
}
