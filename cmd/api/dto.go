package main

import (
	"github.com/agung96tm/golearn-packages/internal/models"
	"github.com/agung96tm/golearn-packages/internal/validator"
)

type ArticleCreateRequest struct {
	Title               string `json:"title"`
	Body                string `json:"body"`
	validator.Validator `json:"-"`
}

func (r *ArticleCreateRequest) Validate(article *models.Article) error {
	if r.Title == "" {
		r.AddErrField("title", "The field is required")
	}
	if r.Body == "" {
		r.AddErrField("body", "The field is required")
	}

	article.ID = uint(len(models.ArticleData) + 1)
	article.Title = r.Title
	article.Body = r.Body

	if !r.IsValid() {
		return validator.ErrValidator{
			Fields:    r.ErrFields,
			NonFields: r.ErrNonFields,
		}
	}
	return nil
}

type ArticleResponse struct {
	ID    uint   `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}
