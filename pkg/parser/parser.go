package parser

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/typusomega/semantic-changelog-gen/pkg/changelog"
)

// Parser is meant to parse semantic commit messages into `changelog.SemanticCommit` objects.
type Parser interface {
	Parse(commitMessage string) (*changelog.SemanticCommit, error)
}

// New creates a new instance of a Parser
func New() Parser {
	return &parser{}
}

type parser struct {
}

func (it *parser) Parse(commitMessage string) (*changelog.SemanticCommit, error) {
	commit := &changelog.SemanticCommit{}

	lines := strings.Split(commitMessage, "\n")
	if len(lines) == 0 {
		return commit, nil
	}

	commit = it.parseFirstLine(lines[0])

	cursor := 2
	commit.Body, cursor = it.parseBody(cursor, lines)
	if cursor < len(lines) {
		commit.Footers = lines[cursor : len(lines)-1]
	}

	if strings.Contains(commitMessage, "BREAKING") {
		commit.CommitType = changelog.Breaking
	}
	return commit, nil
}

func (it *parser) parseBody(start int, lines []string) (string, int) {
	paragraph := ""
	cursor := start
	if len(lines) > 0 {
		for ; cursor < len(lines); cursor++ {
			if lines[cursor] == "" {
				break
			}
			paragraph = fmt.Sprintf("%s%s\n", paragraph, lines[cursor])
		}
	}
	return paragraph, cursor
}

func (it *parser) parseFirstLine(line string) *changelog.SemanticCommit {
	commit := &changelog.SemanticCommit{}
	match := firstLine.FindStringSubmatch(line)
	for i, name := range firstLine.SubexpNames() {
		if i < len(match) {
			switch name {
			case "type":
				commit.CommitType = changelog.CommitType(match[i])
			case "description":
				commit.Description = match[i]
			case "component":
				commit.Component = match[i]
			}
		}
	}
	return commit
}

var firstLine = regexp.MustCompile(`(?P<type>feat|fix|chore|docs|test|refactor|style)\(?(?P<component>[^\)]*)\)?:\W?(?P<description>.*)`)
