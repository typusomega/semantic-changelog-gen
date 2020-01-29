package changelog

import "strings"

// NewChangelog creates a new instance of a Changelog.
func NewChangelog() *Changelog {
	return &Changelog{
		Releases: make([]*Release, 0),
	}
}

// A Changelog represents the collection of the releases.
type Changelog struct {
	Releases []*Release
}

// NewRelease creates a new instance of a Release as part of the Changelog.
func (it *Changelog) NewRelease(version string) *Release {
	release := &Release{
		Version:  version,
		Fixes:    make([]*SemanticCommit, 0),
		Features: make([]*SemanticCommit, 0),
	}
	it.Releases = append(it.Releases, release)
	return release
}

// A Release represents a version containing features, fixes and breaking changes.
type Release struct {
	Features        []*SemanticCommit
	Fixes           []*SemanticCommit
	BreakingChanges []*SemanticCommit
	Version         string
}

// AddEntry adds the given entry to this Release.
func (it *Release) AddEntry(entry *SemanticCommit) {
	switch entry.CommitType {
	case Fix:
		it.Fixes = append(it.Fixes, entry)
	case Feature:
		it.Features = append(it.Features, entry)
	case Breaking:
		it.BreakingChanges = append(it.BreakingChanges, entry)
	}
}

// A SemanticCommit represents a commit following the rules of conventional commit.
// See also:
// - https://www.conventionalcommits.org/en/v1.0.0-beta.4/
// - http://karma-runner.github.io/4.0/dev/git-commit-msg.html
type SemanticCommit struct {
	Hash        string
	Tag         string
	Description string
	Body        string
	Scope       string
	Footers     []string
	CommitType
}

func NewSemanticCommit() *SemanticCommit {
	return &SemanticCommit{Footers: make([]string, 0)}
}

func (it *SemanticCommit) IsTagged() bool {
	return it.Tag == strings.Trim("", "\t\n ")
}

// A CommitType is one of the given set of commit types as specified by http://karma-runner.github.io/4.0/dev/git-commit-msg.html.
type CommitType string

// The diverse semantic commit types stated in http://karma-runner.github.io/4.0/dev/git-commit-msg.html.
const (
	Feature  CommitType = "feat"
	Fix      CommitType = "fix"
	Breaking CommitType = "BREAKING"
	Chore    CommitType = "chore"
	Docs     CommitType = "docs"
	Tests    CommitType = "test"
	Refactor CommitType = "refactor"
	Style    CommitType = "style"
)
