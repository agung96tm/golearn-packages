package main

import (
	"errors"
	"fmt"
	"github.com/agung96tm/golearn-packages/internal/form"
	"github.com/agung96tm/golearn-packages/internal/models"
	"net/http"
)

func (app application) home(w http.ResponseWriter, r *http.Request) {
	app.redirect(w, r, "/articles")
}

func (app application) articleList(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Articles = models.ArticleData

	app.render(w, http.StatusOK, "article_list.tmpl", data)
}

func (app application) articleCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = ArticleForm{}

	app.render(w, http.StatusOK, "article_create.tmpl", data)
}

func (app application) articleCreatePost(w http.ResponseWriter, r *http.Request) {
	var articleForm ArticleForm
	if err := app.PostForm(r, &articleForm); err != nil {
		app.clientError(w, http.StatusBadRequest)
	}

	article, err := app.articleServiceCreate(&articleForm)
	if err != nil {
		switch {
		case errors.Is(err, form.ErrForm):
			data := app.newTemplateData(r)
			data.Form = articleForm
			app.render(w, http.StatusUnprocessableEntity, "article_create.tmpl", data)
		default:
			app.serverError(w, err)
		}
		return
	}

	app.sessionManager.Put(r.Context(), "flash", fmt.Sprintf("Article `%s` Created!", article.Title))
	app.redirect(w, r, "/articles")
}
