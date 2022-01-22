package cmd

import (
	"os"

	"github.com/alekseiapa/glcli/internal/config"
	"github.com/spf13/cobra"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure glcli CLI options",
	Run: func(cmd *cobra.Command, args []string) {
		config.Configure(os.Stdin, os.Stdout)
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
