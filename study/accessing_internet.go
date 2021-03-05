package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getTestRestAPI(url string) {
	resp, _ := http.Get(url)
	bytes, _ := ioutil.ReadAll(resp.Body)
	stringBody := string(bytes)
	fmt.Println(stringBody)
	resp.Body.Close()

}

func main() {
	var url string = "https://jsonplaceholder.typicode.com/todos/1"
	getTestRestAPI(url)
}
