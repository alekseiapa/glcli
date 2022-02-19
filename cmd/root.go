package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/alekseiapa/glcli/internal/config"
	"github.com/alekseiapa/glcli/internal/gitlab"
	"github.com/alekseiapa/glcli/internal/utils"
	"github.com/spf13/cobra"
	"github.com/tcnksm/go-gitconfig"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:           "glcli",
	Short:         "A CLI tool for gitlab",
	SilenceUsage:  true,
	SilenceErrors: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	cobra.OnInitialize(initConfig)
}

var (
	c          *config.Config
	configured = false
)

// initConfig reads in config file.
func initConfig() {
	home, _ := os.UserHomeDir()
	path := filepath.Join(home, ".config")
	c = config.New(path)

	if err := c.Load(); err != nil {
		if err := c.Edit(os.Stdin, os.Stdout); err != nil {
			log.Fatal(err)
		}
		configured = true
	}
}

func currentRepo() string {
	url, err := gitconfig.OriginURL()
	if err != nil {
		log.Fatal("not a git repository")
	}

	return utils.ParseRepo(url)
}

func gitlabManager() *gitlab.Manager {
	m, err := gitlab.NewManager(c.Get("host"), c.Get("token"), os.Stdout)
	if err != nil {
		log.Fatal(err)
	}

	return m
}
