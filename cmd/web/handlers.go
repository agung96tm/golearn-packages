package main

import (
	"github.com/agung96tm/golearn-packages/internal/form"
	"github.com/agung96tm/golearn-packages/internal/models"
	"net/http"
)

type ArticleForm struct {
	Title     string `form:"title" validate:"required,min=3"`
	Body      string `form:"body" validate:"required,min=10"`
	form.Form `form:"-"`
}

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	app.redirect(w, r, "/articles")
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

	if !articleForm.IsValid() {
		data := app.newTemplateData(r)
		data.Form = articleForm
		app.render(w, http.StatusUnprocessableEntity, "article_create.tmpl", data)
		return
	}

	/* SAVE ARTICLES */
	models.ArticleData = append(models.ArticleData, models.Article{
		Title: articleForm.Title,
		Body:  articleForm.Body,
	})

	app.sessionManager.Put(r.Context(), "flash", "Article Success Created!")
	app.redirect(w, r, "/articles")
}
