package main

import (
	"encoding/binary"
	"os"
	"time"
)

// cache file structure:
// first 8 bytes - timestamp      int64
// next byte     - lesson id      uint8
// next byte     - checkpoint id  uint8

type Information struct {
	time         int64
	lessonId     uint8
	checkpointID uint8
}

var fileName = UserHomeDir() + string(os.PathSeparator) + ".csclub_activities"

func openCache() (*os.File, error) {
	// 0666 -> r/w for all users
	return os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0666)
}

func Save(lessonID uint8, checkpointID uint8) error {
	file, err := openCache()
	if err != nil {
		return err
	}
	defer file.Close()

	b := make([]byte, 10)
	binary.LittleEndian.PutUint64(b, uint64(time.Now().Unix()))
	b[8] = lessonID
	b[9] = checkpointID
	_, err = file.Write(b)

	return err
}

func Load() (*Information, error) {
	file, err := openCache()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	b := make([]byte, 10)
	if _, err := file.Read(b); err != nil {
		return &Information{}, err
	}
	return &Information{
		time:         int64(binary.LittleEndian.Uint64(b[:8])),
		lessonId:     b[8],
		checkpointID: b[9],
	}, nil
}
