package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1> Home </h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `
						<h1>Contact info:</h1>
						email: <a href="mailto:support@mail.com">Support</a>
	`)
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `
		<h1>FAQ(s)</h1>
		<h3> How much do you charge ?</h3>
		<p> $160/hour </p>
	`)
}

func pageNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `<h1>You are lost</h1>`)
}

func main() {
	// Creates a new router
	r := mux.NewRouter()

	// Assign functions to handle different paths
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)

	// Converts pageNotFound func to type HandlerFunc
	notFound := http.HandlerFunc(pageNotFound)
	r.NotFoundHandler = notFound

	// Start up the server
	if err := http.ListenAndServe(":9000", r); err != nil {
		fmt.Println("Error occurred while binding to port", err)
	}

}
