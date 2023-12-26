package lesson

func init() {
	err := Master.append(lessonTwo)
	if err != nil {
		panic(err)
	}
}

var lessonTwo = &Lesson{
	Id:          2,
	Author:      "Noah Mercedes",
	Name:        "Linux Filesystem Navigation and Manipulation",
	Description: "Let's ditch file explorer and become hackers* B)",
	Checkpoints: nil,
}
