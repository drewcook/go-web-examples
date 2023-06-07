// Using bcrypt to hash passwords from user input
package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// Takes in a given password string and hashes it, returns the hashed password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Validates that a hash matches for it's given password, returns bool
func ValidateHash(hash string, password string) bool {
	// Needs to take in bytes arrays
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	// Hash the password
	password := "myS3cr3TP455w0rD!"
	hash, err := HashPassword(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Password: ", password)
	fmt.Println("Hash: ", hash)

	// Validate the hash matches
	isValid := ValidateHash(hash, password)
	fmt.Println("Hash is valid: ", isValid)
}
