package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var homeTemplate *template.Template

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "To get in touch, please send an email "+
		"to <a href=\"mailto:support@lenslocked.com\">"+
		"support@lenslocked.com</a>.")
}

func faq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the FAQ page.")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Sorry this page could not be found.")
}

func main() {
	var err error
	homeTemplate, err = template.ParseFiles("views/home.gohtml")
	if err != nil {
		panic(err)
	}

	var h http.Handler = http.HandlerFunc(notFound)
	r := mux.NewRouter()
	r.NotFoundHandler = h

	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", contact)
	http.ListenAndServe(":3000", r)
}
