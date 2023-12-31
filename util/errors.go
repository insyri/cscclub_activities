package util

import "errors"

var (
	NonexistentLesson     = errors.New("impossible index: nonexistent lesson")
	NonexistentCheckpoint = errors.New("impossible index: nonexistent checkpoint")
	ErrReservedIDs        = errors.New("reserved lessonID or checkpointID value: cannot be 0")
)
