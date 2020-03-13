package main

import (
	"log"

	"github.com/ridwanakf/nadc-intro-to-rest/internal/app"
	"github.com/ridwanakf/nadc-intro-to-rest/internal/delivery/rest"
)

func main() {
	// init app
	nadcRestApp, err := app.NewNadcRest()
	if err != nil {
		log.Fatalf("marshal error %+v", err)
	}
	defer func() {
		errs := nadcRestApp.Close()
		for e := range errs {
			log.Println(e)
		}
	}()

	rest.Start(nadcRestApp)
}
