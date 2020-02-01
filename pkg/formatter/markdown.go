package formatter

import (
	"github.com/typusomega/semantic-changelog-gen/pkg/changelog"
)

// NewMarkdownFormatter creates a new instance of a MarkdownFormatter
func NewMarkdownFormatter(opts ...MarkdownOption) (Formatter, error) {
	options := &markdownOptions{}
	for _, opt := range opts {
		opt(options)
	}

	return &MarkdownFormatter{
		MarkdownOptions: options,
	}, nil
}

const markdownTemplate = `# Changelog
{{ $includeScopes := .Options.IncludeScopes }}
{{- range $release := .Releases }}
## {{ $release.Version }}
  {{- if $includeScopes }}
    {{- $scopedFeatures := $release.GetScoped "feat" }}
    {{- if $scopedFeatures }}
### Features
    {{- end }}
    {{- range $scope, $commits := $scopedFeatures }}
		{{- if $scope }}
#### {{ $scope }}
		{{- end }}
      {{- range $commit := $commits }}
- {{ $commit.Description }}
      {{- end }}
    {{- end }}

    {{- $scopedFixes := $release.GetScoped "fix" }}
    {{- if $scopedFixes }}
### Fixes
    {{- end }}
    {{- range $scope, $commits := $scopedFixes }}
		{{- if $scope }}
#### {{ $scope }}
		{{- end }}
      {{- range $commit := $commits }}
- {{ $commit.Description }}
      {{- end }}
    {{- end }}

    {{- $scopedBreaking := $release.GetScoped "BREAKING" }}
    {{- if $scopedBreaking }}
### Breaking
    {{- end }}
    {{- range $scope, $commits := $scopedBreaking }}
		{{- if $scope }}
#### {{ $scope }}
		{{- end }}
      {{- range $commit := $commits }}
- {{ $commit.Description }}
      {{- end }}
    {{- end }}

  {{- else }}

    {{- $features := .GetFeatures }}
    {{- if $features }}
### Features
    {{- end }}
    {{- range $commit := .GetFeatures }}
- {{ $commit.Description }}
    {{- end }}

    {{- $fixes := .GetFixes }}
    {{- if $fixes }}
### Fixes
    {{- end }}
    {{- range $commit := $fixes }}
- {{ $commit.Description }}
    {{- end }}

    {{- $breaking := .GetBreaking }}
    {{- if $breaking }}
### Breaking
    {{- end }}
    {{- range $commit := $breaking }}
- {{ $commit.Description }}
    {{- end }}
{{- end }}
{{- end }}
`

// A MarkdownFormatter is a formatter rendering Markdown formats.
type MarkdownFormatter struct {
	MarkdownOptions *markdownOptions
}

type markdownOptions struct {
	IncludeScopes bool
}

// Format formats the given changelog in Markdown format.
func (it *MarkdownFormatter) Format(changelog *changelog.Changelog) (string, error) {
	return format(changelog, markdownTemplate, it.MarkdownOptions)
}
