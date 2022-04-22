package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version string

var versionCmd = &cobra.Command{
	Use:                   "version",
	Short:                 "Print version number of glcli",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("glcli version " + Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
