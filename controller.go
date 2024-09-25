package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// showLib handles requests to display the library
func showLib(w http.ResponseWriter, r *http.Request) {
	// Ensure the correct path is accessed
	if r.URL.Path != "/library" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// Fetch the library using the service and write the books to the HTTP response
	books := ls.GetLibrary() // Get all books
	for _, book := range books {
		fmt.Fprintf(w, "Name: %s, ID: %s, Price: %.2f\n", book.Name, book.ID, book.Price)
	}
}

func rentBook(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/rent" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	name := r.FormValue("name")
	rented := ls.RentBook(name)
	if !rented {
		fmt.Fprintf(w, "Can't rent book %s", name)
		return
	}
	if rented {
		fmt.Fprintf(w, "Rented book %s", name)
		return
	}
}

// insertBook handles requests to insert a new book
func insertBook(w http.ResponseWriter, r *http.Request) {
	// Ensure the correct path is accessed
	if r.URL.Path != "/newbook" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// Extract form values from the HTTP request
	name := r.FormValue("name")
	id := r.FormValue("id")
	price := r.FormValue("price")

	// Parse the price from string to float
	p, err := strconv.ParseFloat(price, 64)
	if err != nil {
		http.Error(w, "Invalid price format", http.StatusBadRequest)
		return
	}

	// Add the book using the service
	err = ls.AddBook(name, id, p)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error adding book: %v", err), http.StatusInternalServerError)
		return
	}

	// Acknowledge that the book has been added
	fmt.Fprintf(w, "Book %s added successfully!\n", name)
}

func returnBook(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/return" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	name := r.FormValue("name")
	returned := ls.ReturnBook(name)
	if !returned {
		fmt.Fprintf(w, "Couldn't return: %s", name)
	} else {
		fmt.Fprintf(w, "Book %s return successfully", name)
	}
}
