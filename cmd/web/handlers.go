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
	var articleForm ArticleForm
	err := app.PostForm(r, &articleForm)
	if err != nil {
		app.serverError(w, err)
		return
	}

	if app.ArticleServiceCreate(&articleForm); !articleForm.IsValid() {
		data := app.newTemplateData(r)
		data.Form = articleForm
		app.render(w, http.StatusUnprocessableEntity, "article_create.tmpl", data)
		return
	}

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
	data.Form = ArticleEditBindWithModel(article)

	app.render(w, http.StatusOK, "article_edit.tmpl", data)
}

func (app *application) ArticleEditPost(w http.ResponseWriter, r *http.Request) {
	id, _ := app.readIDParam(r)
	article, err := app.ArticleServiceGet(id)
	if err != nil {
		app.notFound(w, r, "/articles")
	}

	var articleForm ArticleEditForm
	if err = app.PostForm(r, &articleForm); err != nil {
		app.serverError(w, err)
		return
	}

	if app.ArticleServiceUpdate(article, &articleForm); !articleForm.IsValid() {
		data := app.newTemplateData(r)
		data.Form = articleForm
		data.Article = article
		app.render(w, http.StatusUnprocessableEntity, "article_edit.tmpl", data)
		return
	}

	app.sessionManager.Put(r.Context(), "flash", "Article Success Updated!")
	app.redirect(w, r, "/articles")
}
