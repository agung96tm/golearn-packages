package main

import "github.com/agung96tm/golearn-packages/internal/form"

type ArticleForm struct {
	Title string `form:"title"`
	Body  string `form:"body"`
	form.Form
}

type ArticleEditForm struct {
	Title string `form:"title"`
	Body  string `form:"body"`
	form.Form
}
