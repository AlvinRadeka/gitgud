package main

import (
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	fmt.Println("do something 1")
	fmt.Println("do something 2")

	// add something later

	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
