package lesson

func init() {
	err := Master.append(lessonOne)
	if err != nil {
		panic(err)
	}
}

var lessonOne = &Lesson{
	Id:          1,
	Author:      "Noah Mercedes",
	Name:        "Introduction to the Terminal + Getting Started",
	Description: "Learn the basics of using the terminal! (todo desc)",
	Checkpoints: nil,
}
