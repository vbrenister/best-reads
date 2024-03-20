package main

import (
	"net/http"
)

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	data := newTemplateData()

	app.render(w, r, http.StatusOK, "home.tmpl", data)
}

func (app *application) booksHandler(w http.ResponseWriter, r *http.Request) {
	data := newTemplateData()

	app.render(w, r, http.StatusOK, "books.tmpl", data)
}
