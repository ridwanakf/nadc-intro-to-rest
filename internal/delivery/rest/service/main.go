package service

import "github.com/ridwanakf/nadc-intro-to-rest/internal/app"

type Services struct {
	*BookService
	*DefaultService
}

func GetServices(app *app.NadcRest) *Services {
	return &Services{
		BookService:    NewBookService(app),
		DefaultService: NewDefaultService(),
	}
}
