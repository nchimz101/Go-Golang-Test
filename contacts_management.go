package main

import (
	"fmt"
	"strings"
)

type Contact struct {
	Name  string
	Phone string
	Email string
}

type ContactsManager struct {
	Contacts []Contact
}

func (manager *ContactsManager) AddContact(name, phone, email string) {
	contact := Contact{Name: name, Phone: phone, Email: email}
	manager.Contacts = append(manager.Contacts, contact)
}

func (manager *ContactsManager) ViewContacts() {
	if len(manager.Contacts) == 0 {
		fmt.Println("No contacts available.")
		return
	}
	for i, contact := range manager.Contacts {
		fmt.Printf("Contact %d: %s, Phone: %s, Email: %s\n", i+1, contact.Name, contact.Phone, contact.Email)
	}
}

func (manager *ContactsManager) UpdateContact(index int, name, phone, email string) {
	if index < 0 || index >= len(manager.Contacts) {
		fmt.Println("Invalid contact index.")
		return
	}
	manager.Contacts[index] = Contact{Name: name, Phone: phone, Email: email}
}

func (manager *ContactsManager) DeleteContact(index int) {
	if index < 0 || index >= len(manager.Contacts) {
		fmt.Println("Invalid contact index.")
		return
	}
	manager.Contacts = append(manager.Contacts[:index], manager.Contacts[index+1:]...)
}

func main() {
	manager := &ContactsManager{}

	for {
		fmt.Println("\nContacts Management System:")
		fmt.Println("1. Add Contact")
		fmt.Println("2. View Contacts")
		fmt.Println("3. Update Contact")
		fmt.Println("4. Delete Contact")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var name, phone, email string
			fmt.Print("Enter name: ")
			fmt.Scan(&name)
			fmt.Print("Enter phone: ")
			fmt.Scan(&phone)
			fmt.Print("Enter email: ")
			fmt.Scan(&email)
			manager.AddContact(name, phone, email)

		case 2:
			manager.ViewContacts()

		case 3:
			var index int
			var name, phone, email string
			fmt.Print("Enter the contact index to update: ")
			fmt.Scan(&index)
			fmt.Print("Enter new name: ")
			fmt.Scan(&name)
			fmt.Print("Enter new phone: ")
			fmt.Scan(&phone)
			fmt.Print("Enter new email: ")
			fmt.Scan(&email)
			manager.UpdateContact(index-1, name, phone, email)

		case 4:
			var index int
			fmt.Print("Enter the contact index to delete: ")
			fmt.Scan(&index)
			manager.DeleteContact(index - 1)

		case 5:
			fmt.Println("Exiting program.")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
