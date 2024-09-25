package main

import (
	"fmt"
	"net/http"
)

var ls LibService // Global service variable

func main() {
	// Initialize the library and service
	library := openLibrary()    // Load or create the library from the Excel file
	ls = NewLibService(library) // Initialize the service with the library

	// Register routes
	http.HandleFunc("/rent", rentBook)
	http.HandleFunc("/library", showLib)
	http.HandleFunc("/newbook", insertBook)
	http.HandleFunc("/return", returnBook)
	// Start the HTTP server
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
