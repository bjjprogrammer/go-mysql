package handlers

import (
	"database/sql"
	"log"

	"github.com/nigerdyanes/go-mysql/models"
)

func ListContacts(db *sql.DB) {
	// List all contacts
	rows, err := db.Query("SELECT * FROM contacts")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		contact := models.Contact{}
		var emailNullValue sql.NullString

		err := rows.Scan(&contact.Id, &contact.Name, &emailNullValue, &contact.Phone)

		if emailNullValue.Valid {
			contact.Email = emailNullValue.String
		} else {
			contact.Email = "WITHOUT EMAIL"
		}

		if err != nil {
			log.Fatal(err)
		}

		log.Println(contact)
	}
}

func GetContactByID(db *sql.DB, id int) {
	// Get a contact by ID
	row := db.QueryRow("SELECT * FROM contacts WHERE id = ?", id)

	contact := models.Contact{}
	var emailNullValue sql.NullString

	err := row.Scan(&contact.Id, &contact.Name, &emailNullValue, &contact.Phone)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("The contact with ID %d was not found", id)
		}
	}

	if emailNullValue.Valid {
		contact.Email = emailNullValue.String
	} else {
		contact.Email = "WITHOUT EMAIL"
	}

	log.Println(contact)
}

func CreateContact(db *sql.DB, contact models.Contact) {
	query := "INSERT INTO contacts (name, email, phone) VALUES (?, ?, ?)"

	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Contact created successfully")
}

func UpdateContact(db *sql.DB, id int, contact models.Contact) {
	query := "UPDATE contacts SET name = ?, email = ?, phone = ? WHERE id = ?"

	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone, id)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Contact updated successfully")
}

func DeleteContact(db *sql.DB, id int) {
	query := "DELETE FROM contacts WHERE id = ?"

	_, err := db.Exec(query, id)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Contact deleted successfully")
}
