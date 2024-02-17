package main

import (
	"github.com/agung96tm/golearn-packages/internal/form"
	"github.com/agung96tm/golearn-packages/internal/models"
)

type ArticleForm struct {
	Title     *string `form:"title" validate:"required,min=3"`
	Body      *string `form:"body" validate:"required,min=10"`
	form.Form `form:"-"`
}

type ArticleEditForm struct {
	Title     *string `form:"title"`
	Body      *string `form:"body"`
	form.Form `form:"-"`
}

func ArticleEditBindWithModel(article *models.Article) *ArticleEditForm {
	form := ArticleEditForm{}
	form.Title = &article.Title
	form.Body = &article.Body
	return &form
}
