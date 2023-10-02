package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a handler function to handle incoming HTTP requests.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Server")
	})

	// Specify the network address and port to listen on.
	address := ":8080"

	// Start the HTTP server.
	fmt.Printf("Server is listening on %s\n", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
