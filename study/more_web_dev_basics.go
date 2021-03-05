package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<h1>HiHIHIHI</h1>
	<p>go is?</p>
	<p>You %s even add %s</p>
	`, "can", "<strong>variables</strong>")
	fmt.Fprintf(w, "<h1>HiHIHIHI</h1>")
	fmt.Fprintf(w, "<p>go is?</p>")
	fmt.Fprintf(w, "<p>You %s even add %s</p>", "can", "<strong>variables</strong>")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}
