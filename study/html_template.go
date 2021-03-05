package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Page is ...
type Page struct {
	Title string
	Msg   string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	p := Page{Title: "first page", Msg: "Hello world"}
	t, _ := template.ParseFiles("index.html")
	err := t.Execute(w, p)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)

}
