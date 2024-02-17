package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/articles", app.articleList)
	router.HandlerFunc(http.MethodGet, "/articles/:id", app.articleDetail)
	router.HandlerFunc(http.MethodPost, "/articles", app.articleCreate)
	router.HandlerFunc(http.MethodPatch, "/articles/:id", app.articleUpdate)

	return router
}
