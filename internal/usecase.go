package internal

import "github.com/ridwanakf/nadc-intro-to-rest/internal/entity"

// BookUC contains logic for all Book Handlers
//go:generate mockgen -destination=usecase/book_mock.go -package=usecase github.com/ridwanakf/nadc-intro-to-rest/internal BookUC
type BookUC interface {
	GetBookByID(bookID int32) (entity.Book, error)
	InsertNewBook(book entity.Book) (bool, error)
	UpdateBookRating(bookID int32, bookRate float32) (bool, error)
	DeleteBookByID(bookID int32) (bool, error)
	SearchBookByName(bookName string) ([]entity.Book, error)
}
