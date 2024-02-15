package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/articles", app.articleList)
	router.HandlerFunc(http.MethodPost, "/articles", app.articleCreate)

	return router
}
