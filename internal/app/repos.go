package app

import (
	"database/sql"

	"github.com/ridwanakf/nadc-intro-to-rest/internal"
	"github.com/ridwanakf/nadc-intro-to-rest/internal/repo"
)

type Repos struct {
	bookRepo internal.BookRepo
}

func newRepos(db *sql.DB) (*Repos, error) {
	r := &Repos{
		bookRepo: repo.NewBookDB(db),
	}

	return r, nil
}

func (r *Repos) Close() []error {
	var errs []error
	return errs
}
