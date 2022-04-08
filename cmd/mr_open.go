package cmd

import (
	"github.com/spf13/cobra"
)

var mrOpenCmd = &cobra.Command{
	Use:                   "open [MRID]",
	Short:                 "Open a merge request page in the default browser",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return projectManager().MergeRequest.Open(args[0])
	},
}

func init() {
	mrCmd.AddCommand(mrOpenCmd)
}
