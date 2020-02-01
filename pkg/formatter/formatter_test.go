package formatter

import (
	"testing"

	"github.com/joomcode/errorx"
	"github.com/stretchr/testify/assert"

	"github.com/typusomega/semantic-changelog-gen/pkg/changelog"
)

func TestNewTemplateFormatter(t *testing.T) {
	type args struct {
		opts []Option
	}
	tests := []struct {
		name  string
		given args
		then  func(t *testing.T, formatter Formatter, err error)
	}{
		{
			name:  "no options",
			given: args{opts: []Option{}},
			then: func(t *testing.T, formatter Formatter, err error) {
				assert.NoError(t, err)
				assert.IsType(t, &TemplateFormatter{}, formatter)
				assert.Equal(t, formatter.(*TemplateFormatter).template, markdownTemplate)
			},
		},
		{
			name:  "with markdown",
			given: args{opts: []Option{WithFormat(MarkdownFormat)}},
			then: func(t *testing.T, formatter Formatter, err error) {
				assert.NoError(t, err)
				assert.IsType(t, &TemplateFormatter{}, formatter)
				assert.Equal(t, formatter.(*TemplateFormatter).template, markdownTemplate)
			},
		},
		{
			name:  "with markdown; given set",
			given: args{opts: []Option{WithFormat(MarkdownFormat), WithTemplate("ignored given")}},
			then: func(t *testing.T, formatter Formatter, err error) {
				assert.NoError(t, err)
				assert.IsType(t, &TemplateFormatter{}, formatter)
				assert.Equal(t, formatter.(*TemplateFormatter).template, markdownTemplate)
			},
		},
		{
			name:  "with custom; no given set",
			given: args{opts: []Option{WithFormat(CustomFormat)}},
			then: func(t *testing.T, formatter Formatter, err error) {
				assert.Error(t, err)
				assert.True(t, errorx.IsOfType(err, errorx.IllegalArgument))
			},
		},
		{
			name:  "with custom; given set",
			given: args{opts: []Option{WithFormat(CustomFormat), WithTemplate("used given")}},
			then: func(t *testing.T, formatter Formatter, err error) {
				assert.NoError(t, err)
				assert.Equal(t, formatter.(*TemplateFormatter).template, "used given")
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTemplateFormatter(tt.given.opts...)
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
			it := &TemplateFormatter{
				template: tt.template,
			}
			format, err := it.Format(tt.chlog)
			tt.then(t, format, err)
		})
	}
}
