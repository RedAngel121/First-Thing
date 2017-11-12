package main

import (
	"net/http"
	"fmt"
)

// Function initialization, port 8080 listen and serve for index

func main() {
http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

// index response and HTTP header and body creation for assigned formatting

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "whatever")
	return
}