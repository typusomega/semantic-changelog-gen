package formatter

const markdownTemplate = `# Changelog
{{- range $release := .Releases }}
## {{ $release.Version }}

{{- if .Features }}
### Features
{{- end }}
{{- range $commit := .Features }}
- {{ $commit.Description }}
{{- end }}

{{- if .Fixes }}
### Fixes
{{- end }}
{{- range $commit := .Fixes }}
- {{ $commit.Description }}
{{- end }}

{{- if .BreakingChanges }}
### Breaking
{{- end }}
{{- range $commit := .BreakingChanges }}
- {{ $commit.Description }}
{{- end }}

{{- end }}
`
