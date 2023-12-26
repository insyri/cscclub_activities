package cmd

import (
	"csclub-activities/util"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	readCmd.Hidden = true
	rootCmd.AddCommand(readCmd)
}

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Clears the lesson checkpoint",
	Run: func(cmd *cobra.Command, args []string) {
		info, err := util.Load()
		if err != nil {
			util.LogErrorAndExit(err)
		}

		fmt.Printf(
			"Timestamp: %d\n"+
				"LessonID:  %d\n"+
				"CheckptID:  %d\n", info.Time, info.LessonId, info.CheckpointID)
	},
}
