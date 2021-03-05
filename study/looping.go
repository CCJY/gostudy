package main

import "fmt"

func main() {
	loop := 2
	i := 0
	for i < loop {
		fmt.Println(i)
		i++
	}

	for x := 0; x < loop; x++ {
		fmt.Println(x)
	}

	y := 5
	for {
		fmt.Println("Hello", y)
		y += 3
		if 10 < y {
			break
		}
	}

	var datas []int = []int{0, 1, 2, 3, 4, 5}

	for data := range datas {
		fmt.Println(data)
	}

}
