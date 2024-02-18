package main

import (
	"github.com/agung96tm/golearn-packages/internal/models"
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	app.redirect(w, r, "/articles")
}

func (app *application) ArticleList(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Articles = models.ArticleData

	app.render(w, http.StatusOK, "article_list.tmpl", data)
}

func (app *application) ArticleCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = ArticleForm{}

	app.render(w, http.StatusOK, "article_create.tmpl", data)
}

func (app *application) ArticleCreatePost(w http.ResponseWriter, r *http.Request) {
	app.sessionManager.Put(r.Context(), "flash", "Article Success Created!")
	app.redirect(w, r, "/articles")
}

func (app *application) ArticleEdit(w http.ResponseWriter, r *http.Request) {
	id, _ := app.readIDParam(r)
	article, err := app.ArticleServiceGet(id)
	if err != nil {
		app.notFound(w, r, "/articles")
		return
	}

	data := app.newTemplateData(r)
	data.Article = article
	data.Form = ArticleEditForm{}

	app.render(w, http.StatusOK, "article_edit.tmpl", data)
}

func (app *application) ArticleEditPost(w http.ResponseWriter, r *http.Request) {
	app.sessionManager.Put(r.Context(), "flash", "Article Success Updated!")
	app.redirect(w, r, "/articles")
}
