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
	}
	if f.Body == "" {
		f.AddErrField("Body", "The field is required")
	}

	article.ID = uint(len(models.ArticleData) + 1)
	article.Title = f.Title
	article.Body = f.Body

	if !f.IsValid() {
		return form.ErrForm
	}
	return nil
}
