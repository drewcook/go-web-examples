// Encoding and decoding JSON & structs
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Our struct with json identifiers
type User struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Age       int    `json:"age"`
}

// Decode request body into schema for a User struct
func decode(w http.ResponseWriter, r *http.Request) {
	// Set an empty var to decode into
	var user User
	// Decode following the schema for user
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Fatal(err)
	}
	// Use the values on 'user' var
	fmt.Fprintf(w, "%s %s is %d years old!", user.Firstname, user.Lastname, user.Age)
}

// Encode struct data into json
func encode(w http.ResponseWriter, r *http.Request) {
	// Work with our struct (payload)
	peter := User{
		Firstname: "Peter",
		Lastname:  "Cottontail",
		Age:       204,
	}
	// Encode it into JSON
	if err := json.NewEncoder(w).Encode(peter); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "%v", peter)
}

func main() {
	http.HandleFunc("/decode", decode)
	http.HandleFunc("/encode", encode)
	http.ListenAndServe(":80", nil)
}
