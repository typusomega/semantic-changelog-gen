package formatter

import (
	"fmt"
	"strings"

	"github.com/typusomega/semantic-changelog-gen/pkg/changelog"
)

// NewMarkdownFormatter creates a new instance of a stateful MarkdownFormatter
func NewMarkdownFormatter() Formatter {
	return &MarkdownFormatter{
		txt: strings.Builder{},
	}
}

// Format renders the given changelog in Markdown format.
// This operation is stateful.
func (it *MarkdownFormatter) Format(chlog *changelog.Changelog) (string, error) {
	for _, release := range chlog.Releases {
		it.txt.WriteString(fmt.Sprintf("## %s", release.Version))
		it.linefeed()
		it.writeSection(featuresHeading, release.Features)
		it.writeSection(fixesHeading, release.Fixes)
		it.writeSection(breakingHeading, release.BreakingChanges)
	}
	return it.txt.String(), nil
}

func (it *MarkdownFormatter) writeSection(heading string, commits []*changelog.SemanticCommit) {
	if len(commits) == 0 {
		return
	}

	it.txt.WriteString(heading)
	it.linefeed()
	for _, commit := range commits {
		it.txt.WriteString(listEntryPrefix)
		it.txt.WriteString(commit.Description)
		it.linefeed()
	}
	it.linefeed()
}

func (it *MarkdownFormatter) linefeed() {
	it.txt.WriteString("\n")
}

// A MarkdownFormatter is a Formatter rendering markdown
type MarkdownFormatter struct {
	txt strings.Builder
}

const fixesHeading = "### Fixes\n"
const featuresHeading = "### Features\n"
const breakingHeading = "### Breaking\n"
const listEntryPrefix = "- "
