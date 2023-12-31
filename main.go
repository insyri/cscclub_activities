package main

import (
	"csclub-activities/cmd"
	"csclub-activities/lesson"
	"csclub-activities/util"
)

func main() {

	// util.Load could return an empty or new information struct.
	// new does not help us.

	information, err := util.Load()
	if err != nil {
		util.LogErrorAndExit(err)
	}

	err = lesson.Master.Run(int(information.LessonId), int(information.CheckpointID))
	if err != nil {
		util.LogErrorAndExit(err)
	} else {
		cmd.Execute()
	}

	// commands during lesson
	// stop -> clears checkpoint
	// previous -> sets checkpoint to previous one + does thing
	// help ->

	// active checkpoint? || stop || previous

}
