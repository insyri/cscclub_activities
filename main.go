package main

import "csclub-activities/cmd"

func main() {

	cmd.Execute()

	// Check if a checkpoint is active
	//info, err := Load()
	//if err != nil {
	//	if errors.Is(err, io.EOF) {
	//		err = NewSave()
	//		if err != nil {
	//			panic(err)
	//		}
	//	} else {
	//		panic(err)
	//	}
	//}

	// LF checkpoint

	// commands during lesson
	// stop -> clears checkpoint
	// previous -> sets checkpoint to previous one + does thing
	// help ->

	// active checkpoint? || stop || previous

}
