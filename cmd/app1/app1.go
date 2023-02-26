package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting Application 1 - This is a server")

	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello from application 1\n")
}
