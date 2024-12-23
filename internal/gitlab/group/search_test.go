package group

import (
	"errors"
	"testing"

	"github.com/alekseiapa/glcli/internal/gitlab/contract/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/xanzy/go-gitlab"
)

func TestMR(t *testing.T) {
	group := "foo"
	query := "query string"
	mockGitlabSearch := &mocks.GitlabSearch{}
	mockGitlabSearch.On("MergeRequestsByGroup", group, query, &gitlab.SearchOptions{}).
		Once().
		Return([]*gitlab.MergeRequest{}, &gitlab.Response{}, errors.New(""))

	s := &searchService{group: group, gitlabSearch: mockGitlabSearch}
	err := s.MR(query)

	assert.Error(t, err)
	mockGitlabSearch.AssertExpectations(t)
}

func TestProject(t *testing.T) {
	group := "foo"
	query := "query string"
	mockGitlabSearch := &mocks.GitlabSearch{}
	mockGitlabSearch.On("ProjectsByGroup", group, query, &gitlab.SearchOptions{}).
		Once().
		Return([]*gitlab.Project{}, &gitlab.Response{}, errors.New(""))

	s := &searchService{group: group, gitlabSearch: mockGitlabSearch}
	err := s.Project(query)

	assert.Error(t, err)
	mockGitlabSearch.AssertExpectations(t)
}
