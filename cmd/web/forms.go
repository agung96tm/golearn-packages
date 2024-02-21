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
	return nil
}
