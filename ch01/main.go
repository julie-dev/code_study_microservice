package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080
	log.Printf("Server starting on port %v\n", 8080)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func init() {
	//http.HandleFunc("/helloworld", helloWorldHandler)
	http.Handle("/helloworld", newRequestHandler(newResponseHandler()))
	http.Handle("/cat/", http.StripPrefix("/cat/", http.FileServer(http.Dir("images"))))
}
