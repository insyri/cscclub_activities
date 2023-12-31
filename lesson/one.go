package lesson

import (
	"csclub-activities/util"
	"fmt"
	"github.com/mitchellh/go-wordwrap"
	"os"
	"runtime"
)

var lessonOne = &Lesson{
	Author:      "Noah Mercedes",
	Name:        "Introduction to the Terminal + Getting Started",
	Description: "Learn the basics of using the terminal! (todo desc)",
	Checkpoints: []Checkpoint{
		{
			Title: "Getting Started",
			Action: func(lessonID uint8, checkpointID uint8) {
				defer func() {
					err := util.Save(lessonID, checkpointID)
					if err != nil {
						util.LogErrorAndExit(err)
					}
				}()

				// TODO: find or develop a word wrapper.

				fmt.Println(wordwrap.WrapString(
					"Welcome to the activities program! This program was designed to interactively learn concepts"+
						" within the terminal, as well as learning the terminal! To get started, first we'll learn basic"+
						" Linux-based terminal concepts through the shell program.", util.WordWrapLimit))

				fmt.Println("\nFirst, we'll ensure that you're running a shell-based command-line interpreter.")
				if runtime.GOOS == "windows" {
					fmt.Println(wordwrap.WrapString(
						"\nWindows users should consider installing a version of MSYS2, the easiest way to get started"+
							" with this is to use the Git Bash program from GitForWindows:\n"+
							"https://git-scm.com/download/win", util.WordWrapLimit))
					fmt.Println()
				}
				fmt.Println(wordwrap.WrapString(
					"If you are running a non-windows based operating system, chances are that Git is already"+
						" installed; to verify, enter \"git\" into your shell and you should see a lengthy dialog.",
					util.WordWrapLimit))
				fmt.Println()
				fmt.Println(
					"After installing Git, or to move on to the next checkpoint, simply re-run this program: `csa`")

			},
			ShouldPromote: func() bool {
				switch os.Getenv("SHELL") {
				case "/usr/bin/bash", "/usr/bin/zsh", "/usr/bin/sh":
					return true
				}
				// for those who run other shells like fish, you could just use the env VALIDSHELL
				return os.Getenv("VALIDSHELL") != ""
			},
		},
	},
}
