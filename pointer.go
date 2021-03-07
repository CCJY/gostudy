package main

import "fmt"

func referenceTest() *int {
	var a = 1
	return &a
}

func TestPointer() {
	x := 15
	a := &x
	fmt.Println(a)
	fmt.Println(*a)

	*a = 5
	fmt.Println(a)
	fmt.Println(*a)

	var vv *int = referenceTest()
	fmt.Println(*vv)

}
