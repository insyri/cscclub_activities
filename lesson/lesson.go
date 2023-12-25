package lesson

import "time"

/*
save checkpoint
check checkpoint (specific lesson)

metadata:
	name
	amount of steps
	description

*/

type Lesson struct {
	id            uint8
	name          string
	author        string
	datePublished time.Time
	description   string
	checkpoints   []Checkpoint
}

type Checkpoint struct {
	preCondition func() bool // precondition could be previous checkpoint condition
	condition    func() bool
	action       func()
	position     uint8
}
