package internal

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *Handler) GetBookByID(w http.ResponseWriter, r *http.Request, param httprouter.Params){
	// ambil param

	// persiapkan prepare statement sql

	// execute query

	// marshal array books

	// return response
}

func (h *Handler) InsertNewBook(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	// baca body

	// parse body

	// persiapkan prepare statement sql

	// execute query

	// return response
}

func (h *Handler) UpdateBookRating(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	// ambil param

	// baca body

	// parse body

	// persiapkan prepare statement sql

	// execute query

	// return response
}

func (h *Handler) DeleteBookByID(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	// baca param

	// persiapkan prepare statement sql

	//execute query

	//return response
}

