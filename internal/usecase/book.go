package usecase

import (
	"github.com/ridwanakf/nadc-intro-to-rest/internal"
	"github.com/ridwanakf/nadc-intro-to-rest/internal/entity"
)

type BookUsecase struct {
	repo internal.BookRepo
}

func NewBookUsecase(repo internal.BookRepo) *BookUsecase {
	return &BookUsecase{
		repo: repo,
	}
}

func (b *BookUsecase)GetBookByID(bookID int32) (entity.Book, error){
	return b.repo.GetBookByID(bookID)
}

func (b *BookUsecase)InsertNewBook(book entity.Book) (bool, error){
	return b.repo.InsertNewBook(book)
}

func (b *BookUsecase)UpdateBookRating(bookID int32, bookRate float32) (bool, error){
	return b.repo.UpdateBookRating(bookID, bookRate)
}

func (b *BookUsecase)DeleteBookByID(bookID int32) (bool, error){
	return b.repo.DeleteBookByID(bookID)
}

func (b *BookUsecase)SearchBookByName(bookName string) ([]entity.Book, error){
	return b.repo.SearchBookByName(bookName)
}
