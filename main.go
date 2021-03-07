package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func HtmlTemplateHandler(w http.ResponseWriter, r *http.Request) {
	p := Page{Title: "first page", Msg: "Hello world"}
	t, _ := template.ParseFiles("index.html")
	err := t.Execute(w, p)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	// 	fmt.Println(contactJSON)
	// 	http.HandleFunc("/", HtmlTemplateHandler)
	// 	http.ListenAndServe(":8000", nil)

	contact := new(Contacts)
	contact.initialContact(contactJSON)
	contact.showContacts()
}
