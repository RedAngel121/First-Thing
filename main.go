package main

import (
	"net/http"
	"fmt"
)

func main() {
http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "whatever")
	return
}