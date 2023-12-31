package lesson

import (
	"csclub-activities/util"
	"fmt"
	"github.com/fatih/color"
)

/*
save checkpoint
check checkpoint (specific lesson)

metadata:
	name
	amount of steps
	description

*/

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

func (m *master) Run(lesson int, checkpoint int) error {

	if m.Lessons == nil ||
		lesson >= len(m.Lessons) {
		return util.NonexistentLesson
	}

	if m.Lessons[lesson].Checkpoints == nil ||
		checkpoint >= len(m.Lessons[lesson].Checkpoints) {
		return util.NonexistentCheckpoint
	}

	for i := checkpoint; i < len(m.Lessons[lesson].Checkpoints); i++ {

		cp := m.Lessons[lesson].Checkpoints[i]
		fmt.Printf("%-5s  %s\n\n",
			color.New(color.Bold).Add(color.FgGreen).Sprintf("%-2d/%-2d", i, len(m.Lessons)),
			color.New(color.Bold).Sprintf("%s", cp.Title),
		)
		cp.Action(uint8(lesson), uint8(checkpoint))
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
	Action func(lessonID uint8, checkpointID uint8)
	// What is the condition that allows the user to go onto the next checkpoint?
	ShouldPromote  func() bool
	ExpectedErrors []ExpectantError
}

type ExpectantError struct {
	MatchesError func() bool
	Feedback     string
}
