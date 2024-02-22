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

func (f *ArticleForm) Validate(article *models.Article) error {
	if f.Title == "" {
		f.AddErrField("Title", "The field is required")
	} else {
		article.Title = f.Title
	}
	if f.Body == "" {
		f.AddErrField("Body", "The field is required")
	} else {
		article.Body = f.Body
	}

	if !f.IsValid() {
		return form.ErrForm
	}

	return nil
}

type ArticleEditForm struct {
	Title string `form:"title"`
	Body  string `form:"body"`
	form.Form
}

func (f *ArticleEditForm) BindModel(article *models.Article) {
	f.Title = article.Title
	f.Body = article.Body
}

func (f *ArticleEditForm) Validate(article *models.Article) error {
	if !f.IsValid() {
		return form.ErrForm
	}

	if f.Title != article.Title && f.Title != "" {
		article.Title = f.Title
	}
	if f.Body != article.Body && f.Body != "" {
		article.Body = f.Body
	}

	return nil
}
