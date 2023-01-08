package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ContactDetails struct {
	Email 	string
	Subject string
	Message string
}

func main() {
	tmpl := template.Must(template.ParseFiles("forms.html"))
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := ContactDetails {
			Email: 	 r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		// do something with details
		fmt.Println(details)
		_ = details

		tmpl.Execute(w, struct {
			Success bool
		} {
			Success: true,
		})
	})

	http.ListenAndServe(":5173", nil)
}
