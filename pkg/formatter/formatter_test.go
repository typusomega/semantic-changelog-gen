package formatter

import (
	"testing"

	"github.com/joomcode/errorx"
	"github.com/stretchr/testify/assert"

	"github.com/typusomega/semantic-changelog-gen/pkg/changelog"
)

func TestNewCustomTemplateFormatter(t *testing.T) {
	tests := []struct {
		name     string
		template string
		then     func(t *testing.T, formatter Formatter, err error)
	}{
		{
			name:     "with custom; no given template",
			template: " ",
			then: func(t *testing.T, formatter Formatter, err error) {
				assert.Error(t, err)
				assert.True(t, errorx.IsOfType(err, errorx.IllegalArgument))
			},
		},
		{
			name:     "with custom; template set",
			template: "template given",
			then: func(t *testing.T, formatter Formatter, err error) {
				assert.NoError(t, err)
				assert.Equal(t, formatter.(*CustomTemplateFormatter).template, "template given")
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCustomTemplateFormatter(tt.template)
			tt.then(t, got, err)
		})
	}
}

func TestTemplateFormatter_Format(t *testing.T) {
	tests := []struct {
		name     string
		template string
		chlog    *changelog.Changelog
		then     func(t *testing.T, fmt string, err error)
	}{
		{
			name:     "nil changelog",
			template: markdownTemplate,
			chlog:    nil,
			then: func(t *testing.T, fmt string, err error) {
				assert.Error(t, err)
				assert.True(t, errorx.IsOfType(err, errorx.IllegalArgument))
			},
		},
		{
			name:     "invalid template; access to non-existing variables",
			template: "{{ .NotExistingValue }}",
			chlog:    &changelog.Changelog{},
			then: func(t *testing.T, fmt string, err error) {
				assert.Error(t, err)
				assert.True(t, errorx.IsOfType(err, errorx.IllegalArgument))
			},
		},
		{
			name:     "invalid template; not a go template",
			template: "{{ .NotExistingValue }",
			chlog:    &changelog.Changelog{},
			then: func(t *testing.T, fmt string, err error) {
				assert.Error(t, err)
				assert.True(t, errorx.IsOfType(err, errorx.IllegalArgument))
			},
		},
		{
			name:     "success",
			template: "just a template",
			chlog:    &changelog.Changelog{},
			then: func(t *testing.T, fmt string, err error) {
				assert.NoError(t, err)
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			it := &CustomTemplateFormatter{
				template: tt.template,
			}
			format, err := it.Format(tt.chlog)
			tt.then(t, format, err)
		})
	}
}
