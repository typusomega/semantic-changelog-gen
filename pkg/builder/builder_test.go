package builder_test

//go:generate mockgen -package mocks -destination=./../mocks/mock_repository.go github.com/typusomega/semantic-changelog-gen/pkg/git Repository

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	bldr "github.com/typusomega/semantic-changelog-gen/pkg/builder"
	"github.com/typusomega/semantic-changelog-gen/pkg/changelog"
	"github.com/typusomega/semantic-changelog-gen/pkg/mocks"
)

func Test_builder_Build(t *testing.T) {
	tests := []struct {
		name  string
		given func(repository *mocks.MockRepository)
		then  func(t *testing.T, chlog *changelog.Changelog, err error)
	}{
		{
			name: "repository.Log returns error",
			given: func(repository *mocks.MockRepository) {
				repository.EXPECT().GetLog().Times(1).Return(nil, errors.New("failed"))
			},
			then: func(t *testing.T, chlog *changelog.Changelog, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name: "no tags",
			given: func(repository *mocks.MockRepository) {
				repository.EXPECT().GetLog().Times(1).Return(onlyUntaggedCommits(), nil)
			},
			then: func(t *testing.T, chlog *changelog.Changelog, err error) {
				assert.Nil(t, err)
				assert.Len(t, chlog.Releases, 1)
				assert.Equal(t, currentVersion, chlog.Releases[0].Version)
			},
		},
		{
			name: "tags",
			given: func(repository *mocks.MockRepository) {
				repository.EXPECT().GetLog().Times(1).Return(commits, nil)
			},
			then: func(t *testing.T, chlog *changelog.Changelog, err error) {
				assert.Nil(t, err)
				assert.Len(t, chlog.Releases, 3)
				assert.Equal(t, currentVersion, chlog.Releases[0].Version)
				assert.Equal(t, chlog.Releases[1].Version, v101)
				assert.Equal(t, chlog.Releases[2].Version, v1)
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()
			repoMock := mocks.NewMockRepository(controller)
			if tt.given != nil {
				tt.given(repoMock)
			}

			b := bldr.New(repoMock)
			chlog, err := b.Build()
			tt.then(t, chlog, err)
		})
	}
}

func onlyUntaggedCommits() []*changelog.SemanticCommit {
	cms := make([]*changelog.SemanticCommit, 0)
	for _, cm := range cms {
		if !cm.IsTagged() {
			cms = append(cms, cm)
		}
	}
	return cms
}

const currentVersion = "tbd"
const v1 = "v1.0.0"
const v101 = "v1.0.1"

var (
	commits = []*changelog.SemanticCommit{
		{
			Hash:        "6",
			Tag:         "",
			Description: "stuff happened",
			Body:        "body",
			CommitType:  changelog.Fix,
		},
		{
			Hash:        "5",
			Tag:         v101,
			Description: "stuff happened",
			Body:        "body",
			CommitType:  changelog.Fix,
		},
		{
			Hash:        "4",
			Tag:         "",
			Description: "stuff happened",
			Body:        "body",
			CommitType:  changelog.Feature,
		},
		{
			Hash:        "3",
			Tag:         v1,
			Description: "stuff happened",
			Body:        "body",
			CommitType:  changelog.Feature,
		},
		{
			Hash:        "2",
			Tag:         "",
			Description: "stuff happened",
			Body:        "body",
			CommitType:  changelog.Chore,
		},
		{
			Hash:        "1",
			Tag:         "",
			Description: "stuff happened",
			Body:        "body",
			CommitType:  changelog.Chore,
		},
	}
)
