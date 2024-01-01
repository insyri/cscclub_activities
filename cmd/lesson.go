package cmd

import (
	"csclub-activities/lesson"
	"csclub-activities/util"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"strconv"
)

func init() {
	rootCmd.AddCommand(lessonCmd)
}

var lessonCmd = &cobra.Command{
	Use:     "lesson",
	Short:   "Starts the activity / Lists available activities",
	Aliases: []string{"lessons", "list", "start"},
	//Args: func(cmd *cobra.Command, args []string) error {
	//	if err := cobra.RangeArgs(0, 2)(cmd, args); err != nil {
	//		return err
	//	}
	//	// Run the custom validation logic
	//	if myapp.IsValidColor(args[0]) {
	//		return nil
	//	}
	//	return fmt.Errorf("invalid color specified: %s", args[0])
	//},
	// TODO: impl arg validator + pre-run (if l/c exists)
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			for i := 1; i < len(lesson.Master.Lessons); i++ {
				L := lesson.Master.Lessons[i]
				fmt.Printf("%-2d  %s\n    %s\n    Author: %s"+
					"\n\n",
					i, color.New(color.Bold).Sprint(L.Name), L.Description, L.Author)
			}
			fmt.Println("Use \"csa start [number]\" to start the lesson that corresponds to the number.")
		} else {
			n, err := strconv.Atoi(args[0])
			if err != nil {
				util.LogErrorAndExit(err)
			}
			if n > len(lesson.Master.Lessons) || n == 0 {
				util.LogErrorAndExit(util.NonexistentLesson)
			}
			err = lesson.Master.Run(n, 0, true)
			if err != nil {
				util.LogErrorAndExit(err)
			}
		}
	},
}
