package app

import (
	"database/sql"
	db2 "github.com/ridwanakf/nadc-intro-to-rest/internal/repo/db"

	"github.com/ridwanakf/nadc-intro-to-rest/internal"
)

type Repos struct {
	bookRepo internal.BookRepo
}

func newRepos(db *sql.DB) (*Repos, error) {
	r := &Repos{
		bookRepo: db2.NewBookDB(db),
	}

	return r, nil
}

func (r *Repos) Close() []error {
	var errs []error
	return errs
}
