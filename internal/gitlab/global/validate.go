package global

import (
	"io"
	"io/ioutil"

	"github.com/alekseiapa/glcli/internal/gitlab/contract"
	"github.com/alekseiapa/glcli/internal/gitlab/render"
)

type validateService struct {
	gitlabValidate contract.GitlabValidate
	out            io.Writer
}

// Lint validate .gitlab-ci.yml whether valid.
func (s *validateService) Lint(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	result, _, err := s.gitlabValidate.Lint(string(file))
	if err != nil {
		return err
	}

	render.New(s.out).LintCI(result)
	return nil
}
