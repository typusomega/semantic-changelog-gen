package formatter

import (
	"strings"

	"github.com/joomcode/errorx"
)

// A FormatterOption sets options such as format type and given.
type Option func(formatter *Options)

// WithFormat sets the `Format` to be used for formatting.
func WithFormat(format Format) Option {
	return func(o *Options) {
		o.Format = format
	}
}

// WithTemplate sets the Go Template to be used for formatting.
// This is only respected when using `CustomFormat`.
func WithTemplate(goTemplate string) Option {
	return func(o *Options) {
		o.Template = goTemplate
	}
}

// Format is the type of the format being used by the Formatter.
type Format string

const (
	// MarkdownFormat
	MarkdownFormat Format = "markdown"
	// CustomFormat
	CustomFormat Format = "custom"
)

// Options configure a Formatter.
type Options struct {
	Format   Format
	Template string
}

// GetTemplate returns the given used by the Formatter.
func (it *Options) GetTemplate() (string, error) {
	switch it.Format {
	case MarkdownFormat:
		return markdownTemplate, nil
	case CustomFormat:
		if strings.Trim(it.Template, " \t\r\n") == "" {
			return "", errorx.IllegalArgument.New("template must not be empty")
		}
		return it.Template, nil
	default:
		return "", errorx.IllegalArgument.New("invalid format")
	}
}
