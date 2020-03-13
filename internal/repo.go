package internal

import "github.com/ridwanakf/nadc-intro-to-rest/internal/entity"

// BookRepo contains logic for getting book from repo
//go:generate mockgen -destination=repo/book_mock.go -package=repo github.com/ridwanakf/nadc-intro-to-rest/internal BookRepo
type BookRepo interface {
	GetBookByID(bookID int32) (entity.Book, error)
	InsertNewBook(book entity.Book) (bool, error)
	UpdateBookRating(bookID int32, bookRate float32) (bool, error)
	DeleteBookByID(bookID int32) (bool, error)
	SearchBookByName(bookName string) ([]entity.Book, error)
}
