package main

import (
	"fmt"
	"net/http"
)

func SimpleWebAppIndexhandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "Whoa, Go is neat!")
}
func SimpleWebAppAboutandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "about!")
}
func TestSimpleWeb() {
	http.HandleFunc("/", SimpleWebAppIndexhandler)
	http.HandleFunc("/about", SimpleWebAppAboutandler)
	http.ListenAndServe(":8000", nil)
}
