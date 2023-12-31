package cmd

import (
	"csclub-activities/util"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(readCmd)
}

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Clears the lesson checkpoint",
	RunE: func(cmd *cobra.Command, args []string) error {
		info, err := util.Load()
		if err != nil {
			return err
		}

		fmt.Printf(
			"Timestamp: %d\n"+
				"LessonID:  %d\n"+
				"CheckptID:  %d\n", info.Time, info.LessonId, info.CheckpointID)
		return nil
	},
}
