package cmd

import (
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Init or list glcli CLI options",
}

func init() {
	rootCmd.AddCommand(configCmd)
}
