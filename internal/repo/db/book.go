package db

import (
	"database/sql"
	"log"

	"github.com/pkg/errors"
	"github.com/ridwanakf/nadc-intro-to-rest/internal/entity"
)

type BookDB struct {
	db *sql.DB
}

func NewBookDB(db *sql.DB) *BookDB {
	return &BookDB{
		db: db,
	}
}

func (b *BookDB) GetBookByID(bookID int32) (entity.Book, error) {
	stmt, err := b.db.Prepare(SQLGetBookByID)
	if err != nil {
		return entity.Book{}, errors.Errorf("[repo][GetBookByID] invalid prepare statement :%+v", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(bookID)
	var books []entity.Book
	for rows.Next() {
		book := entity.Book{}
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Rate,
			&book.Category,
		)
		if err != nil {
			log.Printf("[repo][GetBookByID] fail to scan :%+v\n", err)
			continue
		}
		books = append(books, book)
	}
	if len(books) < 1 {
		return entity.Book{}, errors.New("[repo][GetBookByID] book not found!")
	}
	return books[0], nil
}

func (b *BookDB) InsertNewBook(book entity.Book) (bool, error) {
	stmt, err := b.db.Prepare(SQLInsertNewBook)
	if err != nil {
		return false, errors.Errorf("[repo][InsertNewBook] invalid prepare statement :%+v\n", err)
	}
	defer stmt.Close()

	_, err = stmt.Query(book.Title, book.Author, book.Rate, book.Category)
	if err != nil {
		return false, errors.Errorf("[repo][InsertNewBook] error inserting new book :%+v\n", err)
	}

	return true, nil
}

func (b *BookDB) UpdateBookRating(bookID int32, bookRate float32) (bool, error) {
	stmt, err := b.db.Prepare(SQLUpdateBookRating)
	if err != nil {
		return false, errors.Errorf("[repo][UpdateBookRating] invalid prepare statement :%+v\n", err)
	}
	defer stmt.Close()

	_, err = stmt.Query(bookRate, bookID)
	if err != nil {
		return false, errors.Errorf("[repo][UpdateBookRating] error updating rating :%+v\n", err)
	}

	return true, nil
}

func (b *BookDB) DeleteBookByID(bookID int32) (bool, error) {
	stmt, err := b.db.Prepare(SQLDeleteBookByID)
	if err != nil {
		return false, errors.Errorf("[repo][DeleteBookByID] invalid prepare statement :%+v\n", err)
	}
	defer stmt.Close()

	_, err = stmt.Query(bookID)
	if err != nil {
		return false, errors.Errorf("[repo][DeleteBookByID] error inserting new book :%+v\n", err)
	}

	return true, nil
}

func (b *BookDB) SearchBookByName(bookName string) ([]entity.Book, error) {
	var res []entity.Book

	stmt, err := b.db.Prepare(SQLSearchBookByName)
	if err != nil {
		return res, errors.Errorf("[repo][SearchBookByName] invalid prepare statement :%+v\n", err)
	}
	defer stmt.Close()

	//execute query
	rows, err := stmt.Query("%" + bookName + "%")

	for rows.Next() {
		book := entity.Book{}
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Rate,
			&book.Category,
		)
		if err != nil {
			log.Printf("[repo][SearchBookByName] fail to scan :%+v\n", err)
			continue
		}
		res = append(res, book)
	}

	return res, nil
}
