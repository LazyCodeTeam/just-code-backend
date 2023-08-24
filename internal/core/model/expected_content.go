package model

type ExpectedTechnology struct {
	Technology       Technology
	ExpectedSections []ExpectedSection
}

type ExpectedSection struct {
	Section       Section
	ExpectedTasks []Task
}
