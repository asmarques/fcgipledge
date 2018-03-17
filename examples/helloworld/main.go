package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/asmarques/fcgipledge"
)

func main() {
	log.SetFlags(0)

	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s socket\n", os.Args[0])
	}

	path := os.Args[1]

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world")
	})

	fcgipledge.ListenAndServe(path, nil, nil)
}
