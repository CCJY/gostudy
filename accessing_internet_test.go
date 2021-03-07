package main

import (
	"fmt"
)

func TestAccessingInternet() {
	// var url string = "https://jsonplaceholder.typicode.com/todos/1"
	// var bytes []byte = getBodyBytesFromAPI(url)
	// t, err := todoFromBytes(bytes)

	t, err := todoFromBytes(todoJSON)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(t.toJSONString())
	fmt.Println(t.toString())
}
