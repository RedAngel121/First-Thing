package main

import (
	"fmt"
	"net/http"
)

// opening Listen and serve on port 8080 for index function

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

// index response and HTTP header and body creation

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "MY THING IS WORKING")
	return
}
