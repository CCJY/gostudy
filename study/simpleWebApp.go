package main

import (
	"fmt"
	"net/http"
)

func index_handler(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "Whoa, Go is neat!")
}
func about_handler(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "about!")
}
func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about", about_handler)
	http.ListenAndServe(":8000", nil)
}
