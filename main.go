// This is the file name and main starting point for then program
package main
// Importing the stuff that the language needs to run itself
import (
	"net/http" //web import
	"fmt" //formatting import
)

// Main function that handles the arguments and address info for the response

func main() {
http.HandleFunc("/", index) //if slash then set index
	http.ListenAndServe(":8080", nil) // the address used to show the response
}

// This is the index from above where you need to make a response and display it

func index(writer http.ResponseWriter, request *http.Request) { //the index fall back pattern that uses a HTTP header for easy browser access
	fmt.Fprintf(writer, "whatever") // The Hello World line displayed on the port in the browser
	return // end of the code
}
// COMMENT