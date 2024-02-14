package main

import (
	"github.com/agung96tm/golearn-packages/ui"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	fileServer := http.FileServer(http.FS(ui.Files))
	router.Handler(http.MethodGet, "/static/*filepath", fileServer)

	router.HandlerFunc(http.MethodGet, "/", app.Home)
	router.HandlerFunc(http.MethodGet, "/articles", app.ArticleCreate)
	router.HandlerFunc(http.MethodPost, "/articles", app.ArticleCreatePost)

	return app.sessionManager.LoadAndSave(app.noSurf(router))
}
