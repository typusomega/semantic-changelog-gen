package formatter

const markdownTemplate = `# Changelog
{{- range $release := .Releases }}
## {{ $release.Version }}

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
`
