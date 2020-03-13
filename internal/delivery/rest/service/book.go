package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/ridwanakf/nadc-intro-to-rest/internal"
	"github.com/ridwanakf/nadc-intro-to-rest/internal/app"
	"github.com/ridwanakf/nadc-intro-to-rest/internal/entity"
	"github.com/ridwanakf/nadc-intro-to-rest/internal/helper"
)

type BookService struct {
	uc internal.BookUC
}

func NewBookService(app *app.NadcRest) *BookService {
	return &BookService{
		uc: app.UseCases.Book,
	}
}

func (b *BookService) GetBookByID(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	bookID, err := strconv.Atoi(param.ByName("bookID"))
	if err != nil {
		log.Printf("[service][GetBookByID] fail to get param :%+v\n", err)
		return
	}

	res, err := b.uc.GetBookByID(int32(bookID))
	if err != nil {
		log.Printf("[service][GetBookByID] fail to get data :%+v\n", err)
		helper.RenderJSON(w, []byte(`
		{
			status: "failed",
			message: "book not found!"
		}
		`), http.StatusOK)
		return
	}

	bytes, err := json.Marshal(res)
	if err != nil {
		log.Printf("[service][GetBookByID] fail to marshal data :%+v\n", err)
		return
	}

	helper.RenderJSON(w, bytes, http.StatusOK)
}

func (b *BookService) InsertNewBook(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		helper.RenderJSON(w, []byte(`
			message: "Failed to read body"
		`), http.StatusBadRequest)
		return
	}

	var book entity.Book
	err = json.Unmarshal(body, &book)
	if err != nil {
		log.Printf("[service][InsertNewBook] error unmarshal :%+v\n", err)
		return
	}

	res, err := b.uc.InsertNewBook(book)

	if res {
		helper.RenderJSON(w, []byte(`
		{
			status: "success",
			message: "Insert book success!"
		}
		`), http.StatusOK)
	} else {
		helper.RenderJSON(w, []byte(`
		{
			status: "failed",
			message: `+error.Error(err)+`
		}
		`), http.StatusOK)
	}
}

func (b *BookService) UpdateBookRating(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	bookID, err := strconv.Atoi(param.ByName("bookID"))
	if err != nil {
		log.Printf("[service][UpdateBookRating] fail to get param :%+v\n", err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		helper.RenderJSON(w, []byte(`
			message: "Failed to read body"
		`), http.StatusBadRequest)
		return
	}

	var book entity.Book
	err = json.Unmarshal(body, &book)
	if err != nil {
		log.Printf("[service][UpdateBookRating] error unmarshal :%+v\n", err)
		return
	}

	res, err := b.uc.UpdateBookRating(int32(bookID), book.Rate)

	if res {
		helper.RenderJSON(w, []byte(`
		{
			status: "success",
			message: "Update book success!"
		}
		`), http.StatusOK)
	} else {
		helper.RenderJSON(w, []byte(`
		{
			status: "failed",
			message: `+error.Error(err)+`
		}
		`), http.StatusOK)
	}
}

func (b *BookService) DeleteBookByID(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	bookID, err := strconv.Atoi(param.ByName("bookID"))
	if err != nil {
		log.Printf("[service][DeleteBookByID] fail to get param :%+v\n", err)
		return
	}

	res, err := b.uc.DeleteBookByID(int32(bookID))

	if res {
		helper.RenderJSON(w, []byte(`
		{
			status: "success",
			message: "Delete book success!"
		}
		`), http.StatusOK)
	} else {
		helper.RenderJSON(w, []byte(`
		{
			status: "failed",
			message: `+error.Error(err)+`
		}
		`), http.StatusOK)
	}
}

func (b *BookService) SearchBookByName(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		helper.RenderJSON(w, []byte(`
			message: "Failed to read body"
		`), http.StatusBadRequest)
		return
	}

	// parse body
	var book entity.Book
	err = json.Unmarshal(body, &book)
	if err != nil {
		log.Printf("[service][SearchBookByName] error unmarshal :%+v\n", err)
		return
	}

	res, err := b.uc.SearchBookByName(book.Title)
	if err != nil {
		log.Printf("[service][SearchBookByName] fail to get data :%+v\n", err)
		return
	}

	bytes, err := json.Marshal(res)
	if err != nil {
		log.Printf("[service][SearchBookByName] fail to marshal data :%+v\n", err)
		return
	}

	helper.RenderJSON(w, bytes, http.StatusOK)
}
