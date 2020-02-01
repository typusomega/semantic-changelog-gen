// Package changelog contains the domain model of a Changelog, Release and SemanticCommits.
package changelog

import (
	"strings"
)

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
		Version: version,
		Commits: make([]*SemanticCommit, 0),
	}
	it.Releases = append(it.Releases, release)
	return release
}

// A Release represents a version containing features, fixes and breaking changes.
type Release struct {
	Commits []*SemanticCommit
	Version string
}

// AddEntry adds the given entry to this Release.
func (it *Release) AddEntry(entry *SemanticCommit) {
	it.Commits = append(it.Commits, entry)
}

// GetFixes returns all fixes of this Release.
func (it *Release) GetFixes() []*SemanticCommit {
	return it.filterBy(Fix)
}

// GetFeatures returns all features of this Release.
func (it *Release) GetFeatures() []*SemanticCommit {
	return it.filterBy(Feature)
}

// GetBreaking returns all breaking changes of this Release.
func (it *Release) GetBreaking() []*SemanticCommit {
	return it.filterBy(Breaking)
}

// GetChore returns all chore commits of this Release.
func (it *Release) GetChore() []*SemanticCommit {
	return it.filterBy(Chore)
}

// GetDocs returns all docs commits of this Release.
func (it *Release) GetDocs() []*SemanticCommit {
	return it.filterBy(Docs)
}

// GetTest returns all test commits of this Release.
func (it *Release) GetTest() []*SemanticCommit {
	return it.filterBy(Tests)
}

// GetRefactor returns all refactor commits of this Release.
func (it *Release) GetRefactor() []*SemanticCommit {
	return it.filterBy(Refactor)
}

// GetStyle returns all style commits of this Release.
func (it *Release) GetStyle() []*SemanticCommit {
	return it.filterBy(Style)
}

// GetScoped returns all commits of the given commitType in this Release grouped by their scope.
func (it *Release) GetScoped(commitType CommitType) map[string][]*SemanticCommit {
	scoped := make(map[string][]*SemanticCommit)
	for _, commit := range it.filterBy(commitType) {
		scoped[commit.Scope] = append(scoped[commit.Scope], commit)
	}
	return scoped
}

func (it *Release) filterBy(commitType CommitType) []*SemanticCommit {
	filtered := make([]*SemanticCommit, 0)
	for _, commit := range it.Commits {
		if commit.CommitType == commitType {
			filtered = append(filtered, commit)
		}
	}
	return filtered
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

// NewSemanticCommit creates a new instace of a SemanticCommit.
func NewSemanticCommit() *SemanticCommit {
	return &SemanticCommit{Footers: make([]string, 0)}
}

// IsTagged tells whether the commit is tagged or not.
func (it *SemanticCommit) IsTagged() bool {
	return it.Tag != strings.Trim("", "\t\n ")
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
