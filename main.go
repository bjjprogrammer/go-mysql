package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nigerdyanes/go-mysql/database"
	"github.com/nigerdyanes/go-mysql/handlers"
	"github.com/nigerdyanes/go-mysql/models"
)

func main() {
	db, err := database.Connect()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	for {
		fmt.Println("1. List all contacts")
		fmt.Println("2. Get a contact by ID")
		fmt.Println("3. Create a new contact")
		fmt.Println("4. Update a contact")
		fmt.Println("5. Delete a contact")
		fmt.Println("6. Exit")

		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			handlers.ListContacts(db)
		case 2:
			var id int
			fmt.Print("Enter the ID of the contact: ")
			fmt.Scanln(&id)
			handlers.GetContactByID(db, id)
		case 3:
			contact := readContact()
			handlers.CreateContact(db, contact)
		case 4:
			var id int
			fmt.Print("Enter the ID of the contact: ")
			fmt.Scanln(&id)
			contact := readContact()
			handlers.UpdateContact(db, id, contact)
		case 5:
			var id int
			fmt.Print("Enter the ID of the contact: ")
			fmt.Scanln(&id)
			handlers.DeleteContact(db, id)
		case 6:
			fmt.Println("Bye!")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func readContact() models.Contact {
	reader := bufio.NewReader(os.Stdin)
	var contact models.Contact

	fmt.Print("Enter the name of the contact: ")
	name, _ := reader.ReadString('\n')
	contact.Name = strings.TrimSpace(name)

	fmt.Print("Enter the email of the contact: ")
	email, _ := reader.ReadString('\n')
	contact.Email = strings.TrimSpace(email)

	fmt.Print("Enter the phone of the contact: ")
	phone, _ := reader.ReadString('\n')
	contact.Phone = strings.TrimSpace(phone)

	return contact
}
