package main

import (
	"fmt"
	"net/http"

	"lenslocked.com/views"

	"github.com/gorilla/mux"
)

var (
	homeView    *views.View
	contactView *views.View
	faqView     *views.View
	signupView  *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(faqView.Render(w, nil))
}

func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(signupView.Render(w, nil))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Sorry this page could not be found.")
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	homeView = views.NewView("bootstrap",
		"views/home.gohtml")
	contactView = views.NewView("bootstrap",
		"views/contact.gohtml")
	faqView = views.NewView("bootstrap",
		"views/faq.gohtml")
	signupView = views.NewView("bootstrap",
		"views/signup.gohtml")

	var h http.Handler = http.HandlerFunc(notFound)
	r := mux.NewRouter()
	r.NotFoundHandler = h

	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.HandleFunc("/signup", signup)
	http.ListenAndServe(":3000", r)
}
