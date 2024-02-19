package main

import (
	"github.com/agung96tm/golearn-packages/internal/form"
	"github.com/agung96tm/golearn-packages/internal/models"
)

type ArticleForm struct {
	Title string `form:"title"`
	Body  string `form:"body"`
	form.Form
}

func (r *ArticleForm) Validate() error {
	if r.Title == "" {
		r.AddErrField("Title", "The field is required")
	}
	if r.Body == "" {
		r.AddErrField("Body", "The field is required")
	}

	if !r.IsValid() {
		return form.ErrForm
	}
	return nil
}

type ArticleEditForm struct {
	Title string `form:"title"`
	Body  string `form:"body"`
	form.Form
}

func (r *ArticleEditForm) Bind(article *models.Article) error {
	r.Title = article.Title
	r.Body = article.Body
	return nil
}

func (r *ArticleEditForm) Validate(article *models.Article) error {
	if r.Title == "" {
		r.AddErrField("Title", "The field is required")
	}
	if r.Body == "" {
		r.AddErrField("Body", "The field is required")
	}
	if !r.IsValid() {
		return form.ErrForm
	}

	if article.Title != r.Title {
		article.Title = r.Title
	}
	if article.Body != r.Body {
		article.Body = r.Body
	}

	return nil
}
