package main

import (
	"fmt"
	"net/http"
)

func simpleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "<h1>Welcome home!</h1>")
	case "/contact":
		fmt.Fprint(w, `
						<h1>Contact info:</h1>
						email: <a href="mailto:support@mail.com">Support</a>
		`)
	default:
		fmt.Fprint(w, `<h1>404</h1>`)
	}
}

func main() {
	http.HandleFunc("/", simpleHandler)
	if err := http.ListenAndServe(":9000", nil); err != nil {
		fmt.Println("Error occurred while binding to port", err)
	}

}
