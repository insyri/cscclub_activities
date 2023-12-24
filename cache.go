package main

import (
	"encoding/binary"
	"os"
	"time"
)

var fileName = UserHomeDir() + string(os.PathSeparator) + ".csclub_activities"

func SaveTime(file *os.File) error {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(time.Now().Unix()))
	_, err := file.Write(b)
	return err
}

func GetTime(file *os.File) (int64, error) {
	b := make([]byte, 8)
	if _, err := file.Read(b); err != nil {
		return 0, err
	}
	return int64(binary.LittleEndian.Uint64(b)), nil
}

func GetFile() (*os.File, error) {
	// 0666 -> r/w to all
	return os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0666)
}
