package main

import (
	"log"
	"net/http"
)

const (
	port = ":3000"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	log.Printf("serving on http://localhost%s/", port)
	http.ListenAndServe(port, nil)
}