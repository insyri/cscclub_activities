package lesson

import (
	"fmt"
	"github.com/mitchellh/go-wordwrap"
	"os"
)

var lessonOne = &Lesson{
	Author:      "Noah Mercedes",
	Name:        "Introduction to the Terminal + Getting Started",
	Description: "Learn the basics of using the terminal! (todo desc)",
	Checkpoints: []Checkpoint{
		{
			Title: "Getting Started",
			Action: func() {
				printPrompts([]string{
					"Welcome to the activities program! This program was designed to interactively learn computer" +
						" science concepts through the terminal, as well as learning about the terminal itself! This" +
						" activity will go over basic Linux-based terminal concepts.",

					"To do that, we'll first need to ensure that you're running a shell-based command-line interpreter.",

					"Windows users should consider installing a version of MSYS2; the easiest way to get started" +
						" with this is to use the Git Bash program from the GitForWindows bundle, which installs" +
						" a version of MSYS2 which comes with an alternate terminal program (Git Bash) and Git" +
						" preinstalled:\nhttps://git-scm.com/download/win",

					"If you are running a non-windows based operating system, chances are that you are already using" +
						" a shell-based command-line interpreter.",

					"To move on to the next checkpoint, simply re-run this program: `csa`",

					"If you come back to the same checkpoint, it means that you haven't met the conditions to move" +
						" on to the next checkpoint. Ensure that you:",

					"    1. Have a version of Bash installed on your system\n" +
						"    2. Are running and using Bash at this moment" +
						"    3. Are executing this program (csa) within Bash",
				})
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
		{
			Title: "Introduction to the Terminal",
			Action: func() {
				fmt.Println(wordwrap.WrapString(
					"Awesome! If you got here, you're properly set up.", WordWrapLimit))
			},
		},
	},
}
