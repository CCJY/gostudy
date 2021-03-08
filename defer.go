package main

import "fmt"

func deferDo() {
	defer fmt.Println("Done")
	fmt.Println()
}
