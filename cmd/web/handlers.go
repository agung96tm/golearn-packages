package main

import (
	"net/http"
)

func (app application) home(w http.ResponseWriter, r *http.Request) {
	app.redirect(w, r, "/articles")
}

func (app application) articleList(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Articles = app.articleServiceGetAll()
	app.render(w, http.StatusOK, "article_list.tmpl", data)
}

func (app application) articleCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = ArticleForm{}

	app.render(w, http.StatusOK, "article_create.tmpl", data)
}

func (app application) articleCreatePost(w http.ResponseWriter, r *http.Request) {
	app.sessionManager.Put(r.Context(), "flash", "Article Success Created!")
	app.redirect(w, r, "/articles")
}

func (app application) articleEdit(w http.ResponseWriter, r *http.Request) {
	id, _ := app.readIDParam(r)
	article, err := app.articleServiceGet(id)
	if err != nil {
		app.notFound(w, r, "/articles")
		return
	}

	form := ArticleEditForm{}
	form.BindModel(article)

	data := app.newTemplateData(r)
	data.Article = article
	data.Form = form

	app.render(w, http.StatusOK, "article_edit.tmpl", data)
}

func (app application) articleEditPost(w http.ResponseWriter, r *http.Request) {
	app.sessionManager.Put(r.Context(), "flash", "Article Success Updated!")
	app.redirect(w, r, "/articles")
}

func (app application) articleDeletePost(w http.ResponseWriter, r *http.Request) {
	app.sessionManager.Put(r.Context(), "flash", "Article Success Deleted!")
	app.redirect(w, r, "/articles")
}
