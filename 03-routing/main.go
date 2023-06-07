package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// home
	r.HandleFunc("/", handleHomeRoute).Methods("GET")

	bookrouter := r.PathPrefix("/books").Subrouter()
	// GET books
	bookrouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	}).Methods("GET")
	// /books/go-programming-blueprint
	bookrouter.HandleFunc("/{title}", GetBook).Methods("GET")

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	http.ListenAndServe(":80", r)
}

func handleHomeRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome</h1>"))
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	// get the book
	// navigate to the page
	vars := mux.Vars(r)
	t := vars["title"] // the book title slug
	fmt.Fprintf(w, "Getting the book with title: %s\n", t)
}
