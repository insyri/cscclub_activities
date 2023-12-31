package lesson

import (
	"csclub-activities/util"
	"fmt"
	"github.com/fatih/color"
	"github.com/mitchellh/go-wordwrap"
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

func (m *master) Run(lesson int, checkpoint int) error {

	if m.Lessons == nil ||
		lesson >= len(m.Lessons) {
		return util.NonexistentLesson
	}

	if m.Lessons[lesson].Checkpoints == nil ||
		checkpoint >= len(m.Lessons[lesson].Checkpoints) {
		return util.NonexistentCheckpoint
	}

	cps := m.Lessons[lesson].Checkpoints
	// TODO: remove this for loop for a single run conditional statement
	for i := checkpoint; i < len(cps); i++ {
		if cps[i].ShouldPromote() {
			if i+1 < len(cps) {
				break
			}
			i++
			continue
		}

		fmt.Printf("%-5s  %s\n\n",
			color.New(color.Bold).Add(color.FgGreen).Sprintf("%2d/%-2d", i+1, len(cps)),
			color.New(color.Bold).Sprintf("%s", cps[i].Title),
		)
		cps[i].Action()
		fmt.Println("\n    Note: to get back to the help screen, run: `csa help`")

		err := util.Save(uint8(lesson), uint8(checkpoint))
		if err != nil {
			return err
		}
		break
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

func printPrompts(prompts []string) {
	for x := range prompts {
		if x == 0 {
			fmt.Println(wordwrap.WrapString(prompts[x], WordWrapLimit))
		} else {
			fmt.Println("\n" + wordwrap.WrapString(prompts[x], WordWrapLimit))
		}
	}

}
