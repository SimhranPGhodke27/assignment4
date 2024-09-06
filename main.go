package main

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// Define a model
type User struct {
	gorm.Model
	Name        string
	Email       string
	PhoneNumber string
	Age         int
	Address     string
}

func main() {
	// Connect to the database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{})

	// Create
	user := User{
		Name:        "John Doe",
		Email:       "john.doe@example.com",
		PhoneNumber: "123-456-7890",
		Age:         30,
		Address:     "123 Elm Street",
	}
	db.Create(&user)
	fmt.Println("Created user:", user)

	// Read
	var readUser User
	db.First(&readUser, 1) // Find user with ID 1
	fmt.Println("Read user:", readUser)

	// Update
	db.Model(&readUser).Updates(User{
		Email:       "john.doe@newdomain.com",
		PhoneNumber: "987-654-3210",
		Age:         31,
		Address:     "456 Oak Street",
	})
	fmt.Println("Updated user:", readUser)

	// Delete
	db.Delete(&readUser, 1)
	fmt.Println("Deleted user:", readUser)
}
