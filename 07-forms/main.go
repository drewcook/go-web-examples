package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {
	tmpl := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Guard, only POSTS allowed
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		// extract form values and transform data
		details := ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		// do something with details
		fmt.Printf("Got contact details: %#v\n", details)

		tmpl.Execute(w, struct{ Success bool }{true})
	})

	http.ListenAndServe(":80", nil)
}
