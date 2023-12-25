package main

import "fmt"

func main() {
	// commands during lesson
	// stop -> clears checkpoint
	// previous -> sets checkpoint to previous one + does thing
	// help ->

	// active checkpoint? || stop || previous

	info, err := Load()
	if err != nil {
		panic(err)
	}

	fmt.Println(info)

	err = Save(info.lessonId+1, info.checkpointID+1)
	if err != nil {
		panic(err)
	}

	info, err = Load()
	if err != nil {
		panic(err)
	}

	fmt.Println(info)

}
