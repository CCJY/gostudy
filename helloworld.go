package main

import (
	"fmt"
	"math"
	"math/rand"
)

func foo() {
	fmt.Println("The square root of 4 is", math.Sqrt(4))
}

func bar() {
	fmt.Println("A number from 1-100", rand.Intn(100))
}
func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}
func multiple(a, b string) (string, string) {
	return a, b
}

func TestHelloworld() {
	foo()
	bar()
	var a, b int = 10, 11
	var aa int = add(a, b)
	fmt.Println(aa)
	var bb int = add2(a, b)
	fmt.Println(bb)

	w1, w2 := "Hey", "There"
	fmt.Println(multiple(w1, w2))

	x := aa

	fmt.Println(x)
}
