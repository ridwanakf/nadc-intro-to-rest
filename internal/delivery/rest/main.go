package rest

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/gops/agent"
	"github.com/julienschmidt/httprouter"
	"github.com/ridwanakf/nadc-intro-to-rest/internal/app"
	"github.com/ridwanakf/nadc-intro-to-rest/internal/delivery/rest/service"
	"github.com/ridwanakf/nadc-intro-to-rest/internal/entity"
)

func initFlags(args *entity.Args) {
	port := os.Getenv("PORT")
	if port == "" {
		port = *(flag.String("PORT", "5000", "port number for your apps"))
	}
	args.Port = port
}

func initRouter(router *httprouter.Router, svc *service.Services) {

	//Default path
	router.GET("/", svc.DefaultService.Index)

	//Book API paths
	router.GET("/book/:bookID", svc.BookService.GetBookByID)
	router.POST("/book", svc.BookService.InsertNewBook)
	router.PUT("/book/:bookID", svc.BookService.UpdateBookRating)
	router.DELETE("/book/:bookID", svc.BookService.DeleteBookByID)
	router.GET("/search", svc.BookService.SearchBookByName)

	// `httprouter` library uses `ServeHTTP` method for it's 404 pages
	router.NotFound = svc.DefaultService
}

func Start(app *app.NadcRest) {
	args := new(entity.Args)
	initFlags(args)

	svc := service.GetServices(app)
	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatal(err)
	}

	router := httprouter.New()
	initRouter(router, svc)

	fmt.Println("Apps served on :" + args.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":"+args.Port), router))
}
