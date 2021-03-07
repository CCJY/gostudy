package main

import (
	"encoding/json"
	"fmt"
)

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

func (m MapContact) deleteByID(id int) {
	delete(m, id)
}
