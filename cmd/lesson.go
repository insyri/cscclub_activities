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
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			b := color.New(color.Bold)
			for i := range lesson.Master.Lessons {
				L := lesson.Master.Lessons[i]
				fmt.Printf("%-2d %s\n    %s\n    Author: %s\n\n", L.Id, b.Sprint(L.Name), L.Description, L.Author)
			}
			fmt.Println("Use \"csa start [number]\" to start the lesson that corresponds to the number.")
		} else {
			n, err := strconv.Atoi(args[0])
			if err != nil {
				util.LogErrorAndExit(err)
			}
			for i := range lesson.Master.Lessons {
				if lesson.Master.Lessons[i].Id == uint8(n) {
					fmt.Println(lesson.Master.Lessons[i].Name)
				}
			}
		}
	},
}
