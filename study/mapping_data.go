package main

import (
	"encoding/json"
	"fmt"
)

var contactJSON = []byte(`
[
	{
	"userId": 1,
	"name": "Zed"
	},
	{
	"userId": 2,
	"name": "Yasuo"
	}
]
`)

// Contact is ..
type Contact struct {
	UserID int    `json:"userId"`
	Name   string `json:"name"`
}

// MapContact is ..
type MapContact map[int]Contact

// Contacts are .....
type Contacts []Contact

func (contacts *Contacts) initialContact(bytes []byte) {
	err := json.Unmarshal(bytes, &contacts)
	if err != nil {
		fmt.Println("err:", err.Error())
	}
}

func (contacts Contacts) showContacts() {
	for _, contact := range contacts {
		fmt.Println(contact.UserID, ":", contact.Name)
	}
}

func (contacts Contacts) toMap() MapContact {
	mapContact := make(MapContact)
	for _, contact := range contacts {
		mapContact[contact.UserID] = contact
	}
	fmt.Println(mapContact)
	return mapContact
}

func (m MapContact) showContacts() {
	for k, v := range m {
		fmt.Println("key=", k, " value=", v)
	}
}

func main() {
	contacts := new(Contacts)
	// var contacts Contacts
	contacts.initialContact(contactJSON)
	contacts.showContacts()

	contacts.toMap().showContacts()

}
