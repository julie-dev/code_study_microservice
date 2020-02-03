package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080
	http.Handle("/helloworld", newRequestHandler(newResponseHandler()))

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
