package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("API")

	mux := http.NewServeMux()

	// routes
	mux.HandleFunc("GET /comment", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "get comment")
	})

	mux.HandleFunc("GET /comment/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue(("id"))
		fmt.Fprintf(w, "get comment by id: %s", id)
	})

	mux.HandleFunc("POST /comment", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "post comment")
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	// server
	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println("error: ", err.Error())
	}
}
