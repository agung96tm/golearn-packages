package main

import (
	"errors"
	"fmt"
	"github.com/agung96tm/golearn-packages/internal/form"
	"github.com/agung96tm/golearn-packages/internal/models"
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	app.redirect(w, r, "/articles")
}

func (app *application) ArticleList(w http.ResponseWriter, r *http.Request) {
	articles, err := app.ArticleServiceGetAll()
	if err != nil {
		app.serverError(w, err)
	}

	data := app.newTemplateData(r)
	data.Articles = articles
	app.render(w, http.StatusOK, "article_list.tmpl", data)
}

func (app *application) ArticleCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = ArticleForm{}

	app.render(w, http.StatusOK, "article_create.tmpl", data)
}

func (app *application) ArticleCreatePost(w http.ResponseWriter, r *http.Request) {
	articleForm := ArticleForm{}
	if err := app.PostForm(r, &articleForm); err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	article, err := app.ArticleServiceCreate(&articleForm)
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

	app.sessionManager.Put(
		r.Context(),
		"flash",
		fmt.Sprintf("Article `%s` Success Created!", article.Title),
	)
	app.redirect(w, r, "/articles")
}

func (app *application) ArticleEdit(w http.ResponseWriter, r *http.Request) {
	id, _ := app.readIDParam(r)
	article, err := app.ArticleServiceGet(id)
	if err != nil {
		app.notFound(w, r, "/articles")
		return
	}

	articleForm := ArticleEditForm{}
	if err := articleForm.Bind(article); err != nil {
		app.serverError(w, err)
	}

	data := app.newTemplateData(r)
	data.Article = article
	data.Form = articleForm

	app.render(w, http.StatusOK, "article_edit.tmpl", data)
}

func (app *application) ArticleEditPost(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFound(w, r, "/articles")
		return
	}

	articleForm := ArticleEditForm{}
	if err := app.PostForm(r, &articleForm); err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	article, err := app.ArticleServiceUpdate(id, &articleForm)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNotFound):
			app.notFound(w, r, "/articles")

		case errors.Is(err, form.ErrForm):
			data := app.newTemplateData(r)
			data.Form = articleForm
			data.Article = article
			app.render(w, http.StatusUnprocessableEntity, "article_edit.tmpl", data)

		default:
			app.serverError(w, err)
		}
		return
	}

	app.sessionManager.Put(
		r.Context(),
		"flash",
		fmt.Sprintf("Article `%s` Success Updated!", article.Title),
	)
	app.redirect(w, r, "/articles")
}

func (app *application) ArticleDeletePost(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFound(w, r, "/articles")
		return
	}

	err = app.ArticleServiceDelete(id)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNotFound):
			app.notFound(w, r, "/articles")
		default:
			app.serverError(w, err)
		}
		return
	}

	app.sessionManager.Put(r.Context(), "flash", "Article Success Deleted!")
	app.redirect(w, r, "/articles")
}
