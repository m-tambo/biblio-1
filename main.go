package main

import (
	"biblio/database"
	"biblio/handler"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	db := database.NewDb()
	ad := handler.NewAuthorDAO(db)
	pd := handler.NewPatronDAO(db)
	r := mux.NewRouter()

	r.Handle("/authors", handler.NewGetAuthorsHandler(ad)).Methods(http.MethodGet)
	r.Handle("/authors", handler.NewCreateAuthorHandler(ad)).Methods(http.MethodPost)
	r.Handle("/authors/{id}", handler.NewDeleteAuthorHandler(ad)).Methods(http.MethodDelete)

	r.Handle("/patrons", handler.NewGetPatronsHandler(pd)).Methods(http.MethodGet)

	http.ListenAndServe(":8008", r)
}