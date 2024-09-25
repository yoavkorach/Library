package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type Book struct {
	Name   string
	ID     string
	Price  float64
	rented bool
}

type Library struct {
	books []Book
}

// openLibrary loads the books from an Excel file or creates a new file if it doesn't exist
func openLibrary() *Library {
	lib := &Library{}
	// Check if file exists
	if _, err := os.Stat("library.xlsx"); os.IsNotExist(err) {
		fmt.Println("Library file not found, creating a new one")
		// Create a new Excel file
		f := excelize.NewFile()
		// Add headers to the sheet
		f.SetCellValue("Sheet1", "A1", "Name")
		f.SetCellValue("Sheet1", "B1", "ID")
		f.SetCellValue("Sheet1", "C1", "Price")
		f.SetCellValue("sheet1", "D1", "Rented")
		f.SaveAs("library.xlsx")
		return lib
	}

	// Open existing Excel file
	f, err := excelize.OpenFile("library.xlsx")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer f.Close()

	// Read rows from the file
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	// Skip the first row (header)
	for _, row := range rows[1:] {
		if len(row) < 3 {
			continue
		}
		price, _ := strconv.ParseFloat(row[2], 64)
		book := Book{Name: row[0], ID: row[1], Price: price}
		lib.books = append(lib.books, book)
	}

	return lib
}

// AddBook adds a new book to the library and saves it to the Excel file
func (l *Library) AddBook(name string, id string, price float64) error {
	book := Book{name, id, price, false}
	l.books = append(l.books, book)

	// Open Excel file
	f, err := excelize.OpenFile("library.xlsx")
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer f.Close()

	// Get the next available row
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return fmt.Errorf("error reading rows: %v", err)
	}
	nextRow := len(rows) + 1

	// Write new book data
	rowStr := fmt.Sprintf("%d", nextRow)
	f.SetCellValue("Sheet1", "A"+rowStr, name)
	f.SetCellValue("Sheet1", "B"+rowStr, id)
	f.SetCellValue("Sheet1", "C"+rowStr, price)
	f.SetCellValue("sheet1", "D"+rowStr, false)

	// Save the file
	if err := f.Save(); err != nil {
		return fmt.Errorf("error saving file: %v", err)
	}

	return nil // Successfully added book
}

// GetBooks returns the list of books from the library
func (l *Library) GetBooks() []Book {
	return l.books
}

func (l *Library) RentBook(bookName string) bool {
	f, err := excelize.OpenFile("library.xlsx")
	if err != nil {
		return false
	}
	defer f.Close()
	j := 0
	for i := range l.books {
		if l.books[i].Name == bookName && !l.books[i].rented {
			j = i + 2
			l.books[i].rented = true
			break
		}
	}
	j_str := fmt.Sprintf("%d", j)
	if j != 0 {
		f.SetCellValue("sheet1", "D"+j_str, true)
		if err := f.Save(); err != nil {
			return false
		}
		return true
	}
	return false
}

func (l *Library) returnBook(bookName string) bool {
	f, err := excelize.OpenFile("library.xlsx")
	if err != nil {
		return false
	}
	defer f.Close()
	j := 0
	for i := range l.books {
		if l.books[i].Name == bookName && l.books[i].rented {
			j = i + 2
			l.books[i].rented = false
			break
		}
	}
	j_str := fmt.Sprintf("%d", j)
	if j != 0 {
		f.SetCellValue("sheet1", "D"+j_str, false)
		if err := f.Save(); err != nil {
			return false
		}
		return true
	}
	return false
}
