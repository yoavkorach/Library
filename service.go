package main

// LibService defines the service interface
type LibService interface {
	AddBook(name string, id string, price float64) error
	GetLibrary() []Book
	RentBook(name string) bool
	ReturnBook(name string) bool
}

// libService implements LibService and wraps a Library instance
type libService struct {
	library *Library
}

// NewLibService initializes a new libService with the provided Library
func NewLibService(l *Library) LibService {
	return &libService{library: l}
}

// AddBook adds a book to the library and returns any error encountered
func (ls *libService) AddBook(name string, id string, price float64) error {
	return ls.library.AddBook(name, id, price)
}

// GetLibrary returns the list of books in the library
func (ls *libService) GetLibrary() []Book {
	return ls.library.GetBooks()
}

func (ls *libService) RentBook(name string) bool {
	return ls.library.RentBook(name)
}

func (ls *libService) ReturnBook(name string) bool {
	return ls.library.returnBook(name)
}
