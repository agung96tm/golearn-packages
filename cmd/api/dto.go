package main

type ArticleCreateRequest struct {
	Title string `json:"title" validate:"required,min=10,max=200"`
	Body  string `json:"body" validate:"required,min=10,max=1000"`
}
