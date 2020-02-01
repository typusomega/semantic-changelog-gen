package formatter

// A MarkdownOption sets options for markdown format such as whether or not to include scopes.
type MarkdownOption func(opt *markdownOptions)

// WithScopes includes scopes to markdown format.
func WithScopes() MarkdownOption {
	return func(o *markdownOptions) {
		o.IncludeScopes = true
	}
}
