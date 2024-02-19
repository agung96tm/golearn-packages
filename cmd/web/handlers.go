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
	articles, err := app.articleServiceGetAll()
	if err != nil {
		app.serverError(w, err)
	}

	data := app.newTemplateData(r)
	data.Articles = articles
	app.render(w, http.StatusOK, "article_list.tmpl", data)
}

func (app application) articleCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = ArticleForm{}

	app.render(w, http.StatusOK, "article_create.tmpl", data)
}

func (app application) articleCreatePost(w http.ResponseWriter, r *http.Request) {
	articleForm := ArticleForm{}
	if err := app.PostForm(r, &articleForm); err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	/*
		-- Example with DB Transaction --
		trxHandle := r.Context().Value(constants.DBTransaction).(*gorm.DB)
		resp, err := app.articleServiceCreate(trxHandle, &articleForm)
	*/

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

	app.sessionManager.Put(
		r.Context(),
		"flash",
		fmt.Sprintf("Article `%s` Success Created!", article.Title),
	)
	app.redirect(w, r, "/articles")
}

func (app application) articleEdit(w http.ResponseWriter, r *http.Request) {
	id, _ := app.readIDParam(r)
	article, err := app.articleServiceGet(id)
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

func (app application) articleEditPost(w http.ResponseWriter, r *http.Request) {
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

	article, err := app.articleServiceUpdate(id, &articleForm)
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

func (app application) articleDeletePost(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFound(w, r, "/articles")
		return
	}

	err = app.articleServiceDelete(id)
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
