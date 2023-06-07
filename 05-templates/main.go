package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	// Use template from local file
	tmpl := template.Must(template.ParseFiles("layout.html"))

	// define handler for root
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Construct the interface to pass along to the template engine
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})

	// start server
	http.ListenAndServe(":80", nil)
}
