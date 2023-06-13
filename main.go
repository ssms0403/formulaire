package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const port = ":8082"

func main() {

	http.HandleFunc("/", handleForm)
	http.HandleFunc("/merci", handlethanks)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Printf("(http://localhost%v) - server started on port %v\n", port, port)
	http.ListenAndServe(port, nil)
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "templates/index.html")
	} else if r.Method == "POST" {

	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
func handlethanks(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/thanks.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	r.ParseForm()
	firstName := r.Form.Get("first-name")
	lastName := r.Form.Get("last-name")
	email := r.Form.Get("email")
	new_password := r.Form.Get("password")
	re := response{
		FirstName: firstName,
		LastName:  lastName,
	}
	// Faites quelque chose avec les valeurs des champs du formulaire
	fmt.Printf("First Name: %s\n", firstName)
	fmt.Printf("Last Name: %s\n", lastName)
	fmt.Printf("Email: %s\n", email)
	fmt.Printf("PASSWORD:%s\n ", new_password)
	t.Execute(w, re)
}

type response struct {
	FirstName string
	LastName  string
}
