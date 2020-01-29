package git

import (
	"bufio"
	"io"
	"regexp"
	"strings"

	"github.com/typusomega/semantic-changelog-gen/pkg/changelog"
)

// Parser is meant to parse semantic commit messages into `changelog.SemanticCommit` objects.
type Parser interface {
	ParseLog(reader io.Reader) ([]*changelog.SemanticCommit, error)
	ParseTags(reader io.Reader) map[string]string
}

// NewParser creates a new instance of a Parser
func NewParser() Parser {
	return &parser{}
}

type parser struct {
}

func (it *parser) ParseTags(reader io.Reader) map[string]string {
	tags := make(map[string]string)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		commit, tag := it.parseTag(scanner.Text())
		tags[commit] = tag
	}

	return tags
}

func (it *parser) ParseLog(reader io.Reader) ([]*changelog.SemanticCommit, error) {
	var commits []*changelog.SemanticCommit
	commit := changelog.NewSemanticCommit()

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()

		if isEmpty(line) {
			continue
		}

		if strings.HasPrefix(line, "commit") {
			commit = changelog.NewSemanticCommit()
			commit.Hash, commit.Tag = it.parseHashAndTag(line)
			commits = append(commits, commit)
			continue
		}

		if strings.HasPrefix(line, "Author") {
			continue
		}

		if strings.HasPrefix(line, "Date") {
			continue
		}

		if strings.Contains(line, "BREAKING") {
			commit.CommitType = changelog.Breaking
		}

		if typ, desc, scope, ok := it.parseTitle(line); ok && isEmpty(commit.Description) {
			commit.CommitType = typ
			commit.Description = desc
			commit.Scope = scope
			continue
		}

		if isEmpty(commit.Body) {
			for scanner.Scan() {
				line = scanner.Text()
				if isEmpty(line) {
					break
				}
				commit.Body += line
			}
			continue
		}

		commit.Footers = append(commit.Footers, line)
	}

	return commits, nil
}

func (it *parser) parseHashAndTag(line string) (string, string) {
	hash := ""
	tag := ""

	match := commitHash.FindStringSubmatch(line)
	if len(match) == 0 {
		return hash, tag
	}
	for i, name := range commitHash.SubexpNames() {
		if i < len(match) {
			switch name {
			case "hash":
				hash = match[i]
			case "tag":
				tag = match[i]
			}
		}
	}
	return hash, tag
}

func (it *parser) parseTitle(line string) (tp changelog.CommitType, desc string, scope string, ok bool) {
	match := commitTitle.FindStringSubmatch(line)
	ok = true
	if len(match) == 0 {
		ok = false
	}

	for i, name := range commitTitle.SubexpNames() {
		if i < len(match) {
			switch name {
			case "type":
				tp = changelog.CommitType(match[i])
			case "description":
				desc = match[i]
			case "scope":
				scope = match[i]
			}
		}
	}
	return
}

func (it *parser) parseTag(line string) (string, string) {
	commit := ""
	tag := ""
	match := tagTarget.FindStringSubmatch(line)
	if len(match) == 0 {
		return commit, tag
	}

	for i, name := range tagTarget.SubexpNames() {
		if i < len(match) {
			switch name {
			case "commit":
				commit = match[i]
			case "tag":
				tag = match[i]
			}
		}
	}
	return commit, tag
}

func isEmpty(str string) bool {
	str = strings.TrimSpace(str)
	return str == ""
}

var commitTitle = regexp.MustCompile(`(?P<type>feat|fix|chore|docs|test|refactor|style)(\((?P<scope>[^\)]*)\))?:\W?(?P<description>.*)`)
var commitHash = regexp.MustCompile(`commit\W?(?P<hash>[a-f0-9]+)(\W\(tag:\W?(?P<tag>[^)]*)\))?`)
var tagTarget = regexp.MustCompile(`(?P<commit>[a-f0-9]+)\W?(refs/tags/(?P<tag>.*))?`)
