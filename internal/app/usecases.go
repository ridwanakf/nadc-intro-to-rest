package app

import (
	"github.com/ridwanakf/nadc-intro-to-rest/internal"
	"github.com/ridwanakf/nadc-intro-to-rest/internal/usecase"
)

type Usecases struct {
	Book internal.BookUC //Book Usecase
}

func newUsecases(repos *Repos) *Usecases {
	return &Usecases{
		Book: usecase.NewBookUsecase(repos.bookRepo),
	}
}

func (*Usecases) Close() []error {
	var errs []error
	return errs
}
