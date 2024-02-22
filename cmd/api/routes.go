package main

import (
	"github.com/agung96tm/golearn-packages/ui"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app application) routes() http.Handler {
	router := httprouter.New()

	fileServer := http.FileServer(http.FS(ui.Files))
	router.Handler(http.MethodGet, "/media/*filepath", fileServer)
	router.Handler(http.MethodGet, "/static/*filepath", fileServer)

	router.HandlerFunc(http.MethodGet, "/articles", app.articleList)
	router.HandlerFunc(http.MethodGet, "/articles/:id", app.articleDetail)
	router.HandlerFunc(http.MethodPost, "/articles", app.articleCreate)
	router.HandlerFunc(http.MethodPatch, "/articles/:id", app.articleUpdate)
	router.HandlerFunc(http.MethodDelete, "/articles/:id", app.articleDelete)

	router.HandlerFunc(http.MethodPost, "/articles-upload", app.articleUpload)

	return app.recoverPanic(app.enableCORS(router))
}
