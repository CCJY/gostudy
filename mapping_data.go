package main

func TestMappingData() {
	contacts := new(Contacts)
	// var contacts Contacts
	contacts.initialContact(contactJSON)
	contacts.showContacts()

	contacts.toMap().showContacts()

	m := contacts.toMap()
	m.deleteByID(1)
	m.showContacts()
}
