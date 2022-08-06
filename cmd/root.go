package cmd

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	conf "github.com/alekseiapa/glcli/internal/config"
	"github.com/alekseiapa/glcli/internal/errors"
	"github.com/alekseiapa/glcli/internal/git"
	"github.com/alekseiapa/glcli/internal/gitlab/global"
	"github.com/alekseiapa/glcli/internal/gitlab/group"
	"github.com/alekseiapa/glcli/internal/gitlab/project"
	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:           "glcli",
	Short:         "A CLI tool for gitlab",
	SilenceErrors: true,
	SilenceUsage:  true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if cmd, err := rootCmd.ExecuteC(); err != nil {
		errors.Handle(cmd, err)
	}
}

func init() {
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	rootCmd.SetFlagErrorFunc(func(cmd *cobra.Command, err error) error {
		return &errors.FlagError{Err: err}
	})
	cobra.OnInitialize(initConfig)

	log.SetFlags(0)
}

var (
	config     *conf.Config
	configured = false
)

// initConfig reads in config file.
func initConfig() {
	home, _ := os.UserHomeDir()
	path := filepath.Join(home, ".config")
	config = conf.New(path)

	if err := config.Load(); err != nil {
		if err := config.Init(os.Stdin, os.Stdout); err != nil {
			log.Fatal(err)
		}
		configured = true
	}
}

func gitlabClient() *gitlab.Client {
	client := &http.Client{Timeout: 1000 * time.Second}
	c := gitlab.NewClient(client, config.Get("token"))
	if err := c.SetBaseURL(config.Get("host")); err != nil {
		log.Fatal(err)
		return nil
	}
	return c
}

func projectManager(p *string) *project.Manager {
	if p == nil {
		temp := git.CurrentRepo()
		p = &temp
	}

	return project.NewManager(gitlabClient(), *p, os.Stdout)
}

func groupManager(g string) *group.Manager {
	return group.NewManager(gitlabClient(), g, os.Stdout)
}

func globalManager() *global.Manager {
	return global.NewManager(gitlabClient(), os.Stdout)
}
