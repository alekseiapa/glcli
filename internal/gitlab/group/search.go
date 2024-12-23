package group

import (
	"io"

	"github.com/alekseiapa/glcli/internal/gitlab/contract"
	"github.com/alekseiapa/glcli/internal/gitlab/render"
	"github.com/xanzy/go-gitlab"
)

type searchService struct {
	group        string
	gitlabSearch contract.GitlabSearch
	out          io.Writer
}

func (s *searchService) MR(query string) error {
	mrs, _, err := s.gitlabSearch.MergeRequestsByGroup(s.group, query, &gitlab.SearchOptions{})
	if err != nil {
		return err
	}

	render.New(s.out).GlobalMRs(mrs)
	return nil
}

func (s *searchService) Project(query string) error {
	projects, _, err := s.gitlabSearch.ProjectsByGroup(s.group, query, &gitlab.SearchOptions{})
	if err != nil {
		return err
	}

	render.New(s.out).Projects(projects)
	return nil
}
