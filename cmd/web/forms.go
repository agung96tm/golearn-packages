package main

import (
	"github.com/agung96tm/golearn-packages/internal/form"
	"github.com/agung96tm/golearn-packages/internal/models"
)

type ArticleForm struct {
	Title     *string `form:"title"`
	Body      *string `form:"body"`
	form.Form `form:"-"`
}

func (f *ArticleForm) Validate(article *models.Article) error {
	if f.Title == nil || *f.Title == "" {
		f.AddErrField("Title", "The field is required")
	} else {
		article.Title = *f.Title
	}

	if f.Body == nil || *f.Body == "" {
		f.AddErrField("Body", "The field is required")
	} else {
		article.Body = *f.Body
	}

	if !f.IsValid() {
		return form.ErrForm
	}

	return nil
}

type ArticleEditForm struct {
	Title     *string `form:"title"`
	Body      *string `form:"body"`
	form.Form `form:"-"`
}

func (f *ArticleEditForm) Validate(article *models.Article) error {
	if f.Title == nil || *f.Title == "" {
		f.AddErrField("Title", "The field is required")
	}
	if f.Body == nil || *f.Body == "" {
		f.AddErrField("Body", "The field is required")
	}

	if *f.Title != article.Title {
		article.Title = *f.Title
	}
	if *f.Body != article.Body {
		article.Body = *f.Body
	}

	if !f.IsValid() {
		return form.ErrForm
	}

	return nil
}

func (f *ArticleEditForm) BindModel(article *models.Article) {
	f.Title = &article.Title
	f.Body = &article.Body
}
