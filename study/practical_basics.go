package main

import "fmt"

type Grades map[string]float32

func (grades Grades) initialGrades() {
	grades["Zed"] = 42
	grades["Yasuo"] = 33
	grades["Yummi"] = 22
}

func main() {
	grades := make(Grades)
	grades.initialGrades()

	for k, v := range grades {
		fmt.Println("name: ", k, "grade: ", v)
	}
}
