package global

import (
	"errors"
	"os"
	"testing"

	"github.com/alekseiapa/glcli/internal/gitlab/contract/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/xanzy/go-gitlab"
)

func TestLint(t *testing.T) {
	t.Run("file not exists", func(t *testing.T) {
		s := &validateService{}
		err := s.Lint("foo/bar")

		assert.Error(t, err)
	})

	t.Run("lint", func(t *testing.T) {
		mockGitlabValidate := &mocks.GitlabValidate{}
		mockGitlabValidate.On("Lint", "HelloWord\n").
			Once().
			Return(&gitlab.LintResult{}, &gitlab.Response{}, errors.New(""))

		s := &validateService{gitlabValidate: mockGitlabValidate}
		dir, _ := os.Getwd()
		s.Lint(dir + "/../../../test/.gitlab-ci.yml")

		mockGitlabValidate.AssertExpectations(t)
	})
}
