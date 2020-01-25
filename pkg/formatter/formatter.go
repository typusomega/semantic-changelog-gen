package formatter

import (
	"github.com/typusomega/semantic-changelog-gen/pkg/changelog"
)

// Formatter formats changelogs.
type Formatter interface {
	// Format formats the given changelog.
	Format(changelog *changelog.Changelog) (string, error)
}
