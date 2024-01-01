package lesson

import (
	"csclub-activities/util"
	"fmt"
	"github.com/fatih/color"
)

const WordWrapLimit = 100

func init() {
	// Empty first struct to support the NoOngoingActivity functionality
	l := []*Lesson{{}, lessonOne, lessonTwo}
	for x := range l {
		Master.Lessons = append(Master.Lessons, *l[x])
	}
}

var Master master

type master struct {
	Lessons []Lesson
}

// on program startup, it checks whether there was a previously saved checkpoint.
// if there was, it checks if the condition to pass that checkpoint is valid. (ShouldPromote)
// if true, then the next checkpoint is displayed. otherwise, the same checkpoint is displayed.

// Run
// setManually describes whether the checkpoint was loaded from a file or not.
// if it was not loaded (set manually), we shouldn't skip the current checkpoint
// (in other words, don't check the current checkpoint's ShouldPromote)
func (m *master) Run(lesson int, checkpoint int, setManually bool) error {

	if m.Lessons == nil ||
		lesson >= len(m.Lessons) {
		return util.NonexistentLesson
	}

	if m.Lessons[lesson].Checkpoints == nil ||
		checkpoint >= len(m.Lessons[lesson].Checkpoints) {
		return util.NonexistentCheckpoint
	}

	cps := m.Lessons[lesson].Checkpoints
	if !setManually && cps[checkpoint].ShouldPromote() && util.HasNext(checkpoint, cps) {
		checkpoint++
	}

	fmt.Printf("%-5s  %s\n\n",
		color.New(color.Bold).Add(color.FgGreen).Sprintf("%2d/%-2d", checkpoint+1, len(cps)),
		color.New(color.Bold).Sprintf("%s", cps[checkpoint].Title),
	)
	cps[checkpoint].Action()
	fmt.Println("\n    Note: to get back to the help screen, run: `csa help`")

	err := util.Save(uint8(lesson), uint8(checkpoint))
	if err != nil {
		return err
	}
	return nil
}

type Lesson struct {
	Author      string
	Name        string
	Description string
	Checkpoints []Checkpoint
}

type Checkpoint struct {
	Title  string
	Action func()
	// What is the condition that allows the user to go onto the next checkpoint?
	ShouldPromote  func() bool
	ExpectedErrors []ExpectantError
}

type ExpectantError struct {
	MatchesError func() bool
	Feedback     string
}
