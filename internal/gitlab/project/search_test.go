package project

import (
	"errors"
	"testing"

	"github.com/alekseiapa/glcli/internal/gitlab/contract/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/xanzy/go-gitlab"
)

func TestMR(t *testing.T) {
	project := "foo/bar"
	query := "query string"
	mockGitlabSearch := &mocks.GitlabSearch{}
	mockGitlabSearch.On("MergeRequestsByProject", project, query, &gitlab.SearchOptions{}).
		Once().
		Return([]*gitlab.MergeRequest{}, &gitlab.Response{}, errors.New(""))

	s := &searchService{project: project, gitlabSearch: mockGitlabSearch}
	err := s.MR(query)

	assert.Error(t, err)
	mockGitlabSearch.AssertExpectations(t)
}
