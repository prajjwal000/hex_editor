package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)

	srv := http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	log.Println("Starting server on :8080")
	err := srv.ListenAndServe()
	log.Fatal(err)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Sailor\n")
}
