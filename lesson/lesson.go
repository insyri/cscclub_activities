package lesson

import "errors"

/*
save checkpoint
check checkpoint (specific lesson)

metadata:
	name
	amount of steps
	description

*/

var Master master

type master struct {
	Lessons []Lesson
}

func (m *master) append(newLesson *Lesson) error {
	for i := range m.Lessons {
		if m.Lessons[i].Id == newLesson.Id {
			// here for future contributors
			return errors.New("duplicate lesson ids")
		}
	}
	m.Lessons = append(m.Lessons, *newLesson)
	return nil
}

type Lesson struct {
	Id          uint8
	Author      string
	Name        string
	Description string
	Checkpoints []Checkpoint
}

type Checkpoint struct {
	PreCondition func() bool // precondition could be previous checkpoint condition
	Condition    func() bool
	Action       func()
	Position     uint8
}
