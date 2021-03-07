package main

import "fmt"

// Grades have name and id
type Grades map[string]float32

func (grades Grades) initialGrades() {
	grades["Zed"] = 42
	grades["Yasuo"] = 33
	grades["Yummi"] = 22
}

func (grades Grades) showGrades() {
	fmt.Println(grades)
}

func (grades Grades) showPrettyGrades() {
	for k, v := range grades {
		fmt.Println(k, ":", v)
	}
}

func TestMapBasics() {
	grades := make(Grades)
	grades.initialGrades()

	grades.showGrades()
	delete(grades, "Zed")
	grades.showPrettyGrades()
}
