package cmd

import (
	"errors"
	"fmt"
	"strconv"

	errs "github.com/alekseiapa/glcli/internal/errors"
	"github.com/spf13/cobra"
)

var ownMrShowCmd = &cobra.Command{
	Use:                   "show [PID] [MRID]",
	Short:                 "Show information about a merge request",
	DisableFlagsInUseLine: true,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return &errs.FlagError{Err: errors.New("missing MRID")}
		}
		if _, err := strconv.Atoi(args[0]); err != nil {
			return &errs.FlagError{Err: fmt.Errorf("invalid MRID %q", args[0])}
		}
		if len(args) < 2 {
			return &errs.FlagError{Err: errors.New("missing PID")}
		}
		if _, err := strconv.Atoi(args[1]); err != nil {
			return &errs.FlagError{Err: fmt.Errorf("invalid PID %q", args[0])}
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		pID := args[0]
		mrID, _ := strconv.Atoi(args[1])
		return globalManager().MergeRequest.Show(pID, mrID)
	},
}
