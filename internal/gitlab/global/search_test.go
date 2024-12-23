package global

import (
	"errors"
	"testing"

	"github.com/alekseiapa/glcli/internal/gitlab/contract/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/xanzy/go-gitlab"
)

func TestMR(t *testing.T) {
	query := "query string"
	mockGitlabSearch := &mocks.GitlabSearch{}
	mockGitlabSearch.On("MergeRequests", query, &gitlab.SearchOptions{}).
		Once().
		Return([]*gitlab.MergeRequest{}, &gitlab.Response{}, errors.New(""))

	s := &searchService{gitlabSearch: mockGitlabSearch}
	err := s.MR(query)

	assert.Error(t, err)
	mockGitlabSearch.AssertExpectations(t)
}

func TestProject(t *testing.T) {
	query := "query string"
	mockGitlabSearch := &mocks.GitlabSearch{}
	mockGitlabSearch.On("Projects", query, &gitlab.SearchOptions{}).
		Once().
		Return([]*gitlab.Project{}, &gitlab.Response{}, errors.New(""))

	s := &searchService{gitlabSearch: mockGitlabSearch}
	err := s.Project(query)

	assert.Error(t, err)
	mockGitlabSearch.AssertExpectations(t)
}
