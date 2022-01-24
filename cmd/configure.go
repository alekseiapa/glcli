package cmd

import (
	"log"
	"os"

	"github.com/alekseiapa/glcli/internal/config"
	"github.com/spf13/cobra"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure glcli CLI options",
	Run: func(cmd *cobra.Command, args []string) {
		if configured {
			return
		}

		if err := config.Configure(os.Stdin, os.Stdout); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
