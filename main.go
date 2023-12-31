package main

import (
	"csclub-activities/cmd"
	"csclub-activities/lesson"
	"csclub-activities/util"
	"errors"
	"io"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		cmd.Execute()
		return
	}

	information, err := util.Load()
	if err != nil {
		if errors.Is(err, io.EOF) { // no file created

			err = util.NewSave()
			if err != nil {
				util.LogErrorAndExit(err)
			}
			cmd.Execute()
			return

		} else {
			util.LogErrorAndExit(err)
		}
	}

	if information.LessonId == util.NoOngoingActivity {
		cmd.Execute()
		return
	}

	err = lesson.Master.Run(int(information.LessonId), int(information.CheckpointID))
	if err != nil {
		util.LogErrorAndExit(err)
	}

	// commands during lesson
	// stop -> clears checkpoint
	// previous -> sets checkpoint to previous one + does thing
	// help ->

	// active checkpoint? || stop || previous

}
