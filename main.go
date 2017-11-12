package main

import (
	"net/http"
	"fmt"
)

// opening Listen and serve on port 8080 for index function

func main() {
http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

// index response and HTTP header and body creation for assigned formatting

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "MY THING IS WORKING")
	return
}