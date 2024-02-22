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
	} else {
		article.Title = r.Title
	}

	if r.Body == "" {
		r.AddErrField("body", "The field is required")
	} else {
		article.Body = r.Body
	}

	if !r.IsValid() {
		return validator.ErrValidator{
			Fields:    r.ErrFields,
			NonFields: r.ErrNonFields,
		}
	}
	return nil
}

type ArticleUpdateRequest struct {
	Title               string `json:"title"`
	Body                string `json:"body"`
	validator.Validator `json:"-"`
}

func (r *ArticleUpdateRequest) Validate(article *models.Article) error {
	if !r.IsValid() {
		return validator.ErrValidator{
			Fields:    r.ErrFields,
			NonFields: r.ErrNonFields,
		}
	}

	if r.Title != article.Title && r.Title != "" {
		article.Title = r.Title
	}
	if r.Body != article.Body && r.Body != "" {
		article.Body = r.Body
	}

	return nil
}

type ArticleResponse struct {
	ID    uint   `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}
