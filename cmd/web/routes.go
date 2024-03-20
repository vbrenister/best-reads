package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/vbrenister/best-reads/ui"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	fs := http.FileServerFS(ui.Files)

	router.Handler(http.MethodGet, "/static/*filepath", fs)
	router.HandlerFunc(http.MethodGet, "/", app.homeHandler)
	router.HandlerFunc(http.MethodGet, "/books", app.booksHandler)

	return router
}
