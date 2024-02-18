package main

import (
	"github.com/agung96tm/golearn-packages/internal/models"
	"github.com/agung96tm/golearn-packages/internal/validator"
)

type ArticleCreateRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (a ArticleCreateRequest) Validate(v *validator.Validator) {
	if a.Title == "" {
		v.AddErrField("title", "the field is required")
	}
	if a.Body == "" {
		v.AddErrField("body", "the field is required")
	}
}

type ArticleUpdateRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (a ArticleUpdateRequest) Validate(v *validator.Validator, article *models.Article) {
	if a.Title != article.Title {
		article.Title = a.Title
	}
	if a.Body != article.Body {
		article.Body = a.Body
	}
}

type ArticleResponse struct {
	ID    uint   `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}
