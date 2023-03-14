package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":5000"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "test.page.tmpl")
	})
	fmt.Printf("Starting front end service on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Panic(err)
	}
}
