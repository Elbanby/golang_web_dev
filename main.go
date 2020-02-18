package main

import (
	"fmt"
	"net/http"
)

func simpleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome</h1>")
}

func main() {
	http.HandleFunc("/", simpleHandler)
	if err := http.ListenAndServe(":9000", nil); err != nil {
		fmt.Println("Error occurred while binding to port", err)
	}

}
