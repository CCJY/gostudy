package main

import (
	"net/http"
	"sync"
	"text/template"
)

var hwg sync.WaitGroup

type ContactsPage struct {
	Title    string
	Msg      string
	Contacts Contacts
}

func contactTableHandler(w http.ResponseWriter, r *http.Request) {
	var contacts Contacts
	contacts.initialContact(contactJSON)
	p := ContactsPage{Title: "Example Page", Msg: "This is example", Contacts: contacts}

	t, _ := template.ParseFiles("index.html")

	t.Execute(w, p)

}

func contactWebRun() {

	http.HandleFunc("/", contactTableHandler)
	http.ListenAndServe(":8000", nil)

}
