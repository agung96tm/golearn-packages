package main

import (
	"github.com/agung96tm/golearn-packages/ui"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app application) routes() http.Handler {
	router := httprouter.New()
	fileServer := http.FileServer(http.FS(ui.Files))
	router.Handler(http.MethodGet, "/static/*filepath", fileServer)

	router.HandlerFunc(http.MethodGet, "/", app.home)

	router.HandlerFunc(http.MethodGet, "/articles", app.articleList)
	router.HandlerFunc(http.MethodGet, "/articles/create", app.articleCreate)
	router.HandlerFunc(http.MethodPost, "/articles/create", app.articleCreatePost)
	router.HandlerFunc(http.MethodGet, "/articles/edit/:id", app.articleEdit)
	router.HandlerFunc(http.MethodPost, "/articles/edit/:id", app.articleEditPost)
	router.HandlerFunc(http.MethodPost, "/articles/delete/:id", app.articleDeletePost)

	return app.sessionManager.LoadAndSave(app.coreMiddleware(app.noSurf(router)))
}
