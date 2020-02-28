package internal

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Index is the home page handler
func (h *Handler) Index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	renderJSON(w, []byte(`
		{
			"module": "nadc-intro-to-rest",
			"version": "1.0.0", 
			"date":"29 February 2020",
			"desc": "This is index page"
		}
	`), http.StatusOK)

}

// ServeHTTP is used for 404 page
func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	renderJSON(w, []byte(`
		{
			"message": "There's nothing here"
		}
	`), http.StatusNotFound)
}
