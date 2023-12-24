package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		// compare

		f, err := GetFile()
		if err != nil {
			panic(err)
		}

		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				panic(err)
			}
		}(f)

		x, err := GetTime(f)
		if err != nil {
			panic(err)
		}

		then := time.Unix(x, 0)
		result := time.Now().Sub(then)
		fmt.Printf("%.3f", result.Seconds())
	} else {
		f, err := GetFile()
		if err != nil {
			panic(err)
		}

		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				panic(err)
			}
		}(f)

		err = SaveTime(f)
		if err != nil {
			panic(err)
		}

		fmt.Println("Done!")
	}
}
