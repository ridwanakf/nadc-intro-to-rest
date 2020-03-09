package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/ridwanakf/nadc-intro-to-rest/internal"
)

func initFlags(args *internal.Args) {
	port := os.Getenv("PORT")
	if port == "" {
		port = *(flag.String("PORT", "5000", "port number for your apps"))
	}
	args.Port = port
}

func initHandler(handler *internal.Handler) error {

	// Initialize SQL DB
	db, err := sql.Open("postgres", "postgres://username:password@host-address:port/database-name?sslmode=disable")
	if err != nil {
		return err
	}

	// Check if db connected
	if err = db.PingContext(context.Background()); err != nil {
		return err
	}

	handler.DB = db

	return nil
}

func initRouter(router *httprouter.Router, handler *internal.Handler) {

	//Default path
	router.GET("/", handler.Index)

	//Book API paths
	router.GET("/book/:bookID", handler.GetBookByID)
	router.POST("/book", handler.InsertNewBook)
	router.PUT("/book/:bookID", handler.UpdateBookRating)
	router.DELETE("/book/:bookID", handler.DeleteBookByID)
	router.GET("/search", handler.SearchBookByName)

	// `httprouter` library uses `ServeHTTP` method for it's 404 pages
	router.NotFound = handler
}

func main() {
	args := new(internal.Args)
	initFlags(args)

	handler := new(internal.Handler)
	if err := initHandler(handler); err != nil {
		log.Println("Failed to init handler", err)
		panic(err)
	}

	router := httprouter.New()
	initRouter(router, handler)

	fmt.Println("Apps served on :" + args.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":" + args.Port), router))
}
